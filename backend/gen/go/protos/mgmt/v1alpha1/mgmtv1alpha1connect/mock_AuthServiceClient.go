// Code generated by mockery. DO NOT EDIT.

package mgmtv1alpha1connect

import (
	context "context"

	connect "connectrpc.com/connect"

	mgmtv1alpha1 "github.com/nucleuscloud/neosync/backend/gen/go/protos/mgmt/v1alpha1"
	mock "github.com/stretchr/testify/mock"
)

// MockAuthServiceClient is an autogenerated mock type for the AuthServiceClient type
type MockAuthServiceClient struct {
	mock.Mock
}

type MockAuthServiceClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockAuthServiceClient) EXPECT() *MockAuthServiceClient_Expecter {
	return &MockAuthServiceClient_Expecter{mock: &_m.Mock}
}

// CheckToken provides a mock function with given fields: _a0, _a1
func (_m *MockAuthServiceClient) CheckToken(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.CheckTokenRequest]) (*connect.Response[mgmtv1alpha1.CheckTokenResponse], error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for CheckToken")
	}

	var r0 *connect.Response[mgmtv1alpha1.CheckTokenResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.CheckTokenRequest]) (*connect.Response[mgmtv1alpha1.CheckTokenResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.CheckTokenRequest]) *connect.Response[mgmtv1alpha1.CheckTokenResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[mgmtv1alpha1.CheckTokenResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[mgmtv1alpha1.CheckTokenRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAuthServiceClient_CheckToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CheckToken'
type MockAuthServiceClient_CheckToken_Call struct {
	*mock.Call
}

// CheckToken is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *connect.Request[mgmtv1alpha1.CheckTokenRequest]
func (_e *MockAuthServiceClient_Expecter) CheckToken(_a0 interface{}, _a1 interface{}) *MockAuthServiceClient_CheckToken_Call {
	return &MockAuthServiceClient_CheckToken_Call{Call: _e.mock.On("CheckToken", _a0, _a1)}
}

func (_c *MockAuthServiceClient_CheckToken_Call) Run(run func(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.CheckTokenRequest])) *MockAuthServiceClient_CheckToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*connect.Request[mgmtv1alpha1.CheckTokenRequest]))
	})
	return _c
}

func (_c *MockAuthServiceClient_CheckToken_Call) Return(_a0 *connect.Response[mgmtv1alpha1.CheckTokenResponse], _a1 error) *MockAuthServiceClient_CheckToken_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAuthServiceClient_CheckToken_Call) RunAndReturn(run func(context.Context, *connect.Request[mgmtv1alpha1.CheckTokenRequest]) (*connect.Response[mgmtv1alpha1.CheckTokenResponse], error)) *MockAuthServiceClient_CheckToken_Call {
	_c.Call.Return(run)
	return _c
}

// GetAuthStatus provides a mock function with given fields: _a0, _a1
func (_m *MockAuthServiceClient) GetAuthStatus(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.GetAuthStatusRequest]) (*connect.Response[mgmtv1alpha1.GetAuthStatusResponse], error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetAuthStatus")
	}

	var r0 *connect.Response[mgmtv1alpha1.GetAuthStatusResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.GetAuthStatusRequest]) (*connect.Response[mgmtv1alpha1.GetAuthStatusResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.GetAuthStatusRequest]) *connect.Response[mgmtv1alpha1.GetAuthStatusResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[mgmtv1alpha1.GetAuthStatusResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[mgmtv1alpha1.GetAuthStatusRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAuthServiceClient_GetAuthStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAuthStatus'
type MockAuthServiceClient_GetAuthStatus_Call struct {
	*mock.Call
}

// GetAuthStatus is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *connect.Request[mgmtv1alpha1.GetAuthStatusRequest]
func (_e *MockAuthServiceClient_Expecter) GetAuthStatus(_a0 interface{}, _a1 interface{}) *MockAuthServiceClient_GetAuthStatus_Call {
	return &MockAuthServiceClient_GetAuthStatus_Call{Call: _e.mock.On("GetAuthStatus", _a0, _a1)}
}

func (_c *MockAuthServiceClient_GetAuthStatus_Call) Run(run func(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.GetAuthStatusRequest])) *MockAuthServiceClient_GetAuthStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*connect.Request[mgmtv1alpha1.GetAuthStatusRequest]))
	})
	return _c
}

