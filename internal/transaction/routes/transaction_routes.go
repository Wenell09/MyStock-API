package routes

import (
	"github.com/Wenell09/MyStock/internal/transaction/controller"
	"github.com/gofiber/fiber/v2"
)

func RegisterTransactionRoutes(api fiber.Router, controller controller.TransactionController) {
	transactions := api.Group("/transactions")

	transactions.Get("/", controller.Read)
	transactions.Get("/:public_id", controller.ReadByPublicId)
	transactions.Post("/", controller.CreateOrUpdate)
}
