package userservice

import (
	"errors"
	"goProject/internal/core/dto"
	"goProject/internal/core/model"
	"goProject/internal/interface"
)

type UserServiceImpl struct {
	userRepo _interface.UserPort
}
type UserServiceI interface {
	AddUser(user dto.AddUserRequest) (AddUserResponse, error)
	GetAllUsers(pg dto.PaginationRequest, filters dto.FilterParams) ([]dto.UserResponse, int, error)
	UpdateUser(user dto.UpdateUserRequest) error
	DeleteUser(id dto.DeleteUserRequest) error
	GetUserByID(id dto.GetUserByIdDTO) (*dto.UserResponse, error)
	UpdateProfile(userID int64, request dto.UpdateProfileRequest) error
	UpdatePassword(userID int64, request dto.UpdatePasswordRequest) error
	GetProfile(userID int64) (*dto.UserResponse, error)
}

func NewUserService(userRepo _interface.UserPort) UserServiceI {
	return &UserServiceImpl{userRepo: userRepo}

}
func isValidRole(role string) (model.UserRole, error) {
	switch role {
	case "admin":
		return model.RoleAdmin, nil
	case "user":
		return model.RoleUser, nil
	default:
		return model.UserRole(-1), errors.New("user role not exist")
	}
}
