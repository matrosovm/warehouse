package mapper

import (
	"github.com/matrosovm/warehouse/internal/pkg/domain"
	pbWarehouse "github.com/matrosovm/warehouse/pkg/api/warehouse"
)

// ReleaseOfReservedRequestToFilter ...
func ReleaseOfReservedRequestToFilter(
	req *pbWarehouse.ReleaseOfReservedRequest,
) *domain.Filter {
	return &domain.Filter{
		WarehouseID: req.WarehouseID,
		Products:    req.Products,
	}
}

// ReservationRequestToFilter ...
func ReservationRequestToFilter(
	req *pbWarehouse.ReservationRequest,
) *domain.Filter {
	return &domain.Filter{
		WarehouseID: req.WarehouseID,
		Products:    req.Products,
	}
}
