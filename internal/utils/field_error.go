package utils

import "github.com/go-playground/validator"

type FieldError struct {
	Errors []FieldDetail `json:"errors"`
}

type FieldDetail struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

func (f FieldError) Error() string {
	return "validation error"
}

func NewFieldError(validationErr validator.ValidationErrors) error {
	dataErr := []FieldDetail{}

	for _, fieldErr := range validationErr {
		dataErr = append(dataErr, FieldDetail{
			Field: fieldErr.Field(),
			Error: fieldErr.Tag(),
		})
	}

	return FieldError{
		Errors: dataErr,
	}
}
