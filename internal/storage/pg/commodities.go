package pg

import (
	"context"
	"fmt"

	"github.com/fungicibus/inventory/internal/types"
)

func (pg *Adapter) GetCommodities(ctx context.Context, filters types.CommoditiesFilters) ([]*types.Commodity, error) {
	// TODO: use pagination
	query := `
		SELECT
			id,
			category,
			description,
			"name",
			package,
			price,
			quantity,
			sku
		FROM commodities
		WHERE 
			1=1
			AND ($1 = '' OR id = $1)
			AND ($2 = '' OR LOWER("name") LIKE LOWER($3))
	`

	commodities := make([]*types.Commodity, 0)

	rows, err := pg.roPool.Query(ctx, query,
		filters.CommodityID,
		filters.Name,
		"%"+filters.Name+"%",
	)
	if err != nil {
		return nil, fmt.Errorf("failed to query commodities: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		commodity := &types.Commodity{}
		if err := rows.Scan(
			&commodity.Id,
			&commodity.Category,
			&commodity.Description,
			&commodity.Name,
			&commodity.Package,
			&commodity.Price,
			&commodity.Quantity,
			&commodity.Sku,
		); err != nil {
			return nil, fmt.Errorf("failed to scan commodity: %w", err)
		}
		commodities = append(commodities, commodity)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration: %w", err)
	}

	return commodities, nil
}

func (pg *Adapter) CreateCommodity(ctx context.Context, commodity *types.Commodity) error {
	// TODO: implement
	return nil
}
