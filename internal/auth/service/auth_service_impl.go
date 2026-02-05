package service

import (
	"errors"

	"github.com/Wenell09/MyStock/internal/auth/dto"
	"github.com/Wenell09/MyStock/internal/auth/provider"
	"github.com/Wenell09/MyStock/internal/auth/repository"
	"github.com/Wenell09/MyStock/internal/models"
	"github.com/Wenell09/MyStock/internal/utils"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type AuthServiceImpl struct {
	AuthRepository repository.AuthRepository
	AuthProvider   provider.AuthProvider
	Validate       *validator.Validate
}

func NewAuthService(authRepository repository.AuthRepository, authProvider provider.AuthProvider, validate *validator.Validate) AuthService {
	return &AuthServiceImpl{
		AuthRepository: authRepository,
		AuthProvider:   authProvider,
		Validate:       validate,
	}
}

// Login implements [AuthService].
func (a *AuthServiceImpl) Login(request dto.LoginRequest) (models.User, string, error) {
	if err := a.Validate.Struct(&request); err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			return models.User{}, "", utils.NewFieldError(ve)
		}
		return models.User{}, "", utils.ValidationError{Msg: err.Error()}
	}
	token, err := a.AuthProvider.Login(request.Email, request.Password)
	if err != nil {
		return models.User{}, "", err
	}
	response, err := a.AuthRepository.FindUserByEmail(request.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.User{}, "", utils.NotFoundError{Msg: "User Not Found!"}
		}
		return models.User{}, "", err
	}
	return response, token, nil
}