func (_c *MockAuthServiceClient_GetAuthStatus_Call) Return(_a0 *connect.Response[mgmtv1alpha1.GetAuthStatusResponse], _a1 error) *MockAuthServiceClient_GetAuthStatus_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAuthServiceClient_GetAuthStatus_Call) RunAndReturn(run func(context.Context, *connect.Request[mgmtv1alpha1.GetAuthStatusRequest]) (*connect.Response[mgmtv1alpha1.GetAuthStatusResponse], error)) *MockAuthServiceClient_GetAuthStatus_Call {
	_c.Call.Return(run)
	return _c
}

// GetAuthorizeUrl provides a mock function with given fields: _a0, _a1
func (_m *MockAuthServiceClient) GetAuthorizeUrl(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.GetAuthorizeUrlRequest]) (*connect.Response[mgmtv1alpha1.GetAuthorizeUrlResponse], error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetAuthorizeUrl")
	}

	var r0 *connect.Response[mgmtv1alpha1.GetAuthorizeUrlResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.GetAuthorizeUrlRequest]) (*connect.Response[mgmtv1alpha1.GetAuthorizeUrlResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.GetAuthorizeUrlRequest]) *connect.Response[mgmtv1alpha1.GetAuthorizeUrlResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[mgmtv1alpha1.GetAuthorizeUrlResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[mgmtv1alpha1.GetAuthorizeUrlRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAuthServiceClient_GetAuthorizeUrl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAuthorizeUrl'
type MockAuthServiceClient_GetAuthorizeUrl_Call struct {
	*mock.Call
}

// GetAuthorizeUrl is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *connect.Request[mgmtv1alpha1.GetAuthorizeUrlRequest]
func (_e *MockAuthServiceClient_Expecter) GetAuthorizeUrl(_a0 interface{}, _a1 interface{}) *MockAuthServiceClient_GetAuthorizeUrl_Call {
	return &MockAuthServiceClient_GetAuthorizeUrl_Call{Call: _e.mock.On("GetAuthorizeUrl", _a0, _a1)}
}

func (_c *MockAuthServiceClient_GetAuthorizeUrl_Call) Run(run func(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.GetAuthorizeUrlRequest])) *MockAuthServiceClient_GetAuthorizeUrl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*connect.Request[mgmtv1alpha1.GetAuthorizeUrlRequest]))
	})
	return _c
}

func (_c *MockAuthServiceClient_GetAuthorizeUrl_Call) Return(_a0 *connect.Response[mgmtv1alpha1.GetAuthorizeUrlResponse], _a1 error) *MockAuthServiceClient_GetAuthorizeUrl_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAuthServiceClient_GetAuthorizeUrl_Call) RunAndReturn(run func(context.Context, *connect.Request[mgmtv1alpha1.GetAuthorizeUrlRequest]) (*connect.Response[mgmtv1alpha1.GetAuthorizeUrlResponse], error)) *MockAuthServiceClient_GetAuthorizeUrl_Call {
	_c.Call.Return(run)
	return _c
}

// GetCliIssuer provides a mock function with given fields: _a0, _a1
func (_m *MockAuthServiceClient) GetCliIssuer(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.GetCliIssuerRequest]) (*connect.Response[mgmtv1alpha1.GetCliIssuerResponse], error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetCliIssuer")
	}

	var r0 *connect.Response[mgmtv1alpha1.GetCliIssuerResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.GetCliIssuerRequest]) (*connect.Response[mgmtv1alpha1.GetCliIssuerResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.GetCliIssuerRequest]) *connect.Response[mgmtv1alpha1.GetCliIssuerResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[mgmtv1alpha1.GetCliIssuerResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[mgmtv1alpha1.GetCliIssuerRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAuthServiceClient_GetCliIssuer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCliIssuer'
type MockAuthServiceClient_GetCliIssuer_Call struct {
	*mock.Call
}

// GetCliIssuer is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *connect.Request[mgmtv1alpha1.GetCliIssuerRequest]
func (_e *MockAuthServiceClient_Expecter) GetCliIssuer(_a0 interface{}, _a1 interface{}) *MockAuthServiceClient_GetCliIssuer_Call {
	return &MockAuthServiceClient_GetCliIssuer_Call{Call: _e.mock.On("GetCliIssuer", _a0, _a1)}
}

func (_c *MockAuthServiceClient_GetCliIssuer_Call) Run(run func(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.GetCliIssuerRequest])) *MockAuthServiceClient_GetCliIssuer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*connect.Request[mgmtv1alpha1.GetCliIssuerRequest]))
	})
	return _c
}

