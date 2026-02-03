package routes

import (
	"github.com/Wenell09/MyStock/internal/dashboard/controller"
	"github.com/gofiber/fiber/v2"
)

func RegisterDashboardRoutes(
	api fiber.Router,
	controller controller.DashboardController,
) {
	dashboard := api.Group("/dashboard")
	dashboard.Get("/", controller.CountData)

}
