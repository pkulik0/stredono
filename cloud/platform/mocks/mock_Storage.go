// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	modules "github.com/pkulik0/stredono/cloud/platform/modules"
	mock "github.com/stretchr/testify/mock"
)

// MockStorage is an autogenerated mock type for the Storage type
type MockStorage struct {
	mock.Mock
}

type MockStorage_Expecter struct {
	mock *mock.Mock
}

func (_m *MockStorage) EXPECT() *MockStorage_Expecter {
	return &MockStorage_Expecter{mock: &_m.Mock}
}

// Bucket provides a mock function with given fields: name
func (_m *MockStorage) Bucket(name string) (modules.Bucket, error) {
	ret := _m.Called(name)

	if len(ret) == 0 {
		panic("no return value specified for Bucket")
	}

	var r0 modules.Bucket
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (modules.Bucket, error)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(string) modules.Bucket); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(modules.Bucket)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStorage_Bucket_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Bucket'
type MockStorage_Bucket_Call struct {
	*mock.Call
}

// Bucket is a helper method to define mock.On call
//   - name string
func (_e *MockStorage_Expecter) Bucket(name interface{}) *MockStorage_Bucket_Call {
	return &MockStorage_Bucket_Call{Call: _e.mock.On("Bucket", name)}
}

func (_c *MockStorage_Bucket_Call) Run(run func(name string)) *MockStorage_Bucket_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockStorage_Bucket_Call) Return(_a0 modules.Bucket, _a1 error) *MockStorage_Bucket_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStorage_Bucket_Call) RunAndReturn(run func(string) (modules.Bucket, error)) *MockStorage_Bucket_Call {
	_c.Call.Return(run)
	return _c
}

// DefaultBucket provides a mock function with given fields:
func (_m *MockStorage) DefaultBucket() (modules.Bucket, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for DefaultBucket")
	}

	var r0 modules.Bucket
	var r1 error
	if rf, ok := ret.Get(0).(func() (modules.Bucket, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() modules.Bucket); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(modules.Bucket)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStorage_DefaultBucket_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DefaultBucket'
type MockStorage_DefaultBucket_Call struct {
	*mock.Call
}

// DefaultBucket is a helper method to define mock.On call
func (_e *MockStorage_Expecter) DefaultBucket() *MockStorage_DefaultBucket_Call {
	return &MockStorage_DefaultBucket_Call{Call: _e.mock.On("DefaultBucket")}
}

func (_c *MockStorage_DefaultBucket_Call) Run(run func()) *MockStorage_DefaultBucket_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockStorage_DefaultBucket_Call) Return(_a0 modules.Bucket, _a1 error) *MockStorage_DefaultBucket_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStorage_DefaultBucket_Call) RunAndReturn(run func() (modules.Bucket, error)) *MockStorage_DefaultBucket_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockStorage creates a new instance of MockStorage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockStorage(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockStorage {
	mock := &MockStorage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
