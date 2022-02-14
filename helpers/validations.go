package helpers

import (
	"api-station/langs"
	"api-station/response"

	"github.com/go-playground/validator/v10"
	// "gopkg.in/go-playground/validator.v8"
)

func ValidationInputResponse(err error) (res response.ValidationResponse) {
	res.Success = false
	var validations []response.Validation
	for _, e := range err.(validator.ValidationErrors) {

		field, rule := e.Field(), e.Tag()
		validation := response.Validation{Field: field, Message: langs.GenerateValidationMessage(field, rule)}
		validations = append(validations, validation)
	}
	res.Validations = validations
	return res
}

func ValidationErrorResponse(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}
