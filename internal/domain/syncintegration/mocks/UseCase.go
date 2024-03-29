// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	context "context"

	syncintegration "github.com/azisuazusa/todo-cli/internal/domain/syncintegration"
	mock "github.com/stretchr/testify/mock"
)

// UseCase is an autogenerated mock type for the UseCase type
type UseCase struct {
	mock.Mock
}

type UseCase_Expecter struct {
	mock *mock.Mock
}

func (_m *UseCase) EXPECT() *UseCase_Expecter {
	return &UseCase_Expecter{mock: &_m.Mock}
}

// Download provides a mock function with given fields: ctx
func (_m *UseCase) Download(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Download")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UseCase_Download_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Download'
type UseCase_Download_Call struct {
	*mock.Call
}

// Download is a helper method to define mock.On call
//   - ctx context.Context
func (_e *UseCase_Expecter) Download(ctx interface{}) *UseCase_Download_Call {
	return &UseCase_Download_Call{Call: _e.mock.On("Download", ctx)}
}

func (_c *UseCase_Download_Call) Run(run func(ctx context.Context)) *UseCase_Download_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *UseCase_Download_Call) Return(_a0 error) *UseCase_Download_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *UseCase_Download_Call) RunAndReturn(run func(context.Context) error) *UseCase_Download_Call {
	_c.Call.Return(run)
	return _c
}

// SetSyncIntegration provides a mock function with given fields: ctx, integration
func (_m *UseCase) SetSyncIntegration(ctx context.Context, integration syncintegration.SyncIntegration) error {
	ret := _m.Called(ctx, integration)

	if len(ret) == 0 {
		panic("no return value specified for SetSyncIntegration")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, syncintegration.SyncIntegration) error); ok {
		r0 = rf(ctx, integration)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UseCase_SetSyncIntegration_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetSyncIntegration'
type UseCase_SetSyncIntegration_Call struct {
	*mock.Call
}

// SetSyncIntegration is a helper method to define mock.On call
//   - ctx context.Context
//   - integration syncintegration.SyncIntegration
func (_e *UseCase_Expecter) SetSyncIntegration(ctx interface{}, integration interface{}) *UseCase_SetSyncIntegration_Call {
	return &UseCase_SetSyncIntegration_Call{Call: _e.mock.On("SetSyncIntegration", ctx, integration)}
}

func (_c *UseCase_SetSyncIntegration_Call) Run(run func(ctx context.Context, integration syncintegration.SyncIntegration)) *UseCase_SetSyncIntegration_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(syncintegration.SyncIntegration))
	})
	return _c
}

func (_c *UseCase_SetSyncIntegration_Call) Return(_a0 error) *UseCase_SetSyncIntegration_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *UseCase_SetSyncIntegration_Call) RunAndReturn(run func(context.Context, syncintegration.SyncIntegration) error) *UseCase_SetSyncIntegration_Call {
	_c.Call.Return(run)
	return _c
}

// Upload provides a mock function with given fields: ctx
func (_m *UseCase) Upload(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Upload")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UseCase_Upload_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Upload'
type UseCase_Upload_Call struct {
	*mock.Call
}

// Upload is a helper method to define mock.On call
//   - ctx context.Context
func (_e *UseCase_Expecter) Upload(ctx interface{}) *UseCase_Upload_Call {
	return &UseCase_Upload_Call{Call: _e.mock.On("Upload", ctx)}
}

func (_c *UseCase_Upload_Call) Run(run func(ctx context.Context)) *UseCase_Upload_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *UseCase_Upload_Call) Return(_a0 error) *UseCase_Upload_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *UseCase_Upload_Call) RunAndReturn(run func(context.Context) error) *UseCase_Upload_Call {
	_c.Call.Return(run)
	return _c
}

// NewUseCase creates a new instance of UseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUseCase(t interface {
	mock.TestingT
	Cleanup(func())
}) *UseCase {
	mock := &UseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
