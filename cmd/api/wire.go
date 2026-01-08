//go:build wireinject
// +build wireinject

package api

import (
	"github.com/Wenell09/MyStock/internal/category/controller"
	"github.com/Wenell09/MyStock/internal/category/repository"
	"github.com/Wenell09/MyStock/internal/category/service"
	"github.com/Wenell09/MyStock/internal/database"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var categorySet = wire.NewSet(
	repository.NewCategoryRepository,
	service.NewCategoryService,
	controller.NewCategoryController,
)

// Injector
func InitApp() (*fiber.App, error) {
	wire.Build(
		// external
		godotenv.Load,
		logrus.New,
		validator.New,
		database.DBConnection,
		// layer
		categorySet,
		// app
		NewFiberApp,
	)
	return nil, nil
}
