package dto

import "time"

type TransactionCreatedOrUpdatedResponse struct {
	PublicId          string    `json:"public_id"`
	ItemPublicId      string    `json:"item_public_id"`
	WarehousePublicId string    `json:"warehouse_public_id"`
	Type              string    `json:"type"`
	Quantity          int       `json:"quantity"`
	CreatedAt         time.Time `json:"created_at"`
}
