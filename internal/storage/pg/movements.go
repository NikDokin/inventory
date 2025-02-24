package pg

import (
	"context"
	"fmt"

	"github.com/fungicibus/inventory/internal/types"
)

func (pg *Adapter) CreateMovement(ctx context.Context, tx types.Tx, movement *types.Movement) error {
	const query = `
	INSERT INTO movements (id, commodity_id, amount, "type", created_at, note)
	VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := tx.Exec(ctx, query,
		movement.ID,
		movement.CommodityID,
		movement.Amount,
		movement.Type,
		movement.CreatedAt,
		movement.Note,
	)
	if err != nil {
		return fmt.Errorf("failed to exec: %w", err)
	}

	return nil
}
