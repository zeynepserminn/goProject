package userservice

import (
	"errors"
	"goProject/internal"
	"goProject/internal/core/dto"
	"goProject/internal/core/model"
	"goProject/pkg/bcryptt"
	"golang.org/x/crypto/bcrypt"
)

func (s *UserServiceImpl) UpdatePassword(userID int64, request dto.UpdatePasswordRequest) error {
	user, err := s.userRepo.GetUserByID(int32(userID))
	if err != nil {
		return err
	}
	if user == nil {
		return internal.ErrUserNotFound
	}
	if user.Status != model.Active {
		return errors.New("user is not active")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Old)); err != nil {
		return errors.New("old password is invalid")
	}
	hashedPassword, err := bcryptt.HashPassword(request.New)
	if err != nil {
		return err
	}
	userValue := *user
	user.Password = hashedPassword
	return s.userRepo.UpdateUser(userValue)
}
