package repository

import "github.com/Wenell09/MyStock/internal/models"

type AuthRepository interface {
	FindUserByEmail(email string) (models.User, error)
}
