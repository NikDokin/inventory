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
	query := `
		INSERT INTO commodities (
			id,
			"name",
			sku,
			description,
			category,
			quantity,
			"package",
			price
		) VALUES (
			$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7,
			$8
		)
	`
	_, err := pg.rwPool.Exec(ctx, query,
		commodity.Id,
		commodity.Name,
		commodity.Sku,
		commodity.Description,
		commodity.Category,
		commodity.Quantity,
		commodity.Package,
		commodity.Price,
	)
	if err != nil {
		return fmt.Errorf("failed to exec: %w", err)
	}

	return nil
}
