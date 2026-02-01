package dto

type TransactionRequest struct {
	ItemPublicId      string `json:"item_public_id" validate:"required"`
	WarehousePublicId string `json:"warehouse_public_id" validate:"required"`
	Type              string `json:"type" validate:"required"`
	Quantity          int    `json:"quantity" validate:"required"`
}
