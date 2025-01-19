package mock

import (
	"context"

	"github.com/fungicibus/inventory/internal/types"
)

type mockStorage struct {
}

func New() *mockStorage {
	return &mockStorage{}
}

func (m *mockStorage) GetCommodities(ctx context.Context, filters types.CommoditiesFilters) ([]types.Commodity, error) {
	commodities := []types.Commodity{
		{
			Id:          "266b9823-9b87-4d73-a0f8-41a2b6c5e832",
			Category:    "culinary",
			Name:        "Lion's mane",
			Description: "Hericium erinaceus. The edible fruiting bodies",
			Price:       5,
			Quantity:    100,
			PackageForm: "5 dried pieces",
			Sku:         "CUL-DRY-LNM",
		},
		{
			Id:          "266b9823-9b87-4d73-a0f8-41a2b6c5e833",
			Category:    "exotic",
			Name:        "Fly agaric",
			Description: "Amanita muscaria. The edible caps",
			Price:       1,
			Quantity:    100,
			PackageForm: "1 dried piece",
			Sku:         "EXO-DRY-FAG",
		},
	}

	return commodities, nil
}
