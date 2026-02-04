package dto

type WarehouseWithItemsResponse struct {
	PublicId   string         `json:"public_id"`
	Name       string         `json:"name"`
	Location   string         `json:"location"`
	Items      []ItemResponse `json:"items"`
	TotalStock int            `json:"total_stock"`
}
