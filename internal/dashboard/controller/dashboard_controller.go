package controller

import "github.com/gofiber/fiber/v2"

type DashboardController interface {
	CountData(ctx *fiber.Ctx) error
}
