package mock

import (
	"context"
	"fmt"
	"strings"

	v1 "github.com/fungicibus/inventory/internal/api/v1"
	"github.com/fungicibus/inventory/internal/types"
)

type mockStorage struct {
}

func New() *mockStorage {
	return &mockStorage{}
}

func (m *mockStorage) GetCommodities(ctx context.Context, filters types.CommoditiesFilters) ([]types.Commodity, error) {
	commodities := []types.Commodity{}

	nameQuery := strings.ToLower(filters.Name)
	if nameQuery == "" {
		commodities = []types.Commodity{
			{
				Id:          "266b9823-9b87-4d73-a0f8-41a2b6c5e832",
				Category:    "culinary",
				Name:        "Lion's mane",
				Description: "Hericium erinaceus. The edible fruiting bodies",
				Price:       5,
				Quantity:    100,
				Package:     "5 dried pieces",
				Sku:         "CUL-DRY-LNM",
			},
			{
				Id:          "266b9823-9b87-4d73-a0f8-41a2b6c5e833",
				Category:    "exotic",
				Name:        "Fly agaric",
				Description: "Amanita muscaria. The edible caps",
				Price:       1,
				Quantity:    8,
				Package:     "1 dried piece",
				Sku:         "EXO-DRY-FAG",
			},
			{
				Id:          "266b9823-9b87-4d73-a0f8-41a2b6c5e834",
				Category:    "exotic",
				Name:        "Fly agaric",
				Description: "Amanita muscaria. The powder from caps in capsules",
				Price:       5,
				Quantity:    10,
				Package:     "60 capsules",
				Sku:         "EXO-CAP-FAG",
			},
		}
	} else if strings.Contains("amanita muscaria", nameQuery) || strings.Contains("fly agaric", nameQuery) {
		commodities = []types.Commodity{
			{
				Id:          "266b9823-9b87-4d73-a0f8-41a2b6c5e833",
				Category:    "exotic",
				Name:        "Fly agaric",
				Description: "Amanita muscaria. The edible caps",
				Price:       1,
				Quantity:    8,
				Package:     "1 dried piece",
				Sku:         "EXO-DRY-FAG",
			},
			{
				Id:          "266b9823-9b87-4d73-a0f8-41a2b6c5e834",
				Category:    "exotic",
				Name:        "Fly agaric",
				Description: "Amanita muscaria. The powder from caps in capsules",
				Price:       5,
				Quantity:    10,
				Package:     "60 capsules",
				Sku:         "EXO-CAP-FAG",
			},
		}
	} else if nameQuery == "error" {
		return commodities, fmt.Errorf("mock error")
	}

	return commodities, nil
}

func (m *mockStorage) AddCommodityQuantity(ctx context.Context, commodityID string, amout int) (types.Commodity, error) {
	if commodityID == "266b9823-9b87-4d73-a0f8-41a2b6c5e833" {
		return types.Commodity{
			Id:          "266b9823-9b87-4d73-a0f8-41a2b6c5e833",
			Category:    "exotic",
			Name:        "Fly agaric",
			Description: "Amanita muscaria. The edible caps",
			Price:       1,
			Quantity:    8 + amout,
			Package:     "1 dried piece",
			Sku:         "EXO-DRY-FAG",
		}, nil
	}

	return types.Commodity{}, v1.ErrNoCommodityFound
}
