package user

import (
	"goProject/internal/core/services/userservice"
)

type UserHandler struct {
	userService userservice.UserServiceI
}

func NewUserHandler(userService userservice.UserServiceI) *UserHandler {
	return &UserHandler{userService: userService}
}
