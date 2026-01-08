package routes

import (
	"github.com/Wenell09/MyStock/internal/category/controller"
	"github.com/gofiber/fiber/v2"
)

func RouteApp(app *fiber.App, categoryController controller.CategoryController) {
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Welcome to API MyStock")
	})
	api := app.Group("/api")
	api.Get("/categories", categoryController.Read)
	api.Patch("/categories/:public_id", categoryController.Update)
	api.Post("/categories", categoryController.Create)
	api.Delete("/categories", categoryController.DeleteAll)
	api.Delete("/categories/:public_id", categoryController.Delete)

}
