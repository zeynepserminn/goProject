package dto

import "goProject/internal/core/model"

type UserResponse struct {
	FirstName string           `json:"first_name" `
	Lastname  string           `json:"last_name" `
	Email     string           `json:"email" `
	Phone     string           `json:"phone_number" `
	ID        int64            `json:"id" `
	Status    model.UserStatus `json:"status"`
	Role      model.UserRole   `json:"role"`
}
