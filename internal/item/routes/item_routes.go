package routes

import (
	"github.com/Wenell09/MyStock/internal/item/controller"
	"github.com/Wenell09/MyStock/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func RegisterItemRoutes(
	api fiber.Router,
	controller controller.ItemController,
) {
	items := api.Group("/items", middleware.Auth())

	items.Get("/", controller.Read)
	items.Get("/:public_id", controller.ReadByPublicId)
	items.Post("/", controller.Create)
	items.Patch("/:public_id", controller.Update)
	items.Delete("/", controller.DeleteAll)
	items.Delete("/:public_id", controller.Delete)
}
