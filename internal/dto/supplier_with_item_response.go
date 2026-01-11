package dto

type SupplierWithItemResponse struct {
	PublicID string         `json:"public_id"`
	Name     string         `json:"name"`
	Items    []ItemResponse `json:"items"`
}
