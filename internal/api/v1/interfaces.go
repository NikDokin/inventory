package v1

import (
	"context"

	"github.com/fungicibus/inventory/internal/types"
)

type Storage interface {
	BeginTx(ctx context.Context) (types.Tx, error)

	GetCommodities(ctx context.Context, filters types.CommoditiesFilters) ([]*types.Commodity, error)
	CreateCommodity(ctx context.Context, commodity *types.Commodity) error
	UpdateCommodityQuantity(ctx context.Context, tx types.Tx, commodityID string, quantity int) error

	CreateTransaction(ctx context.Context, tx types.Tx, transaction *types.Transaction) error
}
