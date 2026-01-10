package dto

type ItemRequest struct {
	Name             string `json:"name" validate:"required"`
	CategoryPublicId string `json:"category_public_id" validate:"required"`
	SupplierPublicId string `json:"supplier_public_id" validate:"required"`
}
