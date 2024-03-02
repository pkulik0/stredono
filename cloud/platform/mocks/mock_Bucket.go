// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	modules "github.com/pkulik0/stredono/cloud/platform/modules"
	mock "github.com/stretchr/testify/mock"
)

// MockBucket is an autogenerated mock type for the Bucket type
type MockBucket struct {
	mock.Mock
}

type MockBucket_Expecter struct {
	mock *mock.Mock
}

func (_m *MockBucket) EXPECT() *MockBucket_Expecter {
	return &MockBucket_Expecter{mock: &_m.Mock}
}

// Object provides a mock function with given fields: name
func (_m *MockBucket) Object(name string) modules.Object {
	ret := _m.Called(name)

	if len(ret) == 0 {
		panic("no return value specified for Object")
	}

	var r0 modules.Object
	if rf, ok := ret.Get(0).(func(string) modules.Object); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(modules.Object)
		}
	}

	return r0
}

// MockBucket_Object_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Object'
type MockBucket_Object_Call struct {
	*mock.Call
}

// Object is a helper method to define mock.On call
//   - name string
func (_e *MockBucket_Expecter) Object(name interface{}) *MockBucket_Object_Call {
	return &MockBucket_Object_Call{Call: _e.mock.On("Object", name)}
}

func (_c *MockBucket_Object_Call) Run(run func(name string)) *MockBucket_Object_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockBucket_Object_Call) Return(_a0 modules.Object) *MockBucket_Object_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBucket_Object_Call) RunAndReturn(run func(string) modules.Object) *MockBucket_Object_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockBucket creates a new instance of MockBucket. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockBucket(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockBucket {
	mock := &MockBucket{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}