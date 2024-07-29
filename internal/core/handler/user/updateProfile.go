package user

import (
	"github.com/gin-gonic/gin"
	"goProject/internal/core/dto"
	"goProject/pkg/jwt"
	"net/http"
)

func (h *UserHandler) UpdateProfile(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		sendResponse(c, http.StatusUnauthorized, UserNotExists, "user not found")
		return
	}
	claims, ok := user.(*jwt.UserToken)
	if !ok {
		sendResponse(c, http.StatusInternalServerError, InvalidID, "invalid id")
		return
	}

	var request dto.UpdateProfileRequest
	if err := c.ShouldBind(&request); err != nil {
		sendResponse(c, http.StatusInternalServerError, InvalidData, "invalid data")
		return
	}

	if err := h.userService.UpdateProfile(claims.UserID, request); err != nil {
		sendResponse(c, http.StatusInternalServerError, FailedToUpdateUser, "failed to update user profile")
		return
	}
	sendResponse(c, http.StatusOK, UserProfileUpdated, "user profile updated")
}
