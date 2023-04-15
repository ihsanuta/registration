package repository

import (
	"registration/app/repository/user"

	"github.com/jinzhu/gorm"
)

type Repository struct {
	User user.UserRepository
}

func Init(db *gorm.DB) *Repository {
	repo := &Repository{
		User: user.NewUserRepository(
			db,
		),
	}
	return repo
}
