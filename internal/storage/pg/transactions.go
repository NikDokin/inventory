package pg

import (
	"context"
	"fmt"

	"github.com/fungicibus/inventory/internal/types"
)

func (pg *Adapter) CreateTransaction(ctx context.Context, tx types.Tx, transaction *types.Transaction) error {
	const query = `
	INSERT INTO transactions (id, commodity_id, amount, "type", created_at, note)
	VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := tx.Exec(ctx, query,
		transaction.ID,
		transaction.CommodityID,
		transaction.Amount,
		transaction.Type,
		transaction.CreatedAt,
		transaction.Note,
	)
	if err != nil {
		return fmt.Errorf("failed to exec: %w", err)
	}

	return nil
}
