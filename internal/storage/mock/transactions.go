package mock

import (
	"context"

	"github.com/fungicibus/inventory/internal/types"
)

func (m *mockStorage) CreateTransaction(ctx context.Context, transaction *types.Transaction) error {

	return nil
}
