package dto

type CategoryWithItemsResponse struct {
	PublicId string         `json:"public_id"`
	Name     string         `json:"name"`
	Items    []ItemResponse `json:"items"`
}
