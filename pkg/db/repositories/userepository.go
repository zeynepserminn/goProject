package repositories

import (
	"errors"
	"goProject/internal/core/model"
)

type UserPortImpl struct {
	User     []model.User
	Sequence int
}

func NewUserRepository() *UserPortImpl {
	return &UserPortImpl{}
}
func (us *UserPortImpl) AddUser(user model.User) error {
	us.User = append(us.User, user)
	return nil
}
func (us *UserPortImpl) GetAllUsers() []model.User {
	return us.User
}

func (us *UserPortImpl) UpdateUser(updated model.User) error {
	for i, user := range us.User {
		if user.Email == updated.Email {
			us.User[i] = updated
			return nil
		}
	}
	return errors.New("user not found")
}
func (us *UserPortImpl) DeleteUser(id string) error {
	for i, user := range us.User {
		if user.ID == id {
			us.User = append(us.User[:i], us.User[i+1:]...)
			return nil
		}
	}
	return errors.New("user not found")
}
