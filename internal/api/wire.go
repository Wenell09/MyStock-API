//go:build wireinject
// +build wireinject

package api

import (
	"github.com/Wenell09/MyStock/internal/auth"
	"github.com/Wenell09/MyStock/internal/category"
	"github.com/Wenell09/MyStock/internal/dashboard"
	"github.com/Wenell09/MyStock/internal/database"
	"github.com/Wenell09/MyStock/internal/item"
	"github.com/Wenell09/MyStock/internal/supplier"
	"github.com/Wenell09/MyStock/internal/transaction"
	"github.com/Wenell09/MyStock/internal/warehouse"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
)

// Injector
func InitApp() (*fiber.App, error) {
	wire.Build(
		logrus.New,
		validator.New,
		database.DBConnection,
		category.WireSet,
		warehouse.WireSet,
		supplier.WireSet,
		item.WireSet,
		transaction.WireSet,
		dashboard.WireSet,
		auth.WireSet,
		NewFiberApp,
	)

	return nil, nil
}
