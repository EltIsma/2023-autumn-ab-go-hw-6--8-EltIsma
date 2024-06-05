// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	models "homework/models"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// CreateDevice provides a mock function with given fields: _a0
func (_m *Repository) CreateDevice(_a0 models.Device) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Device) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteDevice provides a mock function with given fields: _a0
func (_m *Repository) DeleteDevice(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetDevice provides a mock function with given fields: _a0
func (_m *Repository) GetDevice(_a0 string) (models.Device, error) {
	ret := _m.Called(_a0)

	var r0 models.Device
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (models.Device, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(string) models.Device); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(models.Device)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateDevice provides a mock function with given fields: _a0
func (_m *Repository) UpdateDevice(_a0 models.Device) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Device) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
