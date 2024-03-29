// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	context "context"

	modules "github.com/pkulik0/stredono/cloud/platform/modules"
	mock "github.com/stretchr/testify/mock"
)

// MockRef is an autogenerated mock type for the Ref type
type MockRef struct {
	mock.Mock
}

type MockRef_Expecter struct {
	mock *mock.Mock
}

func (_m *MockRef) EXPECT() *MockRef_Expecter {
	return &MockRef_Expecter{mock: &_m.Mock}
}

// Child provides a mock function with given fields: path
func (_m *MockRef) Child(path string) modules.Ref {
	ret := _m.Called(path)

	if len(ret) == 0 {
		panic("no return value specified for Child")
	}

	var r0 modules.Ref
	if rf, ok := ret.Get(0).(func(string) modules.Ref); ok {
		r0 = rf(path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(modules.Ref)
		}
	}

	return r0
}

// MockRef_Child_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Child'
type MockRef_Child_Call struct {
	*mock.Call
}

// Child is a helper method to define mock.On call
//   - path string
func (_e *MockRef_Expecter) Child(path interface{}) *MockRef_Child_Call {
	return &MockRef_Child_Call{Call: _e.mock.On("Child", path)}
}

func (_c *MockRef_Child_Call) Run(run func(path string)) *MockRef_Child_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockRef_Child_Call) Return(_a0 modules.Ref) *MockRef_Child_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRef_Child_Call) RunAndReturn(run func(string) modules.Ref) *MockRef_Child_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx
func (_m *MockRef) Delete(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRef_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockRef_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockRef_Expecter) Delete(ctx interface{}) *MockRef_Delete_Call {
	return &MockRef_Delete_Call{Call: _e.mock.On("Delete", ctx)}
}

func (_c *MockRef_Delete_Call) Run(run func(ctx context.Context)) *MockRef_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockRef_Delete_Call) Return(_a0 error) *MockRef_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRef_Delete_Call) RunAndReturn(run func(context.Context) error) *MockRef_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, v
func (_m *MockRef) Get(ctx context.Context, v interface{}) error {
	ret := _m.Called(ctx, v)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) error); ok {
		r0 = rf(ctx, v)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRef_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockRef_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - v interface{}
func (_e *MockRef_Expecter) Get(ctx interface{}, v interface{}) *MockRef_Get_Call {
	return &MockRef_Get_Call{Call: _e.mock.On("Get", ctx, v)}
}

func (_c *MockRef_Get_Call) Run(run func(ctx context.Context, v interface{})) *MockRef_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(interface{}))
	})
	return _c
}

func (_c *MockRef_Get_Call) Return(_a0 error) *MockRef_Get_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRef_Get_Call) RunAndReturn(run func(context.Context, interface{}) error) *MockRef_Get_Call {
	_c.Call.Return(run)
	return _c
}

// Push provides a mock function with given fields: ctx, v
func (_m *MockRef) Push(ctx context.Context, v interface{}) (modules.Ref, error) {
	ret := _m.Called(ctx, v)

	if len(ret) == 0 {
		panic("no return value specified for Push")
	}

	var r0 modules.Ref
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) (modules.Ref, error)); ok {
		return rf(ctx, v)
	}
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) modules.Ref); ok {
		r0 = rf(ctx, v)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(modules.Ref)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, interface{}) error); ok {
		r1 = rf(ctx, v)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRef_Push_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Push'
type MockRef_Push_Call struct {
	*mock.Call
}

// Push is a helper method to define mock.On call
//   - ctx context.Context
//   - v interface{}
func (_e *MockRef_Expecter) Push(ctx interface{}, v interface{}) *MockRef_Push_Call {
	return &MockRef_Push_Call{Call: _e.mock.On("Push", ctx, v)}
}

func (_c *MockRef_Push_Call) Run(run func(ctx context.Context, v interface{})) *MockRef_Push_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(interface{}))
	})
	return _c
}

func (_c *MockRef_Push_Call) Return(_a0 modules.Ref, _a1 error) *MockRef_Push_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRef_Push_Call) RunAndReturn(run func(context.Context, interface{}) (modules.Ref, error)) *MockRef_Push_Call {
	_c.Call.Return(run)
	return _c
}

// Set provides a mock function with given fields: ctx, v
func (_m *MockRef) Set(ctx context.Context, v interface{}) error {
	ret := _m.Called(ctx, v)

	if len(ret) == 0 {
		panic("no return value specified for Set")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) error); ok {
		r0 = rf(ctx, v)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRef_Set_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Set'
type MockRef_Set_Call struct {
	*mock.Call
}

// Set is a helper method to define mock.On call
//   - ctx context.Context
//   - v interface{}
func (_e *MockRef_Expecter) Set(ctx interface{}, v interface{}) *MockRef_Set_Call {
	return &MockRef_Set_Call{Call: _e.mock.On("Set", ctx, v)}
}

func (_c *MockRef_Set_Call) Run(run func(ctx context.Context, v interface{})) *MockRef_Set_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(interface{}))
	})
	return _c
}

func (_c *MockRef_Set_Call) Return(_a0 error) *MockRef_Set_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRef_Set_Call) RunAndReturn(run func(context.Context, interface{}) error) *MockRef_Set_Call {
	_c.Call.Return(run)
	return _c
}

// Transaction provides a mock function with given fields: ctx, f
func (_m *MockRef) Transaction(ctx context.Context, f func(modules.TransactionNode) (interface{}, error)) error {
	ret := _m.Called(ctx, f)

	if len(ret) == 0 {
		panic("no return value specified for Transaction")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, func(modules.TransactionNode) (interface{}, error)) error); ok {
		r0 = rf(ctx, f)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRef_Transaction_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Transaction'
type MockRef_Transaction_Call struct {
	*mock.Call
}

// Transaction is a helper method to define mock.On call
//   - ctx context.Context
//   - f func(modules.TransactionNode)(interface{} , error)
func (_e *MockRef_Expecter) Transaction(ctx interface{}, f interface{}) *MockRef_Transaction_Call {
	return &MockRef_Transaction_Call{Call: _e.mock.On("Transaction", ctx, f)}
}

func (_c *MockRef_Transaction_Call) Run(run func(ctx context.Context, f func(modules.TransactionNode) (interface{}, error))) *MockRef_Transaction_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(func(modules.TransactionNode) (interface{}, error)))
	})
	return _c
}

func (_c *MockRef_Transaction_Call) Return(_a0 error) *MockRef_Transaction_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRef_Transaction_Call) RunAndReturn(run func(context.Context, func(modules.TransactionNode) (interface{}, error)) error) *MockRef_Transaction_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockRef creates a new instance of MockRef. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockRef(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockRef {
	mock := &MockRef{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
