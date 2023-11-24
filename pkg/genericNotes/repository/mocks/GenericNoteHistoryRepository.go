// Code generated by mockery v2.18.0. DO NOT EDIT.

package mocks

import (
	pg "github.com/go-pg/pg"
	mock "github.com/stretchr/testify/mock"

	repository "github.com/devtron-labs/devtron/pkg/genericNotes/repository"
)

// GenericNoteHistoryRepository is an autogenerated mock type for the GenericNoteHistoryRepository type
type GenericNoteHistoryRepository struct {
	mock.Mock
}

// CommitTx provides a mock function with given fields: tx
func (_m *GenericNoteHistoryRepository) CommitTx(tx *pg.Tx) error {
	ret := _m.Called(tx)

	var r0 error
	if rf, ok := ret.Get(0).(func(*pg.Tx) error); ok {
		r0 = rf(tx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindHistoryByNoteId provides a mock function with given fields: id
func (_m *GenericNoteHistoryRepository) FindHistoryByNoteId(id []int) ([]repository.GenericNoteHistory, error) {
	ret := _m.Called(id)

	var r0 []repository.GenericNoteHistory
	if rf, ok := ret.Get(0).(func([]int) []repository.GenericNoteHistory); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]repository.GenericNoteHistory)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RollbackTx provides a mock function with given fields: tx
func (_m *GenericNoteHistoryRepository) RollbackTx(tx *pg.Tx) error {
	ret := _m.Called(tx)

	var r0 error
	if rf, ok := ret.Get(0).(func(*pg.Tx) error); ok {
		r0 = rf(tx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveHistory provides a mock function with given fields: tx, model
func (_m *GenericNoteHistoryRepository) SaveHistory(tx *pg.Tx, model *repository.GenericNoteHistory) error {
	ret := _m.Called(tx, model)

	var r0 error
	if rf, ok := ret.Get(0).(func(*pg.Tx, *repository.GenericNoteHistory) error); ok {
		r0 = rf(tx, model)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StartTx provides a mock function with given fields:
func (_m *GenericNoteHistoryRepository) StartTx() (*pg.Tx, error) {
	ret := _m.Called()

	var r0 *pg.Tx
	if rf, ok := ret.Get(0).(func() *pg.Tx); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pg.Tx)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewGenericNoteHistoryRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewGenericNoteHistoryRepository creates a new instance of GenericNoteHistoryRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGenericNoteHistoryRepository(t mockConstructorTestingTNewGenericNoteHistoryRepository) *GenericNoteHistoryRepository {
	mock := &GenericNoteHistoryRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
