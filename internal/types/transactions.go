package types

import "time"

type Transaction struct {
	ID          string    `json:"id"`
	CommodityID string    `json:"commodityID"`
	Amount      int       `json:"amount"`
	Type        string    `json:"type"`
	Note        string    `json:"note"`
	CreatedAt   time.Time `json:"createdAt"`
	SavedAt     time.Time `json:"savedAt"`
}
