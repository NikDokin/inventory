package v1

import (
	"encoding/json"
	"net/http"
)

func (api *API) WriteJSON(w http.ResponseWriter, r *http.Request, response any) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		api.logger.Error().Err(err).Msgf("failed to encode response body")
		api.WriteError(w, r, err)
	}
}

func (api *API) WriteError(w http.ResponseWriter, r *http.Request, err error) {
	// what message provide here and how?
}
