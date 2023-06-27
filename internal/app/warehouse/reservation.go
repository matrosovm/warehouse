package warehouse

import (
	"github.com/matrosovm/warehouse/internal/app/warehouse/mapper"
	pbWarehouse "github.com/matrosovm/warehouse/pkg/api/warehouse"
)

// Reservation ...
func (s *Service) Reservation(
	req *pbWarehouse.ReservationRequest,
	resp *pbWarehouse.ReservationResponse,
) error {
	status, err := s.store.ReserveProducts(
		mapper.ReservationRequestToFilter(req),
	)
	if err != nil {
		return err
	}

	resp.Status = status
	return nil
}
