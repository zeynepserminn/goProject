package user

import (
	"errors"
	"github.com/gin-gonic/gin"
	"goProject/internal"
	"goProject/internal/core/dto"
	"goProject/pkg/validation"
	"net/http"
	"strconv"
)

func (h *UserHandler) DeleteUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		sendResponse(c, http.StatusBadRequest, InvalidID, err.Error())
		return
	}
	user := dto.DeleteUserRequest{ID: int32(id)}

	if err := validation.ValidateStruct(user); err != nil {
		if errors.Is(err, internal.ErrValidationFailed) {
			sendResponse(c, http.StatusBadRequest, ValidationFailed, err.Error())
			return
		}
	}

	if err := h.userService.DeleteUser(user); err != nil {
		sendResponse(c, http.StatusInternalServerError, FailedToDeleteUser, err.Error())
		return
	}

	sendResponse(c, http.StatusOK, UserDeleted, "user deleted")
}
