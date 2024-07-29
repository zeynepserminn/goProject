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
	if user.Email != request.Email {
		emailExists, err := s.userRepo.IsEmailExists(request.Email, &userID)
		if err != nil {
			return err
		}
		if emailExists {
			return internal.ErrEmailExists
		}
	}

	if user.Phone != request.Phone {
		phoneExists, err := s.userRepo.IsPhoneExists(request.Phone)
		if err != nil {
			return err
		}
		if phoneExists {
			return internal.ErrPhoneExists
		}
	}

	user.Firstname = request.Firstname
	user.Lastname = request.Lastname
	user.Email = request.Email
	user.Phone = request.Phone

	return s.userRepo.UpdateUser(*user)
}
