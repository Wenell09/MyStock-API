package controller

import (
	"github.com/Wenell09/MyStock/internal/auth/dto"
	"github.com/Wenell09/MyStock/internal/auth/service"
	"github.com/Wenell09/MyStock/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type AuthControllerImpl struct {
	AuthService service.AuthService
	Logger      *logrus.Logger
}

func NewAuthController(authService service.AuthService, logger *logrus.Logger) AuthController {
	return &AuthControllerImpl{
		AuthService: authService,
		Logger:      logger,
	}
}

// login implements [AuthController].
func (a *AuthControllerImpl) Login(ctx *fiber.Ctx) error {
	var request dto.LoginRequest
	if err := ctx.BodyParser(&request); err != nil {
		a.Logger.Error(err.Error())
		return utils.NewHandleError(ctx, utils.ValidationError{
			Msg: "Invalid request body",
		})
	}
	responseUser, ResponseToken, err := a.AuthService.Login(request)
	if err != nil {
		a.Logger.Error(err.Error())
		return utils.NewHandleError(ctx, err)
	}
	data := dto.LoginResponse{
		User: dto.UserResponse{
			PublicId: responseUser.ID,
			Email:    responseUser.Email,
			Name:     responseUser.Name,
			Role:     responseUser.Role,
		},
		Token: ResponseToken,
	}
	a.Logger.WithFields(logrus.Fields{
		"status":   fiber.StatusCreated,
		"request":  request,
		"response": data,
	}).Info("Login Success!")
	return ctx.Status(fiber.StatusCreated).JSON(
		utils.NewResponseSuccess(fiber.StatusCreated, "Login Success!", data),
	)
}
