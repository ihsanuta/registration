package usecase

import (
	"registration/app/repository"
	"registration/app/usecase/user"
	"registration/module/token"
)

type Usecase struct {
	User user.UserUsecase
}

func Init(repository *repository.Repository) *Usecase {
	uc := &Usecase{
		User: user.NewUserUsecase(
			repository.User,
			token.NewTokenModule(),
		),
	}
	return uc
}
