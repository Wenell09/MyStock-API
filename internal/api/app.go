package api

import (
	categoryController "github.com/Wenell09/MyStock/internal/category/controller"
	dashboardController "github.com/Wenell09/MyStock/internal/dashboard/controller"
	itemController "github.com/Wenell09/MyStock/internal/item/controller"
	supplierController "github.com/Wenell09/MyStock/internal/supplier/controller"
	transactionController "github.com/Wenell09/MyStock/internal/transaction/controller"
	warehouseController "github.com/Wenell09/MyStock/internal/warehouse/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func NewFiberApp(
	categoryCtrl categoryController.CategoryController,
	warehouseCtrl warehouseController.WarehouseController,
	supplierCtrl supplierController.SupplierController,
	itemCtrl itemController.ItemController,
	transactionCtrl transactionController.TransactionController,
	dashboardCtrl dashboardController.DashboardController,
) *fiber.App {

	app := fiber.New()

	// Global middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
	}))

	RegisterRoutes(app, categoryCtrl, warehouseCtrl, supplierCtrl, itemCtrl, transactionCtrl, dashboardCtrl)

	return app
}
