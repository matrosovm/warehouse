package domain

// WarehouseTable ...
type WarehouseTable struct {
	ID           uint64 `db:"id"`
	Name         string `db:"name"`
	Availability bool   `db:"availability"`
}

// WarehouseProductTable ...
type WarehouseProductTable struct {
	ID             uint64 `db:"id"`
	Quantity       uint64 `db:"quantity"`
	ReservedNumber uint64 `db:"reserved_number"`
	WarehouseID    uint64 `db:"warehouse_id"`
	ProductID      uint64 `db:"product_id"`
}

// Filter ...
type Filter struct {
	WarehouseID uint64            `db:"warehouse_id"`
	Products    map[uint64]uint64 `db:"products"`
}