func (_c *MockAuthServiceClient_GetCliIssuer_Call) Return(_a0 *connect.Response[mgmtv1alpha1.GetCliIssuerResponse], _a1 error) *MockAuthServiceClient_GetCliIssuer_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAuthServiceClient_GetCliIssuer_Call) RunAndReturn(run func(context.Context, *connect.Request[mgmtv1alpha1.GetCliIssuerRequest]) (*connect.Response[mgmtv1alpha1.GetCliIssuerResponse], error)) *MockAuthServiceClient_GetCliIssuer_Call {
	_c.Call.Return(run)
	return _c
}

// LoginCli provides a mock function with given fields: _a0, _a1
func (_m *MockAuthServiceClient) LoginCli(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.LoginCliRequest]) (*connect.Response[mgmtv1alpha1.LoginCliResponse], error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for LoginCli")
	}

	var r0 *connect.Response[mgmtv1alpha1.LoginCliResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.LoginCliRequest]) (*connect.Response[mgmtv1alpha1.LoginCliResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.LoginCliRequest]) *connect.Response[mgmtv1alpha1.LoginCliResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[mgmtv1alpha1.LoginCliResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[mgmtv1alpha1.LoginCliRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAuthServiceClient_LoginCli_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LoginCli'
type MockAuthServiceClient_LoginCli_Call struct {
	*mock.Call
}

// LoginCli is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *connect.Request[mgmtv1alpha1.LoginCliRequest]
func (_e *MockAuthServiceClient_Expecter) LoginCli(_a0 interface{}, _a1 interface{}) *MockAuthServiceClient_LoginCli_Call {
	return &MockAuthServiceClient_LoginCli_Call{Call: _e.mock.On("LoginCli", _a0, _a1)}
}

func (_c *MockAuthServiceClient_LoginCli_Call) Run(run func(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.LoginCliRequest])) *MockAuthServiceClient_LoginCli_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*connect.Request[mgmtv1alpha1.LoginCliRequest]))
	})
	return _c
}

func (_c *MockAuthServiceClient_LoginCli_Call) Return(_a0 *connect.Response[mgmtv1alpha1.LoginCliResponse], _a1 error) *MockAuthServiceClient_LoginCli_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAuthServiceClient_LoginCli_Call) RunAndReturn(run func(context.Context, *connect.Request[mgmtv1alpha1.LoginCliRequest]) (*connect.Response[mgmtv1alpha1.LoginCliResponse], error)) *MockAuthServiceClient_LoginCli_Call {
	_c.Call.Return(run)
	return _c
}

// RefreshCli provides a mock function with given fields: _a0, _a1
func (_m *MockAuthServiceClient) RefreshCli(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.RefreshCliRequest]) (*connect.Response[mgmtv1alpha1.RefreshCliResponse], error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for RefreshCli")
	}

	var r0 *connect.Response[mgmtv1alpha1.RefreshCliResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.RefreshCliRequest]) (*connect.Response[mgmtv1alpha1.RefreshCliResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.RefreshCliRequest]) *connect.Response[mgmtv1alpha1.RefreshCliResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[mgmtv1alpha1.RefreshCliResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[mgmtv1alpha1.RefreshCliRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAuthServiceClient_RefreshCli_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RefreshCli'
type MockAuthServiceClient_RefreshCli_Call struct {
	*mock.Call
}

// RefreshCli is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *connect.Request[mgmtv1alpha1.RefreshCliRequest]
func (_e *MockAuthServiceClient_Expecter) RefreshCli(_a0 interface{}, _a1 interface{}) *MockAuthServiceClient_RefreshCli_Call {
	return &MockAuthServiceClient_RefreshCli_Call{Call: _e.mock.On("RefreshCli", _a0, _a1)}
}

func (_c *MockAuthServiceClient_RefreshCli_Call) Run(run func(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.RefreshCliRequest])) *MockAuthServiceClient_RefreshCli_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*connect.Request[mgmtv1alpha1.RefreshCliRequest]))
	})
	return _c
}

func (_c *MockAuthServiceClient_RefreshCli_Call) Return(_a0 *connect.Response[mgmtv1alpha1.RefreshCliResponse], _a1 error) *MockAuthServiceClient_RefreshCli_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAuthServiceClient_RefreshCli_Call) RunAndReturn(run func(context.Context, *connect.Request[mgmtv1alpha1.RefreshCliRequest]) (*connect.Response[mgmtv1alpha1.RefreshCliResponse], error)) *MockAuthServiceClient_RefreshCli_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockAuthServiceClient creates a new instance of MockAuthServiceClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockAuthServiceClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockAuthServiceClient {
	mock := &MockAuthServiceClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
