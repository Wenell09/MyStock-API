package api

import (
	categoryController "github.com/Wenell09/MyStock/internal/category/controller"
	categoryRoutes "github.com/Wenell09/MyStock/internal/category/routes"
	warehouseController "github.com/Wenell09/MyStock/internal/warehouse/controller"
	warehouseRoutes "github.com/Wenell09/MyStock/internal/warehouse/routes"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(
	app *fiber.App,
	categoryCtrl categoryController.CategoryController,
	warehouseCtrl warehouseController.WarehouseController,
) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to API MyStock")
	})

	api := app.Group("/api")

	categoryRoutes.RegisterCategoryRoutes(api, categoryCtrl)
	warehouseRoutes.RegisterWarehouseRoutes(api, warehouseCtrl)
}
