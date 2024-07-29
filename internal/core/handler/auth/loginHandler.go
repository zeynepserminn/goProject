package auth

import (
	"github.com/gin-gonic/gin"
	"goProject/internal/core/dto"
	"goProject/internal/core/handler/user"
	"goProject/internal/core/services/auth"
	"goProject/pkg/validation"
	"net/http"
)

type AuthHandler struct {
	authService auth.AuthService
}

func NewAuthHandler(authService *auth.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: *authService,
	}
}

func (h *AuthHandler) Login(c *gin.Context) {

	var loginDTO dto.LoginDTO

	if err := c.ShouldBindJSON(&loginDTO); err != nil {
		c.JSON(http.StatusBadRequest, user.Response{Result: nil, Message: "Invalid request."})
		return
	}

	if err := validation.ValidateStruct(&loginDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.authService.Login(loginDTO.Email, loginDTO.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, user.Response{Result: nil, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, user.Response{Result: response, Message: "Success"})

}
