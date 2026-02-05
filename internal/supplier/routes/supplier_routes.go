package routes

import (
	"github.com/Wenell09/MyStock/internal/middleware"
	"github.com/Wenell09/MyStock/internal/supplier/controller"
	"github.com/gofiber/fiber/v2"
)

func RegisterSupplierRoutes(
	api fiber.Router,
	controller controller.SupplierController,
) {
	suppliers := api.Group("/suppliers", middleware.Auth())

	suppliers.Get("/", controller.Read)
	suppliers.Get("/:public_id", controller.ReadByPublicId)
	suppliers.Get("/:public_id/items", controller.ReadBySupplierPublicId)
	suppliers.Post("/", controller.Create)
	suppliers.Patch("/:public_id", controller.Update)
	suppliers.Delete("/", controller.DeleteAll)
	suppliers.Delete("/:public_id", controller.Delete)
}
