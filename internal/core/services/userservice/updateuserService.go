package userservice

import (
	"errors"
	"goProject/internal"
	"goProject/internal/core/dto"
	"goProject/internal/core/model"
	"goProject/pkg/bcryptt"
)

func (u *UserServiceImpl) UpdateUser(userDto dto.UpdateUserRequest) error {
	existinguser, err := u.userRepo.GetUserByID(userDto.ID)
	if err != nil {
		if errors.Is(err, internal.ErrUserNotFound) {
			return internal.ErrUserNotFound
		}
		return err
	}

	if existinguser.Status == model.Passive || existinguser.Status == model.Deleted {
		return internal.ErrUserNotActive

	}

	if userDto.Email != existinguser.Email {
		userID := int64(userDto.ID)
		emailExists, err := u.userRepo.IsEmailExists(userDto.Email, &userID)
		if err != nil {
			return err
		}
		if emailExists {
			return internal.ErrEmailExists
		}
	}
	if userDto.Phone != existinguser.Phone {
		phoneExists, err := u.userRepo.IsPhoneExists(userDto.Phone)
		if err != nil {
			return err
		}
		if phoneExists {
			return internal.ErrPhoneExists
		}
	}

	hashedPassword, err := bcryptt.HashPassword(userDto.Password)
	if err != nil {
		return internal.ErrHashFailed
	}

	user := model.User{
		Firstname: userDto.FirstName,
		Lastname:  userDto.Lastname,
		Email:     userDto.Email,
		Phone:     userDto.Phone,
		ID:        existinguser.ID,
		Password:  hashedPassword,
		Status:    existinguser.Status,
		Role:      userDto.Role,
	}
	if err := u.userRepo.UpdateUser(user); err != nil {
		return err
	}
	return nil
}
