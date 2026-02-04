package dto

import "time"

type WarehouseResponse struct {
	PublicId  string    `json:"public_id"`
	Name      string    `json:"name"`
	Location  string    `json:"location"`
	CreatedAt time.Time `json:"created_at"`
}
