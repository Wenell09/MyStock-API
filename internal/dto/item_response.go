package dto

import "time"

type ItemResponse struct {
	PublicId          string              `json:"public_id"`
	Name              string              `json:"name"`
	CategoryResponse  CategoryResponse    `json:"category"`
	SupplierResponse  SupplierResponse    `json:"supplier"`
	WarehouseResponse []WarehouseResponse `json:"warehouses"`
	CreatedAt         time.Time           `json:"created_at,omitzero"`
}
