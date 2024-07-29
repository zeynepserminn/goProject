package userservice

import (
	"goProject/internal"
	"goProject/internal/core/dto"
)

func (s *UserServiceImpl) GetProfile(userID int64) (*dto.UserResponse, error) {
	user, err := s.userRepo.GetUserByID(int32(userID))
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, internal.ErrUserNotFound
	}

	userResponse := &dto.UserResponse{
		FirstName: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
		Phone:     user.Phone,
		ID:        user.ID,
		Status:    user.Status,
		Role:      user.Role,
	}

	return userResponse, nil
}
