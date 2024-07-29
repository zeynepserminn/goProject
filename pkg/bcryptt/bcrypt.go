package bcryptt

import (
	error2 "goProject/internal"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", error2.ErrHashFailed
	}
	return string(hashedPassword), nil
}
