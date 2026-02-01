package dto

import "time"

type TransactionResponse struct {
	PublicID          string            `json:"public_id"`
	ItemResponse      ItemResponse      `json:"item"`
	WarehouseResponse WarehouseResponse `json:"warehouse"`
	Type              string            `json:"type"`
	Quantity          int               `json:"quantity"`
	CreatedAt         time.Time         `json:"created_at"`
}
