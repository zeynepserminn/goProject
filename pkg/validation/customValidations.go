package validation

import (
	"github.com/go-playground/validator/v10"
	"regexp"
	"strconv"
	"unicode/utf8"
)

func IsAlpha(fl validator.FieldLevel) bool {
	return regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(fl.Field().String())
}
func IsNumeric(fl validator.FieldLevel) bool {
	return regexp.MustCompile(`^[0-9]+$`).MatchString(fl.Field().String())
}

func ValidateMaxLen(fl validator.FieldLevel) bool {
	tag := fl.Param()
	maxlen, err := strconv.Atoi(tag)
	if err != nil {
		return false
	}
	fieldValue := fl.Field().String()
	fieldLen := utf8.RuneCountInString(fieldValue)
	return fieldLen <= maxlen
}

func IsValidEmail(fl validator.FieldLevel) bool {
	email := fl.Field().String()
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)

	return re.MatchString(email)
}

func IsValidPassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	if len(password) < 8 {
		return false
	}
	upper := `[A-Z]`
	reUpper := regexp.MustCompile(upper)
	if !reUpper.MatchString(password) {
		return false
	}
	special := `[!@#~$%^&*()_+{}:"<>?]`
	reSpecial := regexp.MustCompile(special)
	if !reSpecial.MatchString(password) {
		return false
	}
	return true
}

func IsValidPhone(fl validator.FieldLevel) bool {
	phone := fl.Field().String()
	const phoneRegex = `^\+?[1-9]\d{1,14}(?:x.+)?$`
	re := regexp.MustCompile(phoneRegex)
	return re.MatchString(phone)
}
