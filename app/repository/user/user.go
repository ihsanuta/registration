//go:generate mockery --name=UserRepository
package user

import (
	"registration/app/model"

	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	Register(payload *model.User) error
	Login(payload model.Login) (model.User, error)
	GetByPhone(phone string) (model.User, error)
	UpdateName(user *model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) Register(payload *model.User) error {
	db := u.db.Create(&payload)
	if db.Error != nil {
		return db.Error
	}

	return nil
}

func (u *userRepository) Login(payload model.Login) (model.User, error) {
	var user model.User
	db := u.db.Where("phone = ? AND password = ?", payload.Phone, payload.Password).Find(&user)
	if db.Error != nil {
		return user, db.Error
	}

	return user, nil
}

func (u *userRepository) GetByPhone(phone string) (model.User, error) {
	var user model.User
	db := u.db.Where("phone = ?", phone).Find(&user)
	if db.Error != nil {
		return user, db.Error
	}

	return user, nil
}

func (u *userRepository) UpdateName(user *model.User) error {
	db := u.db.Save(user)
	if db.Error != nil {
		return db.Error
	}

	return nil
}
