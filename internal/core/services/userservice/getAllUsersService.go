package userservice

import (
	"goProject/internal"
	"goProject/internal/core/dto"
)

func (u *UserServiceImpl) GetAllUsers(pg dto.PaginationRequest, filters dto.FilterParams) ([]dto.UserResponse, int, error) {
	users, total, err := u.userRepo.GetAllUsers(pg, filters)
	if err != nil {
		return nil, 0, internal.ErrFetchingUsers
	}
	getAllResponse := make([]dto.UserResponse, len(users))
	for i, user := range users {
		getAllResponse[i] = dto.UserResponse{
			ID:        user.ID,
			FirstName: user.Firstname,
			Lastname:  user.Lastname,
			Email:     user.Email,
			Phone:     user.Phone,
		}
	}
	return getAllResponse, total, nil
}
