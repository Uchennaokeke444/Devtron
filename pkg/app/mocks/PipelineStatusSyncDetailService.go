// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// PipelineStatusSyncDetailService is an autogenerated mock type for the PipelineStatusSyncDetailService type
type PipelineStatusSyncDetailService struct {
	mock.Mock
}

// GetSyncTimeAndCountByCdWfrId provides a mock function with given fields: cdWfrId
func (_m *PipelineStatusSyncDetailService) GetSyncTimeAndCountByCdWfrId(cdWfrId int) (time.Time, int, error) {
	ret := _m.Called(cdWfrId)

	var r0 time.Time
	if rf, ok := ret.Get(0).(func(int) time.Time); ok {
		r0 = rf(cdWfrId)
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(int) int); ok {
		r1 = rf(cdWfrId)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(int) error); ok {
		r2 = rf(cdWfrId)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// SaveOrUpdateSyncDetail provides a mock function with given fields: cdWfrId, userId
func (_m *PipelineStatusSyncDetailService) SaveOrUpdateSyncDetail(cdWfrId int, userId int32) error {
	ret := _m.Called(cdWfrId, userId)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int32) error); ok {
		r0 = rf(cdWfrId, userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewPipelineStatusSyncDetailService interface {
	mock.TestingT
	Cleanup(func())
}

// NewPipelineStatusSyncDetailService creates a new instance of PipelineStatusSyncDetailService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPipelineStatusSyncDetailService(t mockConstructorTestingTNewPipelineStatusSyncDetailService) *PipelineStatusSyncDetailService {
	mock := &PipelineStatusSyncDetailService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
