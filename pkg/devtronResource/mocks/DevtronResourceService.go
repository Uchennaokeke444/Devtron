// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	bean "github.com/devtron-labs/devtron/pkg/devtronResource/bean"

	mock "github.com/stretchr/testify/mock"
)

// DevtronResourceService is an autogenerated mock type for the DevtronResourceService type
type DevtronResourceService struct {
	mock.Mock
}

// GetAllSearchableKeyIdNameMap provides a mock function with given fields:
func (_m *DevtronResourceService) GetAllSearchableKeyIdNameMap() map[int]bean.DevtronResourceSearchableKeyName {
	ret := _m.Called()

	var r0 map[int]bean.DevtronResourceSearchableKeyName
	if rf, ok := ret.Get(0).(func() map[int]bean.DevtronResourceSearchableKeyName); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[int]bean.DevtronResourceSearchableKeyName)
		}
	}

	return r0
}

// GetAllSearchableKeyNameIdMap provides a mock function with given fields:
func (_m *DevtronResourceService) GetAllSearchableKeyNameIdMap() map[bean.DevtronResourceSearchableKeyName]int {
	ret := _m.Called()

	var r0 map[bean.DevtronResourceSearchableKeyName]int
	if rf, ok := ret.Get(0).(func() map[bean.DevtronResourceSearchableKeyName]int); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[bean.DevtronResourceSearchableKeyName]int)
		}
	}

	return r0
}

type mockConstructorTestingTNewDevtronResourceService interface {
	mock.TestingT
	Cleanup(func())
}

// NewDevtronResourceService creates a new instance of DevtronResourceService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDevtronResourceService(t mockConstructorTestingTNewDevtronResourceService) *DevtronResourceService {
	mock := &DevtronResourceService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
