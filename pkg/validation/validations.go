package validation

import (
	"fmt"
	"github.com/go-playground/validator/v10"

	"strings"
)

func RegisterValidation(v *validator.Validate) error {
	if err := v.RegisterValidation("max", ValidateMaxLen); err != nil {
		return err
	}
	if err := v.RegisterValidation("validEmail", IsValidEmail); err != nil {
		return err
	}

	if err := v.RegisterValidation("alpha", IsAlpha); err != nil {
		return err
	}
	if err := v.RegisterValidation("numeric", IsNumeric); err != nil {
		return err
	}
	if err := v.RegisterValidation("validPhone", IsValidPhone); err != nil {
		return err
	}
	err := v.RegisterValidation("validPassword", IsValidPassword)

	if err != nil {
		return err
	}

	return nil
}

func ValidateStruct(data interface{}) error {
	validate := validator.New()
	if err := RegisterValidation(validate); err != nil {
		return fmt.Errorf("error registering validations: %w", err)
	}

	err := validate.Struct(data)
	if err != nil {
		var errorMessages []string
		for _, err := range err.(validator.ValidationErrors) {
			fieldName := err.Field()
			tagName := err.Tag()
			errorMessage := fmt.Sprintf("Field validation for '%s' failed on the '%s' tag", fieldName, tagName)
			errorMessages = append(errorMessages, errorMessage)
		}
		return fmt.Errorf(strings.Join(errorMessages, "\n"))
	}
	return nil
}
