package database

import (
	"context"
	"log"
	"sync"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/matrosovm/warehouse/internal/pkg/domain"
	"github.com/matrosovm/warehouse/internal/pkg/helpers"
)

const (
	workerPoolSize = 3
)

func (s *storeImpl) ReserveProducts(filter *domain.Filter) (map[uint64]bool, error) {
	ctx := context.Background()

	productsID := make([]uint64, 0, len(filter.Products))
	for k := range filter.Products {
		productsID = append(productsID, k)
	}

	builder := s.Builder().
		Select("quantity", "reserved_number", "product_id").
		From(warehouseProductTableName).
		Where(squirrel.Eq{"warehouse_id": filter.WarehouseID}).
		Where(squirrel.Eq{"product_id": productsID})

	query, values := builder.MustSql()

	return transactionHelper(ctx, s.conn, func(ctx context.Context, tx pgx.Tx) (map[uint64]bool, error) {
		rows, err := tx.Query(ctx, query, values...)
		if err != nil {
			return nil, err
		}
	
		productsReserved := make(map[uint64]uint64)
	
		for rows.Next() {
			var (
				quantity       uint64
				reservedNumber uint64
				productID      uint64
			)
			err = rows.Scan(&quantity, &reservedNumber, &productID)
			if err != nil {
				log.Printf("scan: %v", err)
				continue
			}
	
			if reservedNumber+filter.Products[productID] > quantity {
				continue
			}
	
			productsReserved[productID] = reservedNumber
		}
	
		res := helpers.NewSyncMapUint64Bool(len(filter.Products))
		workerPool := make(chan struct{}, workerPoolSize)
		wg := sync.WaitGroup{}
		wg.Add(len(productsID))
	
		for _, id := range productsID {
			workerPool <- struct{}{}
			id := id
	
			go func() {
				defer func() {
					<-workerPool
					wg.Done()
				}()
	
				v, ok := productsReserved[id]
				if !ok {
					res.Store(id, false)
					return
				}
	
				builder := s.Builder().
					Update(warehouseProductTableName).
					Set("reserved_number", v+filter.Products[id]).
					Where(squirrel.Eq{"warehouse_id": filter.WarehouseID}).
					Where(squirrel.Eq{"product_id": id})
	
				query, values = builder.MustSql()
				_, err = tx.Exec(ctx, query, values...)
				if err != nil {
					log.Printf("update: %v", err)
					res.Store(id, false)
					return
				}
	
				res.Store(id, true)
			}()
		}
	
		close(workerPool)
		wg.Wait()
	
		return res.GetData(), err
	})
}

func (s *storeImpl) ReleaseOfReserved(filter *domain.Filter) (map[uint64]bool, error) {
	ctx := context.Background()

	productsID := make([]uint64, 0, len(filter.Products))
	for k := range filter.Products {
		productsID = append(productsID, k)
	}

	builder := s.Builder().
		Select("reserved_number", "product_id").
		From(warehouseProductTableName).
		Where(squirrel.Eq{"warehouse_id": filter.WarehouseID}).
		Where(squirrel.Eq{"product_id": productsID})

	query, values := builder.MustSql()

	return transactionHelper(ctx, s.conn, func(ctx context.Context, tx pgx.Tx) (map[uint64]bool, error) {
		rows, err := tx.Query(ctx, query, values...)
		if err != nil {
			return nil, err
		}

		productsReserved := make(map[uint64]uint64)

		for rows.Next() {
			var (
				reservedNumber uint64
				productID      uint64
			)
			err = rows.Scan(&reservedNumber, &productID)
			if err != nil {
				log.Printf("scan: %v", err)
				continue
			}

			productsReserved[productID] = reservedNumber
		}

		res := helpers.NewSyncMapUint64Bool(len(filter.Products))
		workerPool := make(chan struct{}, workerPoolSize)
		wg := sync.WaitGroup{}
		wg.Add(len(productsID))

		for _, id := range productsID {
			workerPool <- struct{}{}
			id := id

			go func() {
				defer func() {
					<-workerPool
					wg.Done()
				}()

				v, ok := productsReserved[id]
				if !ok {
					res.Store(id, false)
					return
				}

				var newReservedNumber uint64
				if v > filter.Products[id] {
					newReservedNumber = v - filter.Products[id]
				}

				builder := s.Builder().
					Update(warehouseProductTableName).
					Set("reserved_number", newReservedNumber).
					Where(squirrel.Eq{"warehouse_id": filter.WarehouseID}).
					Where(squirrel.Eq{"product_id": id})

				query, values = builder.MustSql()
				_, err = tx.Exec(ctx, query, values...)
				if err != nil {
					log.Printf("update: %v", err)
					res.Store(id, false)
					return
				}

				res.Store(id, true)
			}()
		}

		close(workerPool)
		wg.Wait()

		return res.GetData(), err
	})
}

func (s *storeImpl) RemainingProducts(warehouseID *uint64) (map[uint64]uint64, error) {
	builder := s.Builder().
		Select("SUM(quantity - reserved_number) as remaining", "product_id").
		From(warehouseProductTableName).
		Where(squirrel.Eq{"warehouse_id": warehouseID}).
		GroupBy("product_id")

	query, values := builder.MustSql()

	rows, err := s.conn.Query(context.Background(), query, values...)
	if err != nil {
		return nil, err
	}

	remainingProducts := make(map[uint64]uint64)

	for rows.Next() {
		var (
			remaining uint64
			productID uint64
		)
		err = rows.Scan(&remaining, &productID)
		if err != nil {
			log.Printf("scan: %v", err)
			continue
		}

		remainingProducts[productID] = remaining
	}

	return remainingProducts, nil
}

func transactionHelper(
	ctx context.Context, 
	conn *pgx.Conn, 
	f func(context.Context, pgx.Tx) (map[uint64]bool, error),
) (map[uint64]bool, error) {
	tx, err := conn.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.RepeatableRead})
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			err := tx.Rollback(ctx)
			if err != nil {
				log.Println(err)
			}
		} else {
			err := tx.Commit(ctx)
			if err != nil {
				log.Println(err)
			}
		}
	}()
	return f(ctx, tx)
}