package v1

import (
	"context"
	"errors"

	"github.com/fungicibus/inventory/internal/types"
)

var (
	ErrNoCommodityFound = errors.New("no commodity found")
)

type Storage interface {
	GetCommodities(ctx context.Context, filters types.CommoditiesFilters) ([]types.Commodity, error)
	AddCommodityQuantity(ctx context.Context, commodityID string, amout int) (types.Commodity, error)
}
