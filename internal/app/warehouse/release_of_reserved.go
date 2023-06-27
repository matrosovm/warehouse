package warehouse

import (
	"github.com/matrosovm/warehouse/internal/app/warehouse/mapper"
	pbWarehouse "github.com/matrosovm/warehouse/pkg/api/warehouse"
)

// ReleaseOfReserved ...
func (s *Service) ReleaseOfReserved(
	req *pbWarehouse.ReleaseOfReservedRequest,
	resp *pbWarehouse.ReleaseOfReservedResponse,
) error {
	status, err := s.store.ReleaseOfReserved(
		mapper.ReleaseOfReservedRequestToFilter(req),
	)
	if err != nil {
		return err
	}

	resp.Status = status
	return nil
}
