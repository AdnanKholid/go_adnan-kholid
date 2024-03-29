// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	users "weakly-task-3/businesses/users"

	mock "github.com/stretchr/testify/mock"
)

// Usecase is an autogenerated mock type for the Usecase type
type Usecase struct {
	mock.Mock
}

// Login provides a mock function with given fields: userDomain
func (_m *Usecase) Login(userDomain *users.Domain) string {
	ret := _m.Called(userDomain)

	var r0 string
	if rf, ok := ret.Get(0).(func(*users.Domain) string); ok {
		r0 = rf(userDomain)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Register provides a mock function with given fields: userDomain
func (_m *Usecase) Register(userDomain *users.Domain) users.Domain {
	ret := _m.Called(userDomain)

	var r0 users.Domain
	if rf, ok := ret.Get(0).(func(*users.Domain) users.Domain); ok {
		r0 = rf(userDomain)
	} else {
		r0 = ret.Get(0).(users.Domain)
	}

	return r0
}

type mockConstructorTestingTNewUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewUsecase creates a new instance of Usecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUsecase(t mockConstructorTestingTNewUsecase) *Usecase {
	mock := &Usecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
