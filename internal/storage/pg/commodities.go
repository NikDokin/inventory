package pg

import (
	"context"
	"fmt"

	"github.com/fungicibus/inventory/internal/types"
)

func (s *Adapter) GetCommodities(ctx context.Context, filters types.CommoditiesFilters) ([]*types.Commodity, error) {
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
	`
	if filters.Name != "" {
		filterName := fmt.Sprintf(` WHERE LOWER("name") LIKE LOWER('%s')`, "%"+filters.Name+"%")
		query += filterName
	}

	rows, err := s.roPool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query: %w", err)
	}

	commodities := make([]*types.Commodity, 0)
	for rows.Next() {
		var commodity types.Commodity
		err = rows.Scan(
			&commodity.Id,
			&commodity.Category,
			&commodity.Description,
			&commodity.Name,
			&commodity.Package,
			&commodity.Price,
			&commodity.Quantity,
			&commodity.Sku,
		)

		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		commodities = append(commodities, &commodity)
	}

	return commodities, nil
}
