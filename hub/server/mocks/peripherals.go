// Code generated by mockery v2.33.0. DO NOT EDIT.

package mocks

import (
	context "context"

	peripherals "github.com/qwark97/go-versatility-presentation/hub/peripherals"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// Peripherals is an autogenerated mock type for the Peripherals type
type Peripherals struct {
	mock.Mock
}

type Peripherals_Expecter struct {
	mock *mock.Mock
}

func (_m *Peripherals) EXPECT() *Peripherals_Expecter {
	return &Peripherals_Expecter{mock: &_m.Mock}
}

// Add provides a mock function with given fields: ctx, configuration
func (_m *Peripherals) Add(ctx context.Context, configuration peripherals.Configuration) error {
	ret := _m.Called(ctx, configuration)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, peripherals.Configuration) error); ok {
		r0 = rf(ctx, configuration)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Peripherals_Add_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Add'
type Peripherals_Add_Call struct {
	*mock.Call
}

// Add is a helper method to define mock.On call
//   - ctx context.Context
//   - configuration peripherals.Configuration
func (_e *Peripherals_Expecter) Add(ctx interface{}, configuration interface{}) *Peripherals_Add_Call {
	return &Peripherals_Add_Call{Call: _e.mock.On("Add", ctx, configuration)}
}

func (_c *Peripherals_Add_Call) Run(run func(ctx context.Context, configuration peripherals.Configuration)) *Peripherals_Add_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(peripherals.Configuration))
	})
	return _c
}

func (_c *Peripherals_Add_Call) Return(_a0 error) *Peripherals_Add_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Peripherals_Add_Call) RunAndReturn(run func(context.Context, peripherals.Configuration) error) *Peripherals_Add_Call {
	_c.Call.Return(run)
	return _c
}

// All provides a mock function with given fields: ctx
func (_m *Peripherals) All(ctx context.Context) ([]peripherals.Configuration, error) {
	ret := _m.Called(ctx)

	var r0 []peripherals.Configuration
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]peripherals.Configuration, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []peripherals.Configuration); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]peripherals.Configuration)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Peripherals_All_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'All'
type Peripherals_All_Call struct {
	*mock.Call
}

// All is a helper method to define mock.On call
//   - ctx context.Context
func (_e *Peripherals_Expecter) All(ctx interface{}) *Peripherals_All_Call {
	return &Peripherals_All_Call{Call: _e.mock.On("All", ctx)}
}

func (_c *Peripherals_All_Call) Run(run func(ctx context.Context)) *Peripherals_All_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Peripherals_All_Call) Return(_a0 []peripherals.Configuration, _a1 error) *Peripherals_All_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Peripherals_All_Call) RunAndReturn(run func(context.Context) ([]peripherals.Configuration, error)) *Peripherals_All_Call {
	_c.Call.Return(run)
	return _c
}

// ByID provides a mock function with given fields: ctx, id
func (_m *Peripherals) ByID(ctx context.Context, id uuid.UUID) (peripherals.Configuration, error) {
	ret := _m.Called(ctx, id)

	var r0 peripherals.Configuration
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (peripherals.Configuration, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) peripherals.Configuration); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(peripherals.Configuration)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Peripherals_ByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ByID'
type Peripherals_ByID_Call struct {
	*mock.Call
}

// ByID is a helper method to define mock.On call
//   - ctx context.Context
//   - id uuid.UUID
func (_e *Peripherals_Expecter) ByID(ctx interface{}, id interface{}) *Peripherals_ByID_Call {
	return &Peripherals_ByID_Call{Call: _e.mock.On("ByID", ctx, id)}
}

func (_c *Peripherals_ByID_Call) Run(run func(ctx context.Context, id uuid.UUID)) *Peripherals_ByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *Peripherals_ByID_Call) Return(_a0 peripherals.Configuration, _a1 error) *Peripherals_ByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Peripherals_ByID_Call) RunAndReturn(run func(context.Context, uuid.UUID) (peripherals.Configuration, error)) *Peripherals_ByID_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteOne provides a mock function with given fields: ctx, id
func (_m *Peripherals) DeleteOne(ctx context.Context, id uuid.UUID) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Peripherals_DeleteOne_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteOne'
type Peripherals_DeleteOne_Call struct {
	*mock.Call
}

// DeleteOne is a helper method to define mock.On call
//   - ctx context.Context
//   - id uuid.UUID
func (_e *Peripherals_Expecter) DeleteOne(ctx interface{}, id interface{}) *Peripherals_DeleteOne_Call {
	return &Peripherals_DeleteOne_Call{Call: _e.mock.On("DeleteOne", ctx, id)}
}

func (_c *Peripherals_DeleteOne_Call) Run(run func(ctx context.Context, id uuid.UUID)) *Peripherals_DeleteOne_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *Peripherals_DeleteOne_Call) Return(_a0 error) *Peripherals_DeleteOne_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Peripherals_DeleteOne_Call) RunAndReturn(run func(context.Context, uuid.UUID) error) *Peripherals_DeleteOne_Call {
	_c.Call.Return(run)
	return _c
}

// Reload provides a mock function with given fields: ctx
func (_m *Peripherals) Reload(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Peripherals_Reload_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Reload'
type Peripherals_Reload_Call struct {
	*mock.Call
}

// Reload is a helper method to define mock.On call
//   - ctx context.Context
func (_e *Peripherals_Expecter) Reload(ctx interface{}) *Peripherals_Reload_Call {
	return &Peripherals_Reload_Call{Call: _e.mock.On("Reload", ctx)}
}

func (_c *Peripherals_Reload_Call) Run(run func(ctx context.Context)) *Peripherals_Reload_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Peripherals_Reload_Call) Return(_a0 error) *Peripherals_Reload_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Peripherals_Reload_Call) RunAndReturn(run func(context.Context) error) *Peripherals_Reload_Call {
	_c.Call.Return(run)
	return _c
}

// Verify provides a mock function with given fields: ctx, id
func (_m *Peripherals) Verify(ctx context.Context, id uuid.UUID) (bool, error) {
	ret := _m.Called(ctx, id)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (bool, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) bool); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Peripherals_Verify_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Verify'
type Peripherals_Verify_Call struct {
	*mock.Call
}

// Verify is a helper method to define mock.On call
//   - ctx context.Context
//   - id uuid.UUID
func (_e *Peripherals_Expecter) Verify(ctx interface{}, id interface{}) *Peripherals_Verify_Call {
	return &Peripherals_Verify_Call{Call: _e.mock.On("Verify", ctx, id)}
}

func (_c *Peripherals_Verify_Call) Run(run func(ctx context.Context, id uuid.UUID)) *Peripherals_Verify_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *Peripherals_Verify_Call) Return(_a0 bool, _a1 error) *Peripherals_Verify_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Peripherals_Verify_Call) RunAndReturn(run func(context.Context, uuid.UUID) (bool, error)) *Peripherals_Verify_Call {
	_c.Call.Return(run)
	return _c
}

// NewPeripherals creates a new instance of Peripherals. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPeripherals(t interface {
	mock.TestingT
	Cleanup(func())
}) *Peripherals {
	mock := &Peripherals{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}