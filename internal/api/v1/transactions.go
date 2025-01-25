package v1

import (
	"fmt"
	"net/http"
	"time"

	"github.com/fungicibus/inventory/internal/types"
	"github.com/google/uuid"
)

// Create transaction
// (POST /transactions)
func (api *API) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var request CreateTransactionRequest
	if err := api.ReadJSON(w, r, &request); err != nil {
		api.WriteError(w, r, WithStatusCode(http.StatusBadRequest), WithError(err))
		return
	}
	if request.Type != Supply && request.Type != Sale {
		api.WriteError(w, r, WithStatusCode(http.StatusBadRequest), WithDetail("bad transaction type"))
		return
	}
	createdAt, err := time.Parse(time.RFC3339, request.CreatedAt)
	if err != nil {
		api.WriteError(w, r, WithStatusCode(http.StatusBadRequest), WithDetail("bad createdAt value"), WithError(err))
		return
	}

	transactionID := uuid.New().String()
	note := ""
	if request.Note != nil {
		note = *request.Note
	}
	transaction := &types.Transaction{
		ID:          transactionID,
		CommodityID: request.CommodityID,
		Amount:      request.Amount,
		Type:        string(request.Type),
		Note:        note,
		CreatedAt:   createdAt,
	}
	if err := api.storage.CreateTransaction(r.Context(), transaction); err != nil {
		api.WriteError(w, r,
			WithStatusCode(http.StatusInternalServerError),
			WithError(fmt.Errorf("failed to create transaction: %w", err)),
		)
		return
	}

	response := CreateTransactionResponse{
		Id:          transaction.ID,
		CommodityID: request.CommodityID,
		Amount:      request.Amount,
		Type:        request.Type,
		Note:        request.Note,
		CreatedAt:   request.CreatedAt,
	}

	api.WriteJSON(w, r, response)
}
