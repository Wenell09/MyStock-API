package dto

type UserResponse struct {
	PublicId string `json:"public_id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Role     string `json:"role"`
}
