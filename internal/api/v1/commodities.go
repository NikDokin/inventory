package v1

import (
	"fmt"
	"net/http"

	"github.com/fungicibus/inventory/internal/types"
)

// Get list of all commodities
// (GET /commodities)
func (api *API) GetCommodities(w http.ResponseWriter, r *http.Request, params GetCommoditiesParams) {
	filters := types.CommoditiesFilters{}

	if params.Name != nil {
		filters.Name = *params.Name
	}

	commodities, err := api.storage.GetCommodities(r.Context(), filters)
	if err != nil {
		api.WriteError(w, r, WithError(fmt.Errorf("failed to get commodities from storage: %w", err)))
		return
	}

	responseDataItems := make([]CommoditiesItem, 0, len(commodities))
	for _, commodity := range commodities {
		responseDataItems = append(responseDataItems, CommoditiesItem{
			Id:          commodity.Id,
			Category:    CategoryType(commodity.Category),
			Name:        commodity.Name,
			Description: commodity.Description,
			Price:       commodity.Price,
			Quantity:    commodity.Quantity,
			Package:     commodity.Package,
			Sku:         commodity.Sku,
		})
	}

	response := GetCommoditiesResponse{
		Data: responseDataItems,
	}

	api.WriteJSON(w, r, response)
}
