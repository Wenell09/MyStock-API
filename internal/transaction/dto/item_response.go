package dto

type ItemResponse struct {
	PublicId          string              `json:"public_id"`
	Name              string              `json:"name"`
	CategoryResponse  CategoryResponse    `json:"category"`
	SupplierResponse  SupplierResponse    `json:"supplier"`
}
