package mock

import (
	"context"

	"github.com/fungicibus/inventory/internal/types"
)

func (m *mockStorage) CreateMovement(ctx context.Context, movement *types.Movement) error {

	return nil
}
