package v1

import (
	"errors"
	"net/http"
)

// Get list of all commodities
// (GET /commodities)
func (api *API) GetCommodities(w http.ResponseWriter, r *http.Request) {
	if 1 != 2 {
		err := errors.New("some error")
		api.WriteError(w, r, err)
	}


	response := GetCommoditiesResponse{
		Data: []GetCommoditiesResponseDataItem{{
			Attributes: CommoditiesItem{
				Category:    Culinary,
				Name:        "Lion's mane",
				Description: "Hericium erinaceus. The edible fruiting bodies",
				Price:       5,
				Quantity:    100,
				PackageForm: "5 dried pieces",
				Sku:         "CUL-DRY-LNM",
			},
			Id:   "266b9823-9b87-4d73-a0f8-41a2b6c5e832",
			Type: Commodities,
		}},
	}
	_ = response
}
