package user

import (
	"github.com/gin-gonic/gin"
	"goProject/internal/core/dto"
	"goProject/pkg/jwt"
	"net/http"
)

func (h *UserHandler) UpdatePassword(c *gin.Context) {
	var request dto.UpdatePasswordRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		sendResponse(c, http.StatusBadRequest, InvalidData, err.Error())
		return
	}
	user, exists := c.Get("user")
	if !exists {
		sendResponse(c, http.StatusInternalServerError, UserNotExists, "user not exists")
		return
	}
	claims, ok := user.(*jwt.UserToken)
	if !ok {
		sendResponse(c, http.StatusInternalServerError, InvalidID, "invalid id")
		return
	}
	userID := claims.UserID
	if err := h.userService.UpdatePassword(userID, request); err != nil {

		sendResponse(c, http.StatusInternalServerError, FailedToUpdateUser, "failed to update password.")
		return
	}
	sendResponse(c, http.StatusOK, UserPasswordUpdated, "user password updated successfully")
}
