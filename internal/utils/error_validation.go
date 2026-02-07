package utils

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type ValidationError struct {
	Msg string
}

func (v ValidationError) Error() string {
	return v.Msg
}

func (v ValidationError) StatusCode() int {
	return fiber.StatusBadRequest
}

func (v ValidationError) ResponseMessage() string {
	return "Validation Error"
}

func (v ValidationError) ErrorData() interface{} {
	return v.Msg
}

func NewFieldError(validationErr validator.ValidationErrors) error {
	var messages []string
	for _, fieldErr := range validationErr {
		messages = append(messages,
			fmt.Sprintf("%s %s", fieldErr.Field(), fieldErr.Tag()),
		)
	}
	return ValidationError{
		Msg: strings.Join(messages, ", "),
	}
}
