// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	modules "github.com/pkulik0/stredono/cloud/platform/modules"
	mock "github.com/stretchr/testify/mock"
)

// MockQueryIterator is an autogenerated mock type for the QueryIterator type
type MockQueryIterator struct {
	mock.Mock
}

type MockQueryIterator_Expecter struct {
	mock *mock.Mock
}

func (_m *MockQueryIterator) EXPECT() *MockQueryIterator_Expecter {
	return &MockQueryIterator_Expecter{mock: &_m.Mock}
}

// All provides a mock function with given fields:
func (_m *MockQueryIterator) All() ([]modules.DocumentSnapshot, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for All")
	}

	var r0 []modules.DocumentSnapshot
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]modules.DocumentSnapshot, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []modules.DocumentSnapshot); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]modules.DocumentSnapshot)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQueryIterator_All_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'All'
type MockQueryIterator_All_Call struct {
	*mock.Call
}

// All is a helper method to define mock.On call
func (_e *MockQueryIterator_Expecter) All() *MockQueryIterator_All_Call {
	return &MockQueryIterator_All_Call{Call: _e.mock.On("All")}
}

func (_c *MockQueryIterator_All_Call) Run(run func()) *MockQueryIterator_All_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockQueryIterator_All_Call) Return(_a0 []modules.DocumentSnapshot, _a1 error) *MockQueryIterator_All_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockQueryIterator_All_Call) RunAndReturn(run func() ([]modules.DocumentSnapshot, error)) *MockQueryIterator_All_Call {
	_c.Call.Return(run)
	return _c
}

// Next provides a mock function with given fields:
func (_m *MockQueryIterator) Next() (modules.DocumentSnapshot, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Next")
	}

	var r0 modules.DocumentSnapshot
	var r1 error
	if rf, ok := ret.Get(0).(func() (modules.DocumentSnapshot, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() modules.DocumentSnapshot); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(modules.DocumentSnapshot)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQueryIterator_Next_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Next'
type MockQueryIterator_Next_Call struct {
	*mock.Call
}

// Next is a helper method to define mock.On call
func (_e *MockQueryIterator_Expecter) Next() *MockQueryIterator_Next_Call {
	return &MockQueryIterator_Next_Call{Call: _e.mock.On("Next")}
}

func (_c *MockQueryIterator_Next_Call) Run(run func()) *MockQueryIterator_Next_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockQueryIterator_Next_Call) Return(_a0 modules.DocumentSnapshot, _a1 error) *MockQueryIterator_Next_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockQueryIterator_Next_Call) RunAndReturn(run func() (modules.DocumentSnapshot, error)) *MockQueryIterator_Next_Call {
	_c.Call.Return(run)
	return _c
}

// Stop provides a mock function with given fields:
func (_m *MockQueryIterator) Stop() {
	_m.Called()
}

// MockQueryIterator_Stop_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Stop'
type MockQueryIterator_Stop_Call struct {
	*mock.Call
}

// Stop is a helper method to define mock.On call
func (_e *MockQueryIterator_Expecter) Stop() *MockQueryIterator_Stop_Call {
	return &MockQueryIterator_Stop_Call{Call: _e.mock.On("Stop")}
}

func (_c *MockQueryIterator_Stop_Call) Run(run func()) *MockQueryIterator_Stop_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockQueryIterator_Stop_Call) Return() *MockQueryIterator_Stop_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockQueryIterator_Stop_Call) RunAndReturn(run func()) *MockQueryIterator_Stop_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockQueryIterator creates a new instance of MockQueryIterator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockQueryIterator(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockQueryIterator {
	mock := &MockQueryIterator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
