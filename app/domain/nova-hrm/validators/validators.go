package validators

import (
	"regexp"

	"github.com/go-playground/validator"
)

// Initialize the validator
var Validate *validator.Validate

func init() {
	Validate = validator.New()

	// Register custom validation: "startswithalpha"
	Validate.RegisterValidation("startswithalpha", func(fl validator.FieldLevel) bool {
		value := fl.Field().String()
		matched, _ := regexp.MatchString("^[A-Za-z]", value)
		return matched
	})
}
