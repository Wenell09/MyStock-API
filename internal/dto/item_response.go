package dto

import "time"

type ItemResponse struct {
	PublicId          string              `json:"public_id"`
	Name              string              `json:"name"`
	CategoryResponse  CategoryResponse    `json:"category,omitzero"`
	SupplierResponse  SupplierResponse    `json:"supplier,omitzero"`
	WarehouseResponse []WarehouseResponse `json:"warehouses,omitzero"`
	CreatedAt         time.Time           `json:"created_at,omitzero"`
}
