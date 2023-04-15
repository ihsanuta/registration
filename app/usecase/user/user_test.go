package user

import (
	"reflect"
	"registration/app/model"
	"registration/app/repository/user/mocks"
	mocktoken "registration/module/token/mocks"
	"testing"

	"github.com/stretchr/testify/mock"
)

func Test_userUsecase_Register(t *testing.T) {
	mockrepo := mocks.NewUserRepository(t)
	type fields struct {
		userRepository *mocks.UserRepository
	}
	type args struct {
		payload model.UserRequest
	}
	tests := []struct {
		name     string
		fields   fields
		mockfunc func(userRepo *mocks.UserRepository, payload model.UserRequest)
		args     args
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			name: "Success Registrasi",
			fields: fields{
				userRepository: mockrepo,
			},
			mockfunc: func(userRepo *mocks.UserRepository, payload model.UserRequest) {
				userRepo.On("GetByPhone", payload.Phone).Return(
					model.User{},
					nil,
				)

				userRepo.On("Register", mock.Anything).Return(
					nil,
				)
			},
			args: args{
				payload: model.UserRequest{
					Name:     "coba",
					Phone:    "081366677788",
					Password: "testpassword",
				},
			},
			wantErr: false,
		},
		{
			name: "Phone Exists Registrasi",
			fields: fields{
				userRepository: mockrepo,
			},
			mockfunc: func(userRepo *mocks.UserRepository, payload model.UserRequest) {
				userRepo.On("GetByPhone", payload.Phone).Return(
					model.User{
						Name: "test",
					},
					nil,
				)
			},
			args: args{
				payload: model.UserRequest{
					Name:     "coba",
					Phone:    "08129120000",
					Password: "testpassword",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &userUsecase{
				userRepository: tt.fields.userRepository,
			}

			tt.mockfunc(tt.fields.userRepository, tt.args.payload)

			if err := u.Register(tt.args.payload); (err != nil) != tt.wantErr {
				t.Errorf("userUsecase.Register() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_userUsecase_Login(t *testing.T) {
	mockrepo := mocks.NewUserRepository(t)
	mockToken := mocktoken.NewTokenModule(t)
	type fields struct {
		userRepository *mocks.UserRepository
		tokenModule    *mocktoken.TokenModule
	}
	type args struct {
		payload model.Login
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		mockfunc func(userRepo *mocks.UserRepository, tokenModule *mocktoken.TokenModule, payload model.Login)
		want     model.LoginResponse
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			name: "LOGIN SUCCESS",
			fields: fields{
				userRepository: mockrepo,
				tokenModule:    mockToken,
			},
			mockfunc: func(userRepo *mocks.UserRepository, tokenModule *mocktoken.TokenModule, payload model.Login) {
				userRepo.On("Login", payload).Return(
					model.User{
						Name:  "test",
						Phone: "08128129999",
					},
					nil,
				)

				tokenModule.On("GenerateTokenJWT", model.User{
					Phone: "08128129999",
					Name:  "test",
				}).Return(
					mock.Anything,
					nil,
				)
			},
			args: args{
				payload: model.Login{
					Phone:    "08128129999",
					Password: "12345678",
				},
			},
			want: model.LoginResponse{
				Token: "",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &userUsecase{
				userRepository: tt.fields.userRepository,
				tokenModule:    tt.fields.tokenModule,
			}

			tt.mockfunc(tt.fields.userRepository, tt.fields.tokenModule, tt.args.payload)

			_, err := u.Login(tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("userUsecase.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("userUsecase.Login() = %v, want %v", got, tt.want)
			// }
		})
	}
}

func Test_userUsecase_GetByPhone(t *testing.T) {
	mockrepo := mocks.NewUserRepository(t)
	type fields struct {
		userRepository *mocks.UserRepository
	}
	type args struct {
		phone string
	}
	tests := []struct {
		name     string
		fields   fields
		mockfunc func(userRepo *mocks.UserRepository, phone string)
		args     args
		want     model.User
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			name: "Success Get Data",
			fields: fields{
				userRepository: mockrepo,
			},
			args: args{
				phone: "081377281999",
			},
			want: model.User{
				ID:    1,
				Name:  "test",
				Phone: "081377281999",
			},
			wantErr: false,
			mockfunc: func(userRepo *mocks.UserRepository, phone string) {
				userRepo.On("GetByPhone", phone).Return(
					model.User{
						ID:    1,
						Name:  "test",
						Phone: "081377281999",
					},
					nil,
				)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &userUsecase{
				userRepository: tt.fields.userRepository,
			}
			tt.mockfunc(tt.fields.userRepository, tt.args.phone)
			got, err := u.GetByPhone(tt.args.phone)
			if (err != nil) != tt.wantErr {
				t.Errorf("userUsecase.GetByPhone() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userUsecase.GetByPhone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userUsecase_UpdateName(t *testing.T) {
	mockrepo := mocks.NewUserRepository(t)
	type fields struct {
		userRepository *mocks.UserRepository
	}
	type args struct {
		phone   string
		payload model.UserUpdate
	}
	tests := []struct {
		name     string
		fields   fields
		mockfunc func(userRepo *mocks.UserRepository, phone string)
		args     args
		want     model.User
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			name: "SUCCESS UPDATE",
			fields: fields{
				userRepository: mockrepo,
			},
			args: args{
				phone: "0812912999",
				payload: model.UserUpdate{
					Name: "nama baru",
				},
			},
			want: model.User{
				ID:    1,
				Phone: "0812912999",
				Name:  "nama baru",
			},
			wantErr: false,
			mockfunc: func(userRepo *mocks.UserRepository, phone string) {
				userRepo.On("GetByPhone", phone).Return(
					model.User{
						ID:    1,
						Name:  "nama lama",
						Phone: "0812912999",
					},
					nil,
				)

				userRepo.On("UpdateName", &model.User{
					ID:    1,
					Name:  "nama baru",
					Phone: "0812912999",
				}).Return(
					nil,
				)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &userUsecase{
				userRepository: tt.fields.userRepository,
			}
			tt.mockfunc(tt.fields.userRepository, tt.args.phone)
			got, err := u.UpdateName(tt.args.phone, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("userUsecase.UpdateName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userUsecase.UpdateName() = %v, want %v", got, tt.want)
			}
		})
	}
}
