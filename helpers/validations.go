package helpers

import (
	"go-heroku/dtos"
	"go-heroku/langs"

	"gopkg.in/go-playground/validator.v8"
)

func GenerateValidationResponse(err error) (response dtos.ValidationResponse) {
	response.Success = false

	var validations []dtos.Validation

	//get validation errors
	validationErrors := err.(validator.ValidationErrors)

	for _, value := range validationErrors {
		//get field & ruled (tag)
		field, rule := value.Field, value.Tag

		//create validation object
		validation := dtos.Validation{Field: field, Message: langs.GenerateValidationMessage(field, rule)}

		//add validation object to validations
		validations = append(validations, validation)
	}

	response.Validations = validations

	return response

}
