package controller

import "github.com/gofiber/fiber/v2"

type SupplierController interface {
	Create(ctx *fiber.Ctx) error
	Read(ctx *fiber.Ctx) error
	ReadByPublicId(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	DeleteAll(ctx *fiber.Ctx) error
}
