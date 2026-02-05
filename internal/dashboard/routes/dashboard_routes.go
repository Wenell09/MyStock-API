package routes

import (
	"github.com/Wenell09/MyStock/internal/dashboard/controller"
	"github.com/Wenell09/MyStock/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func RegisterDashboardRoutes(
	api fiber.Router,
	controller controller.DashboardController,
) {
	dashboard := api.Group("/dashboard", middleware.Auth())
	dashboard.Get("/", controller.CountData)

}
