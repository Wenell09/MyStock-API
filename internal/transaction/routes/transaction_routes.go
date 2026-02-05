package routes

import (
	"github.com/Wenell09/MyStock/internal/middleware"
	"github.com/Wenell09/MyStock/internal/transaction/controller"
	"github.com/gofiber/fiber/v2"
)

func RegisterTransactionRoutes(api fiber.Router, controller controller.TransactionController) {
	transactions := api.Group("/transactions", middleware.Auth())

	transactions.Get("/", controller.Read)
	transactions.Get("/:public_id", controller.ReadByPublicId)
	transactions.Post("/", controller.CreateOrUpdate)
}
