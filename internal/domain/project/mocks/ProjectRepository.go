// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/azisuazusa/todo-cli/internal/domain/entity"
	mock "github.com/stretchr/testify/mock"
)

// ProjectRepository is an autogenerated mock type for the ProjectRepository type
type ProjectRepository struct {
	mock.Mock
}

type ProjectRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *ProjectRepository) EXPECT() *ProjectRepository_Expecter {
	return &ProjectRepository_Expecter{mock: &_m.Mock}
}

// Delete provides a mock function with given fields: ctx, id
func (_m *ProjectRepository) Delete(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ProjectRepository_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type ProjectRepository_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *ProjectRepository_Expecter) Delete(ctx interface{}, id interface{}) *ProjectRepository_Delete_Call {
	return &ProjectRepository_Delete_Call{Call: _e.mock.On("Delete", ctx, id)}
}

func (_c *ProjectRepository_Delete_Call) Run(run func(ctx context.Context, id string)) *ProjectRepository_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *ProjectRepository_Delete_Call) Return(_a0 error) *ProjectRepository_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ProjectRepository_Delete_Call) RunAndReturn(run func(context.Context, string) error) *ProjectRepository_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// GetAll provides a mock function with given fields: ctx
func (_m *ProjectRepository) GetAll(ctx context.Context) (entity.Projects, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
	}

	var r0 entity.Projects
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (entity.Projects, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) entity.Projects); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(entity.Projects)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProjectRepository_GetAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAll'
type ProjectRepository_GetAll_Call struct {
	*mock.Call
}

// GetAll is a helper method to define mock.On call
//   - ctx context.Context
func (_e *ProjectRepository_Expecter) GetAll(ctx interface{}) *ProjectRepository_GetAll_Call {
	return &ProjectRepository_GetAll_Call{Call: _e.mock.On("GetAll", ctx)}
}

func (_c *ProjectRepository_GetAll_Call) Run(run func(ctx context.Context)) *ProjectRepository_GetAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *ProjectRepository_GetAll_Call) Return(_a0 entity.Projects, _a1 error) *ProjectRepository_GetAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ProjectRepository_GetAll_Call) RunAndReturn(run func(context.Context) (entity.Projects, error)) *ProjectRepository_GetAll_Call {
	_c.Call.Return(run)
	return _c
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *ProjectRepository) GetByID(ctx context.Context, id string) (entity.Project, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetByID")
	}

	var r0 entity.Project
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (entity.Project, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) entity.Project); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(entity.Project)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProjectRepository_GetByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByID'
type ProjectRepository_GetByID_Call struct {
	*mock.Call
}

// GetByID is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *ProjectRepository_Expecter) GetByID(ctx interface{}, id interface{}) *ProjectRepository_GetByID_Call {
	return &ProjectRepository_GetByID_Call{Call: _e.mock.On("GetByID", ctx, id)}
}

func (_c *ProjectRepository_GetByID_Call) Run(run func(ctx context.Context, id string)) *ProjectRepository_GetByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *ProjectRepository_GetByID_Call) Return(_a0 entity.Project, _a1 error) *ProjectRepository_GetByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ProjectRepository_GetByID_Call) RunAndReturn(run func(context.Context, string) (entity.Project, error)) *ProjectRepository_GetByID_Call {
	_c.Call.Return(run)
	return _c
}

// GetSelectedProject provides a mock function with given fields: ctx
func (_m *ProjectRepository) GetSelectedProject(ctx context.Context) (entity.Project, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetSelectedProject")
	}

	var r0 entity.Project
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (entity.Project, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) entity.Project); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(entity.Project)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProjectRepository_GetSelectedProject_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSelectedProject'
type ProjectRepository_GetSelectedProject_Call struct {
	*mock.Call
}

