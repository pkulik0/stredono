// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	platform "github.com/pkulik0/stredono/cloud/platform"
	mock "github.com/stretchr/testify/mock"
)

// MockCollection is an autogenerated mock type for the Collection type
type MockCollection struct {
	mock.Mock
}

type MockCollection_Expecter struct {
	mock *mock.Mock
}

func (_m *MockCollection) EXPECT() *MockCollection_Expecter {
	return &MockCollection_Expecter{mock: &_m.Mock}
}

// Doc provides a mock function with given fields: path
func (_m *MockCollection) Doc(path string) platform.Document {
	ret := _m.Called(path)

	if len(ret) == 0 {
		panic("no return value specified for Doc")
	}

	var r0 platform.Document
	if rf, ok := ret.Get(0).(func(string) platform.Document); ok {
		r0 = rf(path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(platform.Document)
		}
	}

	return r0
}

// MockCollection_Doc_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Doc'
type MockCollection_Doc_Call struct {
	*mock.Call
}

// Doc is a helper method to define mock.On call
//   - path string
func (_e *MockCollection_Expecter) Doc(path interface{}) *MockCollection_Doc_Call {
	return &MockCollection_Doc_Call{Call: _e.mock.On("Doc", path)}
}

func (_c *MockCollection_Doc_Call) Run(run func(path string)) *MockCollection_Doc_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockCollection_Doc_Call) Return(_a0 platform.Document) *MockCollection_Doc_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCollection_Doc_Call) RunAndReturn(run func(string) platform.Document) *MockCollection_Doc_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockCollection creates a new instance of MockCollection. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockCollection(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockCollection {
	mock := &MockCollection{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
