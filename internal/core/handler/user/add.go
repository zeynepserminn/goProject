package user

import (
	"github.com/gin-gonic/gin"
	"goProject/internal/core/dto"
	"goProject/pkg/validation"
	"net/http"
)

func (h *UserHandler) AddUser(c *gin.Context) {
	var user dto.AddUserRequest

	if err := c.ShouldBindJSON(&user); err != nil {
		sendResponse(c, http.StatusBadRequest, InvalidData, err.Error())
		return
	}

	if err := validation.ValidateStruct(user); err != nil {

		sendResponse(c, http.StatusBadRequest, ValidationFailed, err.Error())
		return

	}

	response, err := h.userService.AddUser(user)
	if err != nil {
		sendResponse(c, http.StatusInternalServerError, FailedToAddUser, err.Error())
		return
	}
	sendResponse(c, http.StatusOK, response, "user added")
}
