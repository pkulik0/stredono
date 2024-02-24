// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	modules "github.com/pkulik0/stredono/cloud/platform/modules"
	mock "github.com/stretchr/testify/mock"
)

// MockPubSubClient is an autogenerated mock type for the PubSubClient type
type MockPubSubClient struct {
	mock.Mock
}

type MockPubSubClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPubSubClient) EXPECT() *MockPubSubClient_Expecter {
	return &MockPubSubClient_Expecter{mock: &_m.Mock}
}

// Stop provides a mock function with given fields:
func (_m *MockPubSubClient) Stop() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Stop")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockPubSubClient_Stop_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Stop'
type MockPubSubClient_Stop_Call struct {
	*mock.Call
}

// Stop is a helper method to define mock.On call
func (_e *MockPubSubClient_Expecter) Stop() *MockPubSubClient_Stop_Call {
	return &MockPubSubClient_Stop_Call{Call: _e.mock.On("Stop")}
}

func (_c *MockPubSubClient_Stop_Call) Run(run func()) *MockPubSubClient_Stop_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPubSubClient_Stop_Call) Return(_a0 error) *MockPubSubClient_Stop_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPubSubClient_Stop_Call) RunAndReturn(run func() error) *MockPubSubClient_Stop_Call {
	_c.Call.Return(run)
	return _c
}

// Topic provides a mock function with given fields: name
func (_m *MockPubSubClient) Topic(name string) modules.PubSubTopic {
	ret := _m.Called(name)

	if len(ret) == 0 {
		panic("no return value specified for Topic")
	}

	var r0 modules.PubSubTopic
	if rf, ok := ret.Get(0).(func(string) modules.PubSubTopic); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(modules.PubSubTopic)
		}
	}

	return r0
}

// MockPubSubClient_Topic_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Topic'
type MockPubSubClient_Topic_Call struct {
	*mock.Call
}

// Topic is a helper method to define mock.On call
//   - name string
func (_e *MockPubSubClient_Expecter) Topic(name interface{}) *MockPubSubClient_Topic_Call {
	return &MockPubSubClient_Topic_Call{Call: _e.mock.On("Topic", name)}
}

func (_c *MockPubSubClient_Topic_Call) Run(run func(name string)) *MockPubSubClient_Topic_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockPubSubClient_Topic_Call) Return(_a0 modules.PubSubTopic) *MockPubSubClient_Topic_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPubSubClient_Topic_Call) RunAndReturn(run func(string) modules.PubSubTopic) *MockPubSubClient_Topic_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockPubSubClient creates a new instance of MockPubSubClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPubSubClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPubSubClient {
	mock := &MockPubSubClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
