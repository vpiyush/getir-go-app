package apis

import (
	"errors"
	"github.com/vpiyush/getir-go-app/models"
	"gopkg.in/go-playground/validator.v9"
)

//ValidatePair validates a pair, and returns an error object if validation fails
func validatePair(pair *models.Pair) error {
	validate := validator.New()
	err := validate.Struct(pair)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return errors.New(err.Field() + " value is invalid")
		}
	}
	return nil
}

//ValidateRequest validates a record fetch request object,
//and returns an error object if validation fails
func validateRequest(req *Request) error {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return errors.New(err.Field() + " value is invalid")
		}
	}
	return nil
}
