package service

import (
	"github.com/Wenell09/MyStock/internal/auth/dto"
	"github.com/Wenell09/MyStock/internal/models"
)

type AuthService interface {
	Login(request dto.LoginRequest) (models.User, string, error)
}
