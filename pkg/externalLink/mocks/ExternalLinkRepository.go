// Code generated by mockery v2.14.1. DO NOT EDIT.

package mocks

import (
	externalLink "github.com/devtron-labs/devtron/pkg/externalLink"
	mock "github.com/stretchr/testify/mock"

	pg "github.com/go-pg/pg"
)

// ExternalLinkRepository is an autogenerated mock type for the ExternalLinkRepository type
type ExternalLinkRepository struct {
	mock.Mock
}

// FindAllActive provides a mock function with given fields:
func (_m *ExternalLinkRepository) FindAllActive() ([]externalLink.ExternalLink, error) {
	ret := _m.Called()

	var r0 []externalLink.ExternalLink
	if rf, ok := ret.Get(0).(func() []externalLink.ExternalLink); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]externalLink.ExternalLink)
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

// FindAllClusterLinks provides a mock function with given fields:
func (_m *ExternalLinkRepository) FindAllClusterLinks() ([]externalLink.ExternalLink, error) {
	ret := _m.Called()

	var r0 []externalLink.ExternalLink
	if rf, ok := ret.Get(0).(func() []externalLink.ExternalLink); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]externalLink.ExternalLink)
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

// FindAllFilterOutByIds provides a mock function with given fields: ids
func (_m *ExternalLinkRepository) FindAllFilterOutByIds(ids []int) ([]externalLink.ExternalLink, error) {
	ret := _m.Called(ids)

	var r0 []externalLink.ExternalLink
	if rf, ok := ret.Get(0).(func([]int) []externalLink.ExternalLink); ok {
		r0 = rf(ids)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]externalLink.ExternalLink)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]int) error); ok {
		r1 = rf(ids)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindOne provides a mock function with given fields: id
func (_m *ExternalLinkRepository) FindOne(id int) (externalLink.ExternalLink, error) {
	ret := _m.Called(id)

	var r0 externalLink.ExternalLink
	if rf, ok := ret.Get(0).(func(int) externalLink.ExternalLink); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(externalLink.ExternalLink)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetConnection provides a mock function with given fields:
func (_m *ExternalLinkRepository) GetConnection() *pg.DB {
	ret := _m.Called()

	var r0 *pg.DB
	if rf, ok := ret.Get(0).(func() *pg.DB); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pg.DB)
		}
	}

	return r0
}

// Save provides a mock function with given fields: externalLinks, tx
func (_m *ExternalLinkRepository) Save(externalLinks *externalLink.ExternalLink, tx *pg.Tx) error {
	ret := _m.Called(externalLinks, tx)

	var r0 error
	if rf, ok := ret.Get(0).(func(*externalLink.ExternalLink, *pg.Tx) error); ok {
		r0 = rf(externalLinks, tx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: link, tx
func (_m *ExternalLinkRepository) Update(link *externalLink.ExternalLink, tx *pg.Tx) error {
	ret := _m.Called(link, tx)

	var r0 error
	if rf, ok := ret.Get(0).(func(*externalLink.ExternalLink, *pg.Tx) error); ok {
		r0 = rf(link, tx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewExternalLinkRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewExternalLinkRepository creates a new instance of ExternalLinkRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewExternalLinkRepository(t mockConstructorTestingTNewExternalLinkRepository) *ExternalLinkRepository {
	mock := &ExternalLinkRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
