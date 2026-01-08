package utils

import "github.com/gofiber/fiber/v2"

type NotFoundError struct {
	Msg string
}

func (n NotFoundError) Error() string {
	return n.Msg
}

func (n NotFoundError) StatusCode() int {
	return fiber.StatusNotFound
}

func (n NotFoundError) ResponseMessage() string {
	return "Not Found"
}

func (n NotFoundError) ErrorData() interface{} {
	return n.Msg
}
