package user

import (
	"errors"
	"registration/app/model"
	"registration/app/repository/user"
	"registration/module/password"
	tk "registration/module/token"
)

type UserUsecase interface {
	Register(payload model.UserRequest) error
	Login(payload model.Login) (model.LoginResponse, error)
	GetByPhone(phone string) (model.User, error)
	UpdateName(phone string, payload model.UserUpdate) (model.User, error)
}

type userUsecase struct {
	userRepository user.UserRepository
	tokenModule    tk.TokenModule
}

func NewUserUsecase(userRepository user.UserRepository, tokenModule tk.TokenModule) UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
		tokenModule:    tokenModule,
	}
}

func (u *userUsecase) Register(payload model.UserRequest) error {
	check, err := u.userRepository.GetByPhone(payload.Phone)
	if err != nil {
		return err
	}

	if check.Name != "" {
		return errors.New("phone number user exists")
	}

	user := &model.User{
		Name:  payload.Name,
		Phone: payload.Phone,
	}

	// salt and hash password
	user.Password = password.HashAndSalt(payload.Password)
	err = u.userRepository.Register(user)
	if err != nil {
		return err
	}
	return nil
}

func (u *userUsecase) Login(payload model.Login) (model.LoginResponse, error) {
	var resp model.LoginResponse
	user, err := u.userRepository.Login(payload)
	if err != nil {
		return resp, err
	}

	resp.Token, err = u.tokenModule.GenerateTokenJWT(user)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (u *userUsecase) GetByPhone(phone string) (model.User, error) {
	var resp model.User
	resp, err := u.userRepository.GetByPhone(phone)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (u *userUsecase) UpdateName(phone string, payload model.UserUpdate) (model.User, error) {
	user, err := u.userRepository.GetByPhone(phone)
	if err != nil {
		return user, err
	}

	user.Name = payload.Name
	err = u.userRepository.UpdateName(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}
