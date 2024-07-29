package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"goProject/internal/core/model"
	"time"
)

type UserToken struct {
	Firstname string         `json:"first_name"`
	Lastname  string         `json:"last_name"`
	Email     string         `json:"email"`
	Phone     string         `json:"phone"`
	UserID    int64          `json:"user_id"`
	Role      model.UserRole `json:"role"`

	jwt.RegisteredClaims
}
type Jwt struct {
	AccessSecret    []byte
	AccessInterval  time.Duration
	RefreshSecret   []byte
	RefreshInterval time.Duration
}

func NewJwt() *Jwt {
	return &Jwt{
		AccessSecret:    []byte("secret_key"),
		AccessInterval:  30 * time.Minute,
		RefreshSecret:   []byte("refresh_secret"),
		RefreshInterval: 24 * 180 * time.Hour,
	}
}

func (j *Jwt) generateToken(user UserToken, secretKey []byte, duration time.Duration) (string, time.Time, error) {
	expirationTime := time.Now().Add(duration)
	user.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(expirationTime),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, user)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", time.Time{}, err
	}
	return tokenString, expirationTime, nil
}
func (j *Jwt) GenerateAccessToken(user UserToken) (string, time.Time, error) {
	return j.generateToken(user, j.AccessSecret, j.AccessInterval)
}
func (j *Jwt) GenerateRefreshToken(user UserToken) (string, time.Time, error) {
	j.RefreshInterval = 24 * 180 * time.Hour
	return j.generateToken(user, j.RefreshSecret, j.RefreshInterval)
}

func (j *Jwt) ValidateRefreshToken(refreshToken string) (*UserToken, error) {
	token, err := jwt.ParseWithClaims(refreshToken, &UserToken{}, func(token *jwt.Token) (interface{}, error) {
		return j.RefreshSecret, nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	if claims, ok := token.Claims.(*UserToken); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("invalid refresh token")
	}
}

func (j *Jwt) RefreshAccessToken(refreshToken string) (string, error) {
	claims, err := j.ValidateRefreshToken(refreshToken)
	if err != nil {
		return "", err
	}

	newAccessToken, _, err := j.GenerateAccessToken(*claims)
	if err != nil {
		return "", err
	}
	return newAccessToken, nil
}
