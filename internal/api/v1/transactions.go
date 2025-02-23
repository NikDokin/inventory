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

	// TODO: check if commodity exists

	switch request.Data.Type {
	case Supply, Sale:
		// valid
	default:
		api.WriteError(w, r,
			WithStatusCode(http.StatusUnprocessableEntity),
			WithDetail(fmt.Sprintf("invalid type: %s", request.Data.Type)),
			WithSourcePointer("/data/type"),
		)
		return
	}
	createdAt, err := time.Parse(time.RFC3339, request.Data.CreatedAt)
	if err != nil {
		api.WriteError(w, r,
			WithStatusCode(http.StatusBadRequest),
			WithDetail(fmt.Sprintf("bad createdAt value: %s", request.Data.CreatedAt)),
			WithSourcePointer("/data/createdAt"),
			WithError(err),
		)
		return
	}

	transactionID := uuid.New().String()
	note := ""
	if request.Data.Note != nil {
		note = *request.Data.Note
	}
	transaction := &types.Transaction{
		ID:          transactionID,
		CommodityID: request.Data.CommodityID,
		Amount:      request.Data.Amount,
		Type:        string(request.Data.Type),
		Note:        note,
		CreatedAt:   createdAt,
		SavedAt:     time.Now().UTC(),
	}
	if err := api.storage.CreateTransaction(r.Context(), transaction); err != nil {
		api.WriteError(w, r,
			WithStatusCode(http.StatusInternalServerError),
			WithError(fmt.Errorf("failed to create transaction: %w", err)),
		)
		return
	}

	quantity := transaction.Amount
	if TransactionType(transaction.Type) == Sale {
		quantity = -transaction.Amount
	}

	if err != api.storage.UpdateCommodityQuantity(r.Context(), transaction.CommodityID, quantity) {
		api.WriteError(w, r,
			WithStatusCode(http.StatusInternalServerError),
			WithError(fmt.Errorf("failed to update commodity quantity: %w", err)),
		)
		return
	}

	response := CreateTransactionResponse{
		Data: TransactionItem{
			Id:          transaction.ID,
			CommodityID: transaction.CommodityID,
			Amount:      transaction.Amount,
			Type:        request.Data.Type,
			Note:        request.Data.Note,
			CreatedAt:   request.Data.CreatedAt,
			SavedAt:     transaction.SavedAt.Format(time.RFC3339),
		},
	}

	api.WriteJSON(w, r, response)
}
