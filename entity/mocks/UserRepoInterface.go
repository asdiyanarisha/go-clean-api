// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	entity "go-clean-api/entity"

	mock "github.com/stretchr/testify/mock"
)

// UserRepoInterface is an autogenerated mock type for the UserRepoInterface type
type UserRepoInterface struct {
	mock.Mock
}

// GetByEmail provides a mock function with given fields: Username
func (_m *UserRepoInterface) GetByEmail(Username string) entity.User {
	ret := _m.Called(Username)

	var r0 entity.User
	if rf, ok := ret.Get(0).(func(string) entity.User); ok {
		r0 = rf(Username)
	} else {
		r0 = ret.Get(0).(entity.User)
	}

	return r0
}

// GetByUsername provides a mock function with given fields: Username
func (_m *UserRepoInterface) GetByUsername(Username string) entity.User {
	ret := _m.Called(Username)

	var r0 entity.User
	if rf, ok := ret.Get(0).(func(string) entity.User); ok {
		r0 = rf(Username)
	} else {
		r0 = ret.Get(0).(entity.User)
	}

	return r0
}

// GetByUsernameEmail provides a mock function with given fields: Username, Password
func (_m *UserRepoInterface) GetByUsernameEmail(Username string, Password string) entity.User {
	ret := _m.Called(Username, Password)

	var r0 entity.User
	if rf, ok := ret.Get(0).(func(string, string) entity.User); ok {
		r0 = rf(Username, Password)
	} else {
		r0 = ret.Get(0).(entity.User)
	}

	return r0
}

// StoreUser provides a mock function with given fields: user
func (_m *UserRepoInterface) StoreUser(user *entity.User) {
	_m.Called(user)
}

type mockConstructorTestingTNewUserRepoInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserRepoInterface creates a new instance of UserRepoInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserRepoInterface(t mockConstructorTestingTNewUserRepoInterface) *UserRepoInterface {
	mock := &UserRepoInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}