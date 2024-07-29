package userservice

import (
	"goProject/internal"
	"goProject/internal/core/dto"
)

func (s *UserServiceImpl) UpdateProfile(userID int64, request dto.UpdateProfileRequest) error {
	user, err := s.userRepo.GetUserByID(int32(userID))
	if err != nil {
		return err
	}
	if user == nil {
		return internal.ErrUserNotFound
	}

	user.Firstname = request.Firstname
	user.Lastname = request.Lastname
	user.Email = request.Email
	user.Phone = request.Phone

	return s.userRepo.UpdateUser(*user)
}
