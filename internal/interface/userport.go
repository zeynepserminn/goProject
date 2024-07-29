package _interface

import (
	"goProject/internal/core/dto"
	"goProject/internal/core/model"
)

type UserPort interface {
	AddUser(user model.User) (int64, error)
	GetAllUsers(pagination dto.PaginationRequest, filters dto.FilterParams) ([]model.User, int, error)
	UpdateUser(user model.User) error
	DeleteUser(id int32) error
	GetUserByID(id int32) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	IsEmailExists(email string, exclude *int64) (bool, error)
	IsPhoneExists(phone string) (bool, error)
}
