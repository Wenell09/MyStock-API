package routes

import (
	"github.com/Wenell09/MyStock/internal/warehouse/controller"
	"github.com/gofiber/fiber/v2"
)

func RegisterWarehouseRoutes(
	api fiber.Router,
	controller controller.WarehouseController,
) {
	warehouses := api.Group("/warehouses")

	warehouses.Get("/", controller.Read)
	warehouses.Get("/:public_id", controller.ReadByPublicId)
	warehouses.Post("/", controller.Create)
	warehouses.Patch("/:public_id", controller.Update)
	warehouses.Delete("/", controller.DeleteAll)
	warehouses.Delete("/:public_id", controller.Delete)
}
