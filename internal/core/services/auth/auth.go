package auth

import (
	"errors"
	"goProject/internal"
	"goProject/internal/core/model"
	"goProject/pkg/jwt"
	"goProject/pkg/postgres/repositories"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo   *repositories.UserPortImpl
	jwtService *jwt.Jwt
}

func NewAuthService(userRepo *repositories.UserPortImpl, jwtService *jwt.Jwt) *AuthService {
	return &AuthService{userRepo: userRepo,
		jwtService: jwtService,
	}
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (a AuthService) Login(email, password string) (*LoginResponse, error) {

	user, err := a.userRepo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found") //TODO:şifren yok
	}
	if user.Status != model.Active {
		return nil, internal.ErrUserNotActive
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("password doesn't match")
	}
	userToken := jwt.UserToken{

		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
		Phone:     user.Phone,
		UserID:    user.ID,
		Role:      user.Role,
	}
	accessToken, _, err := a.jwtService.GenerateAccessToken(userToken)
	if err != nil {
		return nil, err
	}

	refreshToken, _, err := a.jwtService.GenerateRefreshToken(userToken)
	if err != nil {
		return nil, err
	}

	response := &LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return response, nil
}
