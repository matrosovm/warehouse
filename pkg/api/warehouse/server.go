package warehouse

// Server ...
type Server interface {
	Reservation(*ReservationRequest, *ReservationResponse) error
	ReleaseOfReserved(*ReleaseOfReservedRequest, *ReleaseOfReservedResponse) error
	RemainingProducts(*RemainingProductsRequest, *RemainingProductsResponse) error
}

// ReservationRequest ...
type ReservationRequest struct {
	WarehouseID uint64            `json:"warehouse_id"`
	Products    map[uint64]uint64 `json:"products"`
}

// ReservationResponse ...
type ReservationResponse struct {
	Status map[uint64]bool `json:"status"`
}

// ReleaseOfReservedRequest ...
type ReleaseOfReservedRequest struct {
	WarehouseID uint64            `json:"warehouse_id"`
	Products    map[uint64]uint64 `json:"products"`
}

// ReleaseOfReservedResponse ...
type ReleaseOfReservedResponse struct {
	Status map[uint64]bool `json:"status"`
}

// RemainingProductsRequest ...
type RemainingProductsRequest struct {
	WarehouseID uint64 `json:"warehouse_id"`
}

// RemainingProductsResponse ...
type RemainingProductsResponse struct {
	Products map[uint64]uint64 `json:"products"`
}
