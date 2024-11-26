package utils

import (
	"eventapp/models"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateEvent(event *models.Event) error {
	return validate.Struct(event)
}
