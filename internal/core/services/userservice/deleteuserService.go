package userservice

import (
	"goProject/internal"
	"goProject/internal/core/dto"
	"goProject/internal/core/model"
)

func (u *UserServiceImpl) DeleteUser(id dto.DeleteUserRequest) error {
	user, err := u.userRepo.GetUserByID(id.ID)
	if err != nil {
		return err
	}
	if user == nil {
		return internal.ErrUserNotFound
	}
	if user.Status != model.Active {
		return internal.ErrUserNotActive
	}

	if err := u.userRepo.DeleteUser(id.ID); err != nil {

		return err
	}
	return nil
}
