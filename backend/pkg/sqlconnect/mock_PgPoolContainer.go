// Code generated by mockery. DO NOT EDIT.

package sqlconnect

import (
	context "context"

	pg_queries "github.com/nucleuscloud/neosync/backend/gen/go/db/dbschemas/postgresql"
	mock "github.com/stretchr/testify/mock"
)

// MockPgPoolContainer is an autogenerated mock type for the PgPoolContainer type
type MockPgPoolContainer struct {
	mock.Mock
}

type MockPgPoolContainer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPgPoolContainer) EXPECT() *MockPgPoolContainer_Expecter {
	return &MockPgPoolContainer_Expecter{mock: &_m.Mock}
}

// Close provides a mock function with given fields:
func (_m *MockPgPoolContainer) Close() {
	_m.Called()
}

// MockPgPoolContainer_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type MockPgPoolContainer_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *MockPgPoolContainer_Expecter) Close() *MockPgPoolContainer_Close_Call {
	return &MockPgPoolContainer_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *MockPgPoolContainer_Close_Call) Run(run func()) *MockPgPoolContainer_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPgPoolContainer_Close_Call) Return() *MockPgPoolContainer_Close_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockPgPoolContainer_Close_Call) RunAndReturn(run func()) *MockPgPoolContainer_Close_Call {
	_c.Call.Return(run)
	return _c
}

// Open provides a mock function with given fields: _a0
func (_m *MockPgPoolContainer) Open(_a0 context.Context) (pg_queries.DBTX, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Open")
	}

	var r0 pg_queries.DBTX
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (pg_queries.DBTX, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(context.Context) pg_queries.DBTX); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(pg_queries.DBTX)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPgPoolContainer_Open_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Open'
type MockPgPoolContainer_Open_Call struct {
	*mock.Call
}

// Open is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *MockPgPoolContainer_Expecter) Open(_a0 interface{}) *MockPgPoolContainer_Open_Call {
	return &MockPgPoolContainer_Open_Call{Call: _e.mock.On("Open", _a0)}
}

func (_c *MockPgPoolContainer_Open_Call) Run(run func(_a0 context.Context)) *MockPgPoolContainer_Open_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockPgPoolContainer_Open_Call) Return(_a0 pg_queries.DBTX, _a1 error) *MockPgPoolContainer_Open_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPgPoolContainer_Open_Call) RunAndReturn(run func(context.Context) (pg_queries.DBTX, error)) *MockPgPoolContainer_Open_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockPgPoolContainer creates a new instance of MockPgPoolContainer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPgPoolContainer(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPgPoolContainer {
	mock := &MockPgPoolContainer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
