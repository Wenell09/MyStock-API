//go:build wireinject
// +build wireinject

package api

import (
	"github.com/Wenell09/MyStock/internal/category"
	"github.com/Wenell09/MyStock/internal/database"
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
		NewFiberApp,
	)

	return nil, nil
}
