package dto

type WarehouseRequest struct {
	Name string `json:"name" validate:"required"`
}
