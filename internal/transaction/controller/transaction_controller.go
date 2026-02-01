package controller

import "github.com/gofiber/fiber/v2"

type TransactionController interface {
	CreateOrUpdate(ctx *fiber.Ctx) error
	Read(ctx *fiber.Ctx) error
	ReadByPublicId(ctx *fiber.Ctx) error
}
