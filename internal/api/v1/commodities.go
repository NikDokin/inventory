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
		api.WriteError(w, r, WithError(err), WithDetail("failed to get commodities"))
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

// Add to commodity quantity
// (POST /commodities/{commodityID}/quantity/add)
func (api *API) AddCommodityQuantity(w http.ResponseWriter, r *http.Request, commodityID string) {
	if commodityID == "" {
		msg := "commodityID must not be empty"
		api.WriteError(w, r, WithDetail(msg), WithStatusCode(http.StatusBadRequest))
		return
	}
}
