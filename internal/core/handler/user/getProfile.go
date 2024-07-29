package user

import (
	"github.com/gin-gonic/gin"
	"goProject/pkg/jwt"
	"net/http"
)

func (h *UserHandler) GetProfile(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		sendResponse(c, http.StatusUnauthorized, UserNotExists, "user not exists")
		return
	}

	claims, ok := user.(*jwt.UserToken)
	if !ok {
		sendResponse(c, http.StatusInternalServerError, InvalidID, "Invalid user ID type")
		return
	}

	profile, err := h.userService.GetProfile(claims.UserID)
	if err != nil {
		sendResponse(c, http.StatusInternalServerError, FailedToGetUserByID, "failed to get profile")
		return
	}

	sendResponse(c, http.StatusOK, profile, "user successfully retrieved")
}