// GetSelectedProject is a helper method to define mock.On call
//   - ctx context.Context
func (_e *ProjectRepository_Expecter) GetSelectedProject(ctx interface{}) *ProjectRepository_GetSelectedProject_Call {
	return &ProjectRepository_GetSelectedProject_Call{Call: _e.mock.On("GetSelectedProject", ctx)}
}

func (_c *ProjectRepository_GetSelectedProject_Call) Run(run func(ctx context.Context)) *ProjectRepository_GetSelectedProject_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *ProjectRepository_GetSelectedProject_Call) Return(_a0 entity.Project, _a1 error) *ProjectRepository_GetSelectedProject_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ProjectRepository_GetSelectedProject_Call) RunAndReturn(run func(context.Context) (entity.Project, error)) *ProjectRepository_GetSelectedProject_Call {
	_c.Call.Return(run)
	return _c
}

// Insert provides a mock function with given fields: ctx, _a1
func (_m *ProjectRepository) Insert(ctx context.Context, _a1 entity.Project) error {
	ret := _m.Called(ctx, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Insert")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.Project) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ProjectRepository_Insert_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Insert'
type ProjectRepository_Insert_Call struct {
	*mock.Call
}

// Insert is a helper method to define mock.On call
//   - ctx context.Context
//   - _a1 entity.Project
func (_e *ProjectRepository_Expecter) Insert(ctx interface{}, _a1 interface{}) *ProjectRepository_Insert_Call {
	return &ProjectRepository_Insert_Call{Call: _e.mock.On("Insert", ctx, _a1)}
}

func (_c *ProjectRepository_Insert_Call) Run(run func(ctx context.Context, _a1 entity.Project)) *ProjectRepository_Insert_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(entity.Project))
	})
	return _c
}

func (_c *ProjectRepository_Insert_Call) Return(_a0 error) *ProjectRepository_Insert_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ProjectRepository_Insert_Call) RunAndReturn(run func(context.Context, entity.Project) error) *ProjectRepository_Insert_Call {
	_c.Call.Return(run)
	return _c
}

// SetSelectedProject provides a mock function with given fields: ctx, id
func (_m *ProjectRepository) SetSelectedProject(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for SetSelectedProject")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ProjectRepository_SetSelectedProject_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetSelectedProject'
type ProjectRepository_SetSelectedProject_Call struct {
	*mock.Call
}

// SetSelectedProject is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *ProjectRepository_Expecter) SetSelectedProject(ctx interface{}, id interface{}) *ProjectRepository_SetSelectedProject_Call {
	return &ProjectRepository_SetSelectedProject_Call{Call: _e.mock.On("SetSelectedProject", ctx, id)}
}

func (_c *ProjectRepository_SetSelectedProject_Call) Run(run func(ctx context.Context, id string)) *ProjectRepository_SetSelectedProject_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *ProjectRepository_SetSelectedProject_Call) Return(_a0 error) *ProjectRepository_SetSelectedProject_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ProjectRepository_SetSelectedProject_Call) RunAndReturn(run func(context.Context, string) error) *ProjectRepository_SetSelectedProject_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, _a1
func (_m *ProjectRepository) Update(ctx context.Context, _a1 entity.Project) error {
	ret := _m.Called(ctx, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.Project) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ProjectRepository_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type ProjectRepository_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - _a1 entity.Project
func (_e *ProjectRepository_Expecter) Update(ctx interface{}, _a1 interface{}) *ProjectRepository_Update_Call {
	return &ProjectRepository_Update_Call{Call: _e.mock.On("Update", ctx, _a1)}
}

func (_c *ProjectRepository_Update_Call) Run(run func(ctx context.Context, _a1 entity.Project)) *ProjectRepository_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(entity.Project))
	})
	return _c
}

func (_c *ProjectRepository_Update_Call) Return(_a0 error) *ProjectRepository_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ProjectRepository_Update_Call) RunAndReturn(run func(context.Context, entity.Project) error) *ProjectRepository_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewProjectRepository creates a new instance of ProjectRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewProjectRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *ProjectRepository {
	mock := &ProjectRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
