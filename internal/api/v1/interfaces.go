package v1

import (
	"context"

	"github.com/fungicibus/inventory/internal/types"
)

type Storage interface {
	GetCommodities(ctx context.Context, filters types.CommoditiesFilters) ([]*types.Commodity, error)
	CreateCommodity(ctx context.Context, commodity *types.Commodity) error

	CreateTransaction(ctx context.Context, transaction *types.Transaction) error
}
