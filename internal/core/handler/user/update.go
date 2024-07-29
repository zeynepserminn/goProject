package user

import (
	"github.com/gin-gonic/gin"
	"goProject/internal/core/dto"
	"goProject/pkg/validation"
	"net/http"
	"strconv"
)

func (h *UserHandler) UpdateUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	user := dto.UpdateUserRequest{ID: int32(id)}

	if err := c.ShouldBindJSON(&user); err != nil {
		sendResponse(c, http.StatusBadRequest, InvalidData, err.Error())
		return
	}

	if err != nil {
		sendResponse(c, http.StatusBadRequest, InvalidID, err.Error())
		return
	}

	if err := validation.ValidateStruct(user); err != nil {

		sendResponse(c, http.StatusBadRequest, ValidationFailed, err.Error())
		return

	}

	if err := h.userService.UpdateUser(user); err != nil {
		sendResponse(c, http.StatusInternalServerError, FailedToUpdateUser, err.Error())
		return
	}
	sendResponse(c, http.StatusOK, UserUpdated, "user updated")
}
