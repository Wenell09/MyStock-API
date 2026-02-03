package api

import (
	categoryController "github.com/Wenell09/MyStock/internal/category/controller"
	categoryRoutes "github.com/Wenell09/MyStock/internal/category/routes"
	dashboardController "github.com/Wenell09/MyStock/internal/dashboard/controller"
	dashboardRoutes "github.com/Wenell09/MyStock/internal/dashboard/routes"
	itemController "github.com/Wenell09/MyStock/internal/item/controller"
	itemRoutes "github.com/Wenell09/MyStock/internal/item/routes"
	supplierController "github.com/Wenell09/MyStock/internal/supplier/controller"
	SupplierRoutes "github.com/Wenell09/MyStock/internal/supplier/routes"
	transactionController "github.com/Wenell09/MyStock/internal/transaction/controller"
	transactionRoutes "github.com/Wenell09/MyStock/internal/transaction/routes"
	warehouseController "github.com/Wenell09/MyStock/internal/warehouse/controller"
	warehouseRoutes "github.com/Wenell09/MyStock/internal/warehouse/routes"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(
	app *fiber.App,
	categoryCtrl categoryController.CategoryController,
	warehouseCtrl warehouseController.WarehouseController,
	supplierCtrl supplierController.SupplierController,
	itemCtrl itemController.ItemController,
	transactionCtrl transactionController.TransactionController,
	dashboardCtrl dashboardController.DashboardController,
) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to API MyStock")
	})

	api := app.Group("/api")

	categoryRoutes.RegisterCategoryRoutes(api, categoryCtrl)
	warehouseRoutes.RegisterWarehouseRoutes(api, warehouseCtrl)
	SupplierRoutes.RegisterSupplierRoutes(api, supplierCtrl)
	itemRoutes.RegisterItemRoutes(api, itemCtrl)
	transactionRoutes.RegisterTransactionRoutes(api, transactionCtrl)
	dashboardRoutes.RegisterDashboardRoutes(api, dashboardCtrl)
}
