package auth

import (
	"github.com/gin-gonic/gin"
	"goProject/internal/core/dto"
	"goProject/internal/core/handler/user"
	"goProject/pkg/jwt"
	"net/http"
)

type RefreshTokenHandler struct {
	jwtService *jwt.Jwt
}

func NewRefreshToken(jwtService *jwt.Jwt) *RefreshTokenHandler {
	return &RefreshTokenHandler{jwtService: jwtService}
}
func (h *RefreshTokenHandler) RefreshAccessToken(c *gin.Context) {
	var request dto.RefreshAccessTokenRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, user.Response{Result: nil, Message: "invalid request"})
		return
	}
	if _, err := h.jwtService.ValidateRefreshToken(request.RefreshToken); err != nil {
		c.JSON(http.StatusBadRequest, user.Response{Result: nil, Message: "Validation failed for refreshing access token."})
		return
	}
	newAccessToken, err := h.jwtService.RefreshAccessToken(request.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, user.Response{Result: nil, Message: "Failed to refresh access token"})
		return
	}
	c.JSON(http.StatusOK, user.Response{Result: newAccessToken, Message: "Successfully refreshed access token"})
}
