package user

import (
	"github.com/gin-gonic/gin"
	"goProject/internal/core/dto"
	"goProject/pkg/validation"
	"net/http"
)

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	var pagination dto.PaginationRequest
	var filters dto.FilterParams
	var err error

	if err = c.ShouldBind(&pagination); err != nil {
		sendResponse(c, http.StatusBadRequest, InvalidPagination, err.Error())
		return
	}
	if err = validation.ValidateStruct(&pagination); err != nil {

		sendResponse(c, http.StatusBadRequest, ValidationFailed, err.Error())
		return

	}

	if err = c.ShouldBind(&filters); err != nil {
		sendResponse(c, http.StatusBadRequest, InvalidFilters, err.Error())
		return
	}

	if err = validation.ValidateStruct(&filters); err != nil {

		sendResponse(c, http.StatusBadRequest, ValidationFailed, err.Error())
		return

	}

	users, total, err := h.userService.GetAllUsers(pagination, filters)
	if err != nil {
		sendResponse(c, http.StatusInternalServerError, FailedToGetUsers, err.Error())
		return
	}

	c.JSON(http.StatusOK, Response{Result: gin.H{
		"users": users,
		"total": total,
	}, Message: "Successfully get users."})
}
