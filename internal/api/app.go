package api

import (
	categoryController "github.com/Wenell09/MyStock/internal/category/controller"
	warehouseController "github.com/Wenell09/MyStock/internal/warehouse/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func NewFiberApp(
	categoryCtrl categoryController.CategoryController,
	warehouseCtrl warehouseController.WarehouseController,
) *fiber.App {

	app := fiber.New()

	// Global middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
	}))

	RegisterRoutes(app, categoryCtrl, warehouseCtrl)

	return app
}
