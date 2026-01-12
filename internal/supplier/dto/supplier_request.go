package dto

type SupplierRequest struct {
	Name string `json:"name" validate:"required"`
}
