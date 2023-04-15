// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	model "registration/app/model"

	mock "github.com/stretchr/testify/mock"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// GetByPhone provides a mock function with given fields: phone
func (_m *UserRepository) GetByPhone(phone string) (model.User, error) {
	ret := _m.Called(phone)

	var r0 model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (model.User, error)); ok {
		return rf(phone)
	}
	if rf, ok := ret.Get(0).(func(string) model.User); ok {
		r0 = rf(phone)
	} else {
		r0 = ret.Get(0).(model.User)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(phone)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: payload
func (_m *UserRepository) Login(payload model.Login) (model.User, error) {
	ret := _m.Called(payload)

	var r0 model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(model.Login) (model.User, error)); ok {
		return rf(payload)
	}
	if rf, ok := ret.Get(0).(func(model.Login) model.User); ok {
		r0 = rf(payload)
	} else {
		r0 = ret.Get(0).(model.User)
	}

	if rf, ok := ret.Get(1).(func(model.Login) error); ok {
		r1 = rf(payload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: payload
func (_m *UserRepository) Register(payload *model.User) error {
	ret := _m.Called(payload)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.User) error); ok {
		r0 = rf(payload)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateName provides a mock function with given fields: _a0
func (_m *UserRepository) UpdateName(_a0 *model.User) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.User) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewUserRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserRepository creates a new instance of UserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserRepository(t mockConstructorTestingTNewUserRepository) *UserRepository {
	mock := &UserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}