package dto

import "goProject/internal/core/model"

type AddUserRequest struct {
	FirstName string         `json:"first_name" validate:"required,max=50,alpha" `
	Lastname  string         `json:"last_name" validate:"required,max=50,alpha"`
	Email     string         `json:"email" validate:"required,validEmail"`
	Phone     string         `json:"phone_number" validate:"required,max=20,validPhone" `
	Password  string         `json:"password" validate:"required,validPassword" `
	Role      model.UserRole `json:"role" validate:"required"`
}

type UpdateUserRequest struct {
	FirstName string `json:"first_name" validate:"required,max=50,alpha" `
	Lastname  string `json:"last_name" validate:"required,max=50,alpha"`
	Email     string `json:"email" validate:"required,validEmail"`
	Phone     string `json:"phone_number" validate:"required,max=20,validPhone"`
	ID        int32  `json:"id" validate:"required"`
	Password  string `json:"password" `
	Status    int    `json:"status" `
	Role      int    `json:"role" validate:"required"`
}

type DeleteUserRequest struct {
	ID   int32 `json:"id" validate:"required"`
	Role int   `json:"role" validate:"required"`
}

type GetUserByIdDTO struct {
	ID int32 `json:"id" validate:"required"`
}
type GetUserByEmailRequest struct {
	Email string `json:"email" validate:"required,validEmail"`
}
type LoginDTO struct {
	Email    string `json:"email" validate:"required,validEmail"`
	Password string `json:"password" validate:"required,validPassword"`
}
type RefreshAccessTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}
type PaginationRequest struct {
	Limit   int    `form:"limit" validate:"required,min=1,max=100"`
	Skip    int    `form:"skip" validate:"omitempty,min=0,max=100"`
	SortBy  string `form:"sort_by" validate:"required,max=50,alpha"`
	OrderBy string `form:"order" validate:"required,max=50,oneof=desc asc"`
}
type FilterParams struct {
	Firstname *string `form:"firstname,omitempty" validate:"omitempty,max=50,alpha"`
	Lastname  *string `form:"lastname,omitempty" validate:"omitempty,max=50,alpha"`
	Email     *string `form:"email,omitempty" validate:"omitempty,max=50"`
	Phone     *string `form:"phone,omitempty" validate:"omitempty,max=50"`
}
type UpdateProfileRequest struct {
	Firstname string `json:"firstname" validate:"required,max=50,alpha"`
	Lastname  string `json:"lastname" validate:"required,max=50,alpha"`
	Email     string `json:"email" validate:"required,email"`
	Phone     string `json:"phone" validate:"required,max=20,numeric"`
}
type UpdatePasswordRequest struct {
	Old string `json:"old_password" validate:"required,max=50,alpha"`
	New string `json:"new_password" validate:"required,max=50,alpha"`
}
