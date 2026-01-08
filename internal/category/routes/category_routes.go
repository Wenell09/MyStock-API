package routes

import (
	"github.com/Wenell09/MyStock/internal/category/controller"
	"github.com/gofiber/fiber/v2"
)

func RegisterCategoryRoutes(
	api fiber.Router,
	controller controller.CategoryController,
) {
	categories := api.Group("/categories")

	categories.Get("/", controller.Read)
	categories.Post("/", controller.Create)
	categories.Patch("/:public_id", controller.Update)
	categories.Delete("/", controller.DeleteAll)
	categories.Delete("/:public_id", controller.Delete)
}
