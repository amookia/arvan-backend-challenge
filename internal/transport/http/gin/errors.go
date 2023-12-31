package gin

import (
	"github.com/amookia/arvan-backend-challenge/internal/transport/http/response"
	"github.com/go-playground/validator/v10"
)

func ErrorHandler(message string,errs ...error) response.Error {
	var errors []string
	for _, err := range errs {
		errors = append(errors, err.Error())
	}
	return response.Error{Errors: errors,Message: message}
}

func FormErrorHandler(err error) response.Error {
	var errors []string
	for _, fieldErr := range err.(validator.ValidationErrors) {
		errors = append(errors, fieldErr.Error())
	}
	return response.Error{Errors: errors, Message: "invalid form"}
}
