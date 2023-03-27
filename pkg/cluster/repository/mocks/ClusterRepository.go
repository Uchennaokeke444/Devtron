// Code generated by mockery v2.18.0. DO NOT EDIT.

package mocks

import (
	repository "github.com/devtron-labs/devtron/pkg/cluster/repository"
	mock "github.com/stretchr/testify/mock"
)

// ClusterRepository is an autogenerated mock type for the ClusterRepository type
type ClusterRepository struct {
	mock.Mock
}

// Delete provides a mock function with given fields: model
func (_m *ClusterRepository) Delete(model *repository.Cluster) error {
	ret := _m.Called(model)

	var r0 error
	if rf, ok := ret.Get(0).(func(*repository.Cluster) error); ok {
		r0 = rf(model)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindAll provides a mock function with given fields:
func (_m *ClusterRepository) FindAll() ([]repository.Cluster, error) {
	ret := _m.Called()

	var r0 []repository.Cluster
	if rf, ok := ret.Get(0).(func() []repository.Cluster); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]repository.Cluster)
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

// FindAllActive provides a mock function with given fields:
func (_m *ClusterRepository) FindAllActive() ([]repository.Cluster, error) {
	ret := _m.Called()

	var r0 []repository.Cluster
	if rf, ok := ret.Get(0).(func() []repository.Cluster); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]repository.Cluster)
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

// FindById provides a mock function with given fields: id
func (_m *ClusterRepository) FindById(id int) (*repository.Cluster, error) {
	ret := _m.Called(id)

	var r0 *repository.Cluster
	if rf, ok := ret.Get(0).(func(int) *repository.Cluster); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*repository.Cluster)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByIds provides a mock function with given fields: id
func (_m *ClusterRepository) FindByIds(id []int) ([]repository.Cluster, error) {
	ret := _m.Called(id)

	var r0 []repository.Cluster
	if rf, ok := ret.Get(0).(func([]int) []repository.Cluster); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]repository.Cluster)
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

// FindOne provides a mock function with given fields: clusterName
func (_m *ClusterRepository) FindOne(clusterName string) (*repository.Cluster, error) {
	ret := _m.Called(clusterName)

	var r0 *repository.Cluster
	if rf, ok := ret.Get(0).(func(string) *repository.Cluster); ok {
		r0 = rf(clusterName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*repository.Cluster)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(clusterName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindOneActive provides a mock function with given fields: clusterName
func (_m *ClusterRepository) FindOneActive(clusterName string) (*repository.Cluster, error) {
	ret := _m.Called(clusterName)

	var r0 *repository.Cluster
	if rf, ok := ret.Get(0).(func(string) *repository.Cluster); ok {
		r0 = rf(clusterName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*repository.Cluster)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(clusterName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MarkClusterDeleted provides a mock function with given fields: model
func (_m *ClusterRepository) MarkClusterDeleted(model *repository.Cluster) error {
	ret := _m.Called(model)

	var r0 error
	if rf, ok := ret.Get(0).(func(*repository.Cluster) error); ok {
		r0 = rf(model)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Save provides a mock function with given fields: model
func (_m *ClusterRepository) Save(model *repository.Cluster) error {
	ret := _m.Called(model)

	var r0 error
	if rf, ok := ret.Get(0).(func(*repository.Cluster) error); ok {
		r0 = rf(model)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: model
func (_m *ClusterRepository) Update(model *repository.Cluster) error {
	ret := _m.Called(model)

	var r0 error
	if rf, ok := ret.Get(0).(func(*repository.Cluster) error); ok {
		r0 = rf(model)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateClusterConnectionStatus provides a mock function with given fields: clusterId, errorInConnecting
func (_m *ClusterRepository) UpdateClusterConnectionStatus(clusterId int, errorInConnecting string) error {
	ret := _m.Called(clusterId, errorInConnecting)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, string) error); ok {
		r0 = rf(clusterId, errorInConnecting)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewClusterRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewClusterRepository creates a new instance of ClusterRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewClusterRepository(t mockConstructorTestingTNewClusterRepository) *ClusterRepository {
	mock := &ClusterRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}