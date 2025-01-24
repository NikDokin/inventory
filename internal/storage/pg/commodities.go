package pg

import (
	"context"
	"fmt"

	"github.com/fungicibus/inventory/internal/types"
)

func (s *Adapter) GetCommodities(ctx context.Context, filters types.CommoditiesFilters) ([]*types.Commodity, error) {
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
		WHERE ($1 = '' OR LOWER("name") LIKE LOWER($2))
	`

	commodities := make([]*types.Commodity, 0, 10)

	rows, err := s.roPool.Query(ctx, query,
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
