package utils

import (
	"github.com/gofiber/fiber/v2"
)

type ValidationError struct {
	Msg string
}

func (v ValidationError) Error() string {
	return v.Msg
}

type NotFoundError struct {
	Msg string
}

func (n NotFoundError) Error() string {
	return n.Msg
}

func NewHandleError(ctx *fiber.Ctx, err error, message string) error {
	switch e := err.(type) {
	case ValidationError:
		return ctx.Status(fiber.StatusBadRequest).JSON(
			NewResponseError(fiber.StatusBadRequest, "Validation Error", e.Error()),
		)
	case NotFoundError:
		return ctx.Status(fiber.StatusNotFound).JSON(
			NewResponseError(fiber.StatusNotFound, "Not Found", e.Error()),
		)
	default:
		return ctx.Status(fiber.StatusInternalServerError).JSON(
			NewResponseError(fiber.StatusInternalServerError, message, e.Error()),
		)
	}
}
