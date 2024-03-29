// Code generated by mockery v2.18.0. DO NOT EDIT.

package mocks

import (
	bytes "bytes"

	terminal "github.com/devtron-labs/devtron/pkg/terminal"
	mock "github.com/stretchr/testify/mock"
)

// TerminalSessionHandler is an autogenerated mock type for the TerminalSessionHandler type
type TerminalSessionHandler struct {
	mock.Mock
}

// AutoSelectShell provides a mock function with given fields: req
func (_m *TerminalSessionHandler) AutoSelectShell(req *terminal.TerminalSessionRequest) (string, error) {
	ret := _m.Called(req)

	var r0 string
	if rf, ok := ret.Get(0).(func(*terminal.TerminalSessionRequest) string); ok {
		r0 = rf(req)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*terminal.TerminalSessionRequest) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Close provides a mock function with given fields: sessionId, statusCode, msg
func (_m *TerminalSessionHandler) Close(sessionId string, statusCode uint32, msg string) {
	_m.Called(sessionId, statusCode, msg)
}

// GetTerminalSession provides a mock function with given fields: req
func (_m *TerminalSessionHandler) GetTerminalSession(req *terminal.TerminalSessionRequest) (int, *terminal.TerminalMessage, error) {
	ret := _m.Called(req)

	var r0 int
	if rf, ok := ret.Get(0).(func(*terminal.TerminalSessionRequest) int); ok {
		r0 = rf(req)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 *terminal.TerminalMessage
	if rf, ok := ret.Get(1).(func(*terminal.TerminalSessionRequest) *terminal.TerminalMessage); ok {
		r1 = rf(req)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*terminal.TerminalMessage)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*terminal.TerminalSessionRequest) error); ok {
		r2 = rf(req)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// RunCmdInRemotePod provides a mock function with given fields: req, cmds
func (_m *TerminalSessionHandler) RunCmdInRemotePod(req *terminal.TerminalSessionRequest, cmds []string) (*bytes.Buffer, *bytes.Buffer, error) {
	ret := _m.Called(req, cmds)

	var r0 *bytes.Buffer
	if rf, ok := ret.Get(0).(func(*terminal.TerminalSessionRequest, []string) *bytes.Buffer); ok {
		r0 = rf(req, cmds)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bytes.Buffer)
		}
	}

	var r1 *bytes.Buffer
	if rf, ok := ret.Get(1).(func(*terminal.TerminalSessionRequest, []string) *bytes.Buffer); ok {
		r1 = rf(req, cmds)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*bytes.Buffer)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*terminal.TerminalSessionRequest, []string) error); ok {
		r2 = rf(req, cmds)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ValidateSession provides a mock function with given fields: sessionId
func (_m *TerminalSessionHandler) ValidateSession(sessionId string) bool {
	ret := _m.Called(sessionId)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(sessionId)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ValidateShell provides a mock function with given fields: req
func (_m *TerminalSessionHandler) ValidateShell(req *terminal.TerminalSessionRequest) (bool, error) {
	ret := _m.Called(req)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*terminal.TerminalSessionRequest) bool); ok {
		r0 = rf(req)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*terminal.TerminalSessionRequest) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewTerminalSessionHandler interface {
	mock.TestingT
	Cleanup(func())
}

// NewTerminalSessionHandler creates a new instance of TerminalSessionHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTerminalSessionHandler(t mockConstructorTestingTNewTerminalSessionHandler) *TerminalSessionHandler {
	mock := &TerminalSessionHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
