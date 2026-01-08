package utils

import "github.com/gofiber/fiber/v2"

func NewHandleError(ctx *fiber.Ctx, err error) error {
	if appErr, ok := err.(AppError); ok {
		return ctx.Status(appErr.StatusCode()).JSON(
			NewResponseError(
				appErr.StatusCode(),
				appErr.ResponseMessage(),
				appErr.ErrorData(),
			),
		)
	}
	return ctx.Status(fiber.StatusInternalServerError).JSON(
		NewResponseError(
			fiber.StatusInternalServerError,
			"Internal Server Error",
			err.Error(),
		),
	)
}
