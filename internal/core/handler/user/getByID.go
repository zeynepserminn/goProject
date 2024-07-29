package user

import (
	"github.com/gin-gonic/gin"
	"goProject/internal/core/dto"
	"goProject/pkg/validation"
	"net/http"
	"strconv"
)

func (h *UserHandler) GetUserByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		sendResponse(c, http.StatusBadRequest, InvalidID, err.Error())
		return
	}

	user := dto.GetUserByIdDTO{ID: int32(id)}

	if err := validation.ValidateStruct(user); err != nil {
		sendResponse(c, http.StatusBadRequest, ValidationFailed, err.Error())
		return
	}

	found, err := h.userService.GetUserByID(user)
	if err != nil {
		sendResponse(c, http.StatusInternalServerError, FailedToGetUserByID, err.Error())
		return
	}

	sendResponse(c, http.StatusOK, found, UserRetrieved)
}
