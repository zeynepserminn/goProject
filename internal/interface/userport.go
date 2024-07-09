package _interface

import (
	"goProject/internal/core/model"
)

type UserPort interface {
	AddUser(user model.User) error
	GetAllUsers() []model.User
	UpdateUser(user model.User) error
	DeleteUser(id string) error
}
