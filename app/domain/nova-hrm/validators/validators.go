package validators

import (
	"log"
	"regexp"

	"github.com/go-playground/validator"
)

// Initialize the validator
var Validate *validator.Validate

func init() {
	log.Println("Initializing validator")
	Validate = validator.New()

	// Register custom validation: "startswithalpha"
	Validate.RegisterValidation("startswithalpha", func(fl validator.FieldLevel) bool {
		log.Println("Validating startsWithAlpha")
		value := fl.Field().String()
		matched, _ := regexp.MatchString("^[A-Za-z]", value)
		return matched
	})
}
