package services

import (
	"goProject/internal/core/dto"
	"goProject/internal/core/model"
	"goProject/internal/interface"
)

type UserServiceImpl struct {
	userRepo _interface.UserPort
}
type UserServiceI interface {
	AddUser(user dto.AddUserDTO) error
	GetAllUsers() []model.User
	UpdateUser(user dto.UpdateUserDTO) error
	DeleteUser(id dto.DeleteUserDTO) error
}

func NewUserService(userRepo _interface.UserPort) UserServiceI {
	return &UserServiceImpl{userRepo: userRepo}

}
func (u *UserServiceImpl) AddUser(userDto dto.AddUserDTO) error {
	user := model.User{
		Firstname: userDto.FirstName,
		Lastname:  userDto.Lastname,
		Email:     userDto.Email,
		Phone:     userDto.Phone,
		ID:        userDto.Id,
	}
	if err := u.userRepo.AddUser(user); err != nil {
		return err
	}
	return nil
}

func (u *UserServiceImpl) GetAllUsers() []model.User {
	return u.userRepo.GetAllUsers()
}

func (u *UserServiceImpl) UpdateUser(userDto dto.UpdateUserDTO) error {
	user := model.User{
		Firstname: userDto.FirstName,
		Lastname:  userDto.Lastname,
		Email:     userDto.Email,
		Phone:     userDto.Phone,
		ID:        userDto.Id,
	}
	if err := u.userRepo.UpdateUser(user); err != nil {
		return err
	}
	return nil
}

func (u *UserServiceImpl) DeleteUser(id dto.DeleteUserDTO) error {
	return u.userRepo.DeleteUser(id.Id)

}
