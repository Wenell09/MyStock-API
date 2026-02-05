package dto

type LoginResponse struct {
	User  UserResponse `json:"user"`
	Token string       `json:"token"`
}
