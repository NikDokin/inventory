package v1

import (
	"net/http"

	"github.com/fungicibus/inventory/internal/types"
)

// Get list of all commodities
// (GET /commodities)
func (api *API) GetCommodities(w http.ResponseWriter, r *http.Request, params GetCommoditiesParams) {
	filters := types.CommoditiesFilters{}

	if params.FilterCommodityName != nil {
		filters.Name = *params.FilterCommodityName
	}

	commodities, err := api.storage.GetCommodities(r.Context(), filters)
	if err != nil {
		api.WriteError(w, r, err, WithDetail("failed to get commodities"))
		return
	}

	responseDataItems := make([]GetCommoditiesResponseDataItem, 0, len(commodities))
	for _, commodity := range commodities {
		responseDataItems = append(responseDataItems, GetCommoditiesResponseDataItem{
			Attributes: CommoditiesItem{
				Category:    CategoryType(commodity.Category),
				Name:        commodity.Name,
				Description: commodity.Description,
				Price:       commodity.Price,
				Quantity:    commodity.Quantity,
				PackageForm: commodity.PackageForm,
				Sku:         commodity.Sku,
			},
			Id:   commodity.Id,
			Type: Commodities,
		})
	}

	response := GetCommoditiesResponse{
		Data: responseDataItems,
	}

	api.WriteJSON(w, r, response)
}
