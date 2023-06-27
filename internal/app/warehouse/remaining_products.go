package warehouse

import (
	pbWarehouse "github.com/matrosovm/warehouse/pkg/api/warehouse"
)

// RemainingProducts ...
func (s *Service) RemainingProducts(
	req *pbWarehouse.RemainingProductsRequest,
	resp *pbWarehouse.RemainingProductsResponse,
) error {
	products, err := s.store.RemainingProducts(&req.WarehouseID)
	if err != nil {
		return err
	}

	resp.Products = products
	return nil
}
