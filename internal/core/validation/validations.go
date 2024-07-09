package validation

import (
	"github.com/go-playground/validator/v10"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

func RegisterValidation(v *validator.Validate) error {
	if err := v.RegisterValidation("maxlenght", ValidateMaxLen); err != nil {
		return err
	}
	if err := v.RegisterValidation("emailcheck", ValidateEmail); err != nil {
		return err
	}
	return nil
}

func ValidateMaxLen(fl validator.FieldLevel) bool {
	tag := fl.Field().String() //?
	maxlen, err := strconv.Atoi(tag)
	if err != nil {
		return false
	}
	fieldLen := utf8.RuneCountInString(fl.Field().String())
	return fieldLen <= maxlen
}

func ValidateEmail(fl validator.FieldLevel) bool {
	email := fl.Field().String()

	if len(email) == 0 || !strings.Contains(email, "@") {
		return false
	}
	return true
}

func ValidateID(fl validator.FieldLevel) bool {
	id := fl.Field().String()
	if id == "" {
		return false
	}
	idMatch, err := regexp.MatchString(`^[a-zA-Z0-9\-]+$`, id)
	if err != nil || !idMatch {
		return false
	}
	return true
}
