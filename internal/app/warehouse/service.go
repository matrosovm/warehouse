package warehouse

import (
	"github.com/matrosovm/warehouse/internal/pkg/database"
	pbWarehouse "github.com/matrosovm/warehouse/pkg/api/warehouse"
)

// Service ...
type Service struct {
	store database.Store
}

// NewService ...
func NewService(store database.Store) pbWarehouse.Server {
	return &Service{
		store: store,
	}
}
