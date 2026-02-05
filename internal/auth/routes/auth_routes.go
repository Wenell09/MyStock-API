package routes

import (
	"github.com/Wenell09/MyStock/internal/auth/controller"
	"github.com/gofiber/fiber/v2"
)

func RegisterAuthRoutes(
	api fiber.Router,
	controller controller.AuthController,
) {
	auth := api.Group("/auth")
	auth.Post("/login", controller.Login)
}
