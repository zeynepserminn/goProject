package repositories

import (
	"errors"
	"goProject/internal/core/dto"
	"goProject/internal/core/model"
	"goProject/pkg/postgres"
	"gorm.io/gorm"
)

type UserPortImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserPortImpl {
	return &UserPortImpl{db: db}
}

func (us *UserPortImpl) AddUser(user model.User) (int64, error) {
	if us.db == nil {
		return 0, errors.New("database connection is nil")
	}
	err := us.db.Create(&user).Error
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}

func (us *UserPortImpl) GetAllUsers(pg dto.PaginationRequest, filters dto.FilterParams) ([]model.User, int, error) {
	var users []model.User
	var total int64

	query := us.db.Model(&users)

	if filters.Firstname != nil {
		query.Where("firstname ILIKE ?", "%"+*filters.Firstname+"%")
	}
	if filters.Lastname != nil {
		query.Where("lastname ILIKE ?", "%"+*filters.Lastname+"%")
	}
	if filters.Email != nil {
		query.Where("email ILIKE ?", "%"+*filters.Email+"%")
	}
	if filters.Phone != nil {
		query.Where("phone ILIKE ?", "%"+*filters.Phone+"%")
	}
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Scopes(postgres.PaginatedResult(pg)).Find(&users).Error
	if err != nil {
		return nil, 0, err
	}

	return users, int(total), nil
}

func (us *UserPortImpl) UpdateUser(updated model.User) error {
	err := us.db.Model(&model.User{}).Where("id=?", updated.ID).Updates(updated).Error

	return err
}
func (us *UserPortImpl) DeleteUser(id int32) error {
	err := us.db.Model(&model.User{}).Where("id=?", id).Update("status", model.Deleted).Error

	return err
}
func (us *UserPortImpl) GetUserByID(ID int32) (*model.User, error) {
	var user model.User
	err := us.db.Where("id=? AND status=?", ID, model.Active).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
func (us *UserPortImpl) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	err := us.db.Where("email=? AND status=?", email, model.Active).First(&user).Error
	if err != nil {

		return nil, err
	}
	return &user, nil
}

func (us *UserPortImpl) IsEmailExists(email string, exclude *int64) (bool, error) {
	var count int64

	query := us.db.Model(&model.User{}).Where("email=?", email)
	if exclude != nil {
		query.Where("id!=?", *exclude)
	}
	err := query.Count(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (us *UserPortImpl) IsPhoneExists(phone string) (bool, error) {
	var count int64
	err := us.db.Model(&model.User{}).Where("phone=?", phone).Count(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (us *UserPortImpl) UpdateUserStatus(userID int32, status int) error {
	err := us.db.Model(&model.User{}).Where("id=?", userID).Update("status", status).Error

	return err
}
