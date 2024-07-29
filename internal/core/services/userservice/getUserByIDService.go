package userservice

import (
	"goProject/internal"
	"goProject/internal/core/dto"
)

func (u *UserServiceImpl) GetUserByID(id dto.GetUserByIdDTO) (*dto.UserResponse, error) {
	user, err := u.userRepo.GetUserByID(id.ID)
	if err != nil {

		return nil, err
	}
	if user == nil {
		return nil, internal.ErrUserNotFound
	}
	response := &dto.UserResponse{
		ID:        user.ID,
		FirstName: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
		Phone:     user.Phone,
		Status:    user.Status,
		Role:      user.Role,
	}
	return response, nil
}
