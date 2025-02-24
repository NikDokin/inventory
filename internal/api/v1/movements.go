package v1

import (
	"fmt"
	"net/http"
	"time"

	"github.com/fungicibus/inventory/internal/types"
	"github.com/google/uuid"
)

// Create movement
// (POST /movements)
func (api *API) CreateMovement(w http.ResponseWriter, r *http.Request) {
	var request CreateMovementResponse
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

	tx, err := api.storage.BeginTx(r.Context())
	if err != nil {
		api.WriteError(w, r,
			WithStatusCode(http.StatusInternalServerError),
			WithError(fmt.Errorf("failed to begin storage movement: %w", err)),
		)
		return
	}

	movementID := uuid.New().String()
	note := ""
	if request.Data.Note != nil {
		note = *request.Data.Note
	}
	movement := &types.Movement{
		ID:          movementID,
		CommodityID: request.Data.CommodityID,
		Amount:      request.Data.Amount,
		Type:        string(request.Data.Type),
		Note:        note,
		CreatedAt:   createdAt,
		SavedAt:     time.Now().UTC(),
	}
	if err := api.storage.CreateMovement(r.Context(), tx, movement); err != nil {
		api.WriteError(w, r,
			WithStatusCode(http.StatusInternalServerError),
			WithError(fmt.Errorf("failed to create movement: %w", err)),
		)
		return
	}

	quantity := movement.Amount
	if MovementType(movement.Type) == Sale {
		quantity = -movement.Amount
	}

	if err != api.storage.UpdateCommodityQuantity(r.Context(), tx, movement.CommodityID, quantity) {
		api.WriteError(w, r,
			WithStatusCode(http.StatusInternalServerError),
			WithError(fmt.Errorf("failed to update commodity quantity: %w", err)),
		)
		return
	}

	if err = tx.Commit(r.Context()); err != nil {
		api.WriteError(w, r,
			WithStatusCode(http.StatusInternalServerError),
			WithError(fmt.Errorf("failed to commit storage movement: %w", err)),
		)
		return
	}

	response := CreateMovementResponse{
		Data: MovementItem{
			Id:          movement.ID,
			CommodityID: movement.CommodityID,
			Amount:      movement.Amount,
			Type:        request.Data.Type,
			Note:        request.Data.Note,
			CreatedAt:   request.Data.CreatedAt,
			SavedAt:     movement.SavedAt.Format(time.RFC3339),
		},
	}

	api.WriteJSON(w, r, response)
}
