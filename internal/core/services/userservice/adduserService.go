package userservice

import (
	"goProject/internal"
	"goProject/internal/core/dto"
	"goProject/internal/core/model"
	"goProject/pkg/bcryptt"
)

type AddUserResponse struct {
	ID int64 `json:"id"`
}

func (u *UserServiceImpl) AddUser(userDto dto.AddUserRequest) (AddUserResponse, error) {
	emailExists, err := u.userRepo.IsEmailExists(userDto.Email, nil)
	if err != nil {
		return AddUserResponse{}, err
	}
	if emailExists {
		return AddUserResponse{}, internal.ErrEmailExists
	}
	phoneExists, err := u.userRepo.IsPhoneExists(userDto.Phone)
	if err != nil {
		return AddUserResponse{}, err
	}
	if phoneExists {
		return AddUserResponse{}, internal.ErrPhoneExists
	}

	hashedPassword, err := bcryptt.HashPassword(userDto.Password)
	if err != nil {
		return AddUserResponse{}, internal.ErrHashFailed
	}
	user := model.User{
		Firstname: userDto.FirstName,
		Lastname:  userDto.Lastname,
		Email:     userDto.Email,
		Phone:     userDto.Phone,
		Password:  hashedPassword,
		Status:    model.Active,
		Role:      userDto.Role,
	}
	userID, err := u.userRepo.AddUser(user)
	if err != nil {
		return AddUserResponse{}, err
	}
	return AddUserResponse{ID: userID}, err
}
