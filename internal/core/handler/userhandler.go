package handler

import (
	"github.com/gin-gonic/gin"
	"goProject/internal/core/dto"
	"goProject/internal/core/services"
	"net/http"
)

type UserHandler struct {
	userService services.UserServiceI
}

func NewUserHandler(userService services.UserServiceI) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) AddUser(c *gin.Context) {
	var user dto.AddUserDTO

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.userService.AddUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"message": "user added"})
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {

	users := h.userService.GetAllUsers()
	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	var user dto.UpdateUserDTO
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.userService.UpdateUser(user); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user updated"})
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	var user dto.DeleteUserDTO
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.userService.DeleteUser(user); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user deleted"})
}
