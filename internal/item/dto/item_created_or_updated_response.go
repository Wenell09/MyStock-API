package dto

import "time"

type ItemCreatedOrUpdatedResponse struct {
	PublicId  string    `json:"public_id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}
