package api

import (
	"github.com/Wenell09/MyStock/internal/category/controller"
	"github.com/Wenell09/MyStock/internal/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func NewFiberApp(
	categoryController controller.CategoryController,
) *fiber.App {

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
	}))

	routes.RouteApp(app, categoryController)

	return app
}
