package dto

type AddUserDTO struct {
	FirstName string `json:"first_name" validate:"required,maxlen=50" `
	Lastname  string `json:"lastname,omitempty"  validate:"required,maxlen=50"`
	Email     string `json:"email" validate:"required,email"`
	Phone     string `json:"phone,omitempty" validate:"required,maxlen=20" `
	Id        string `json:"id" validate:"required,id" `
}

type UpdateUserDTO struct {
	FirstName string `json:"first_name" validate:"required,maxlen=50" `
	Lastname  string `json:"lastname,omitempty" validate:"required,maxlen=50"`
	Email     string `json:"email" validate:"required,email"`
	Phone     string `json:"phone,omitempty" validate:"required,maxlen=20" `
	Id        string `json:"id" validate:"required,id"`
}

type DeleteUserDTO struct {
	Id string `json:"id" validate:"required,email" `
}
