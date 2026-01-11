package dto

import "time"

type CategoryResponse struct {
	PublicId  string    `json:"public_id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at,omitzero"`
}
