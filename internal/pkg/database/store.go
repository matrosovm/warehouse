package database

import (
	"context"
	"sync"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/matrosovm/warehouse/internal/pkg/domain"
	"google.golang.org/appengine/log"
)

const (
	productTableName          = "product"
	warehouseTableName        = "warehouse"
	warehouseProductTableName = "warehouse_product"
)

// Store ...
type Store interface {
	ReserveProducts(*domain.Filter) (map[uint64]bool, error)
	ReleaseOfReserved(*domain.Filter) (map[uint64]bool, error)
	RemainingProducts(*uint64) (map[uint64]uint64, error)
	Close(context.Context)
}

type storeImpl struct {
	conn *pgx.Conn
	mu   sync.RWMutex
}

// NewStore ...
func NewStore(conn *pgx.Conn) Store {
	return &storeImpl{
		conn: conn,
	}
}

// Builder return squirrel SQL Builder object
func (s *storeImpl) Builder() sq.StatementBuilderType {
	return sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
}

func (s *storeImpl) Close(ctx context.Context) {
	if err := s.conn.Close(ctx); err != nil {
		log.Errorf(ctx, "conn close failed: %v", err.Error())
	}
}
