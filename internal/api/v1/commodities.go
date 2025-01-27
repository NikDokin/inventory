package v1

import (
	"fmt"
	"net/http"

	"github.com/fungicibus/inventory/internal/types"
	"github.com/google/uuid"
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

	responseDataItems := make([]CommodityItem, 0, len(commodities))
	for _, commodity := range commodities {
		responseDataItems = append(responseDataItems, CommodityItem{
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

// Get single commodity by id
// (GET /commdities/{commodityID})
func (api *API) GetCommodity(w http.ResponseWriter, r *http.Request, commodityID string) {
	commodities, err := api.storage.GetCommodities(r.Context(), types.CommoditiesFilters{
		CommodityID: commodityID,
	})
	if err != nil {
		api.WriteError(w, r, WithError(fmt.Errorf("failed to get commodities from storage: %w", err)))
		return
	}

	if len(commodities) == 0 {
		api.WriteError(w, r, WithStatusCode(http.StatusNotFound))
		return
	}

	commodity := commodities[0]
	response := GetCommodityResponse{
		Data: CommodityItem{
			Id:          commodity.Id,
			Category:    CategoryType(commodity.Category),
			Name:        commodity.Name,
			Description: commodity.Description,
			Price:       commodity.Price,
			Quantity:    commodity.Quantity,
			Package:     commodity.Package,
			Sku:         commodity.Sku,
		},
	}

	api.WriteJSON(w, r, response)
}

// Create commodity
// (POST /commodities)
func (api *API) CreateCommodity(w http.ResponseWriter, r *http.Request) {
	var request CreateCommodityRequest
	if err := api.ReadJSON(w, r, &request); err != nil {
		api.WriteError(w, r, WithError(fmt.Errorf("failed to read request body: %w", err)))
		return
	}

	// TODO: validate

	commodityID := uuid.New().String()
	sku := "" // TODO: implement sku generator
	commodity := types.Commodity{
		Id:          commodityID,
		Category:    string(request.Data.Category),
		Name:        request.Data.Name,
		Description: request.Data.Description,
		Price:       request.Data.Price,
		Quantity:    request.Data.Quantity,
		Package:     request.Data.Package,
		Sku:         sku,
	}

	if err := api.storage.CreateCommodity(r.Context(), &commodity); err != nil {
		api.WriteError(w, r, WithError(fmt.Errorf("failed to create commodity in storage: %w", err)))
		return
	}

	response := CreateCommodityResponse{
		Data: CommodityItem{
			Id:          commodity.Id,
			Category:    CategoryType(commodity.Category),
			Name:        commodity.Name,
			Description: commodity.Description,
			Price:       commodity.Price,
			Quantity:    commodity.Quantity,
			Package:     commodity.Package,
			Sku:         commodity.Sku,
		},
	}

	api.WriteJSON(w, r, response)
}
