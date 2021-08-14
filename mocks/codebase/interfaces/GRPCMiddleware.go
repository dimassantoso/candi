// Code generated by mockery v2.8.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	types "pkg.agungdp.dev/candi/codebase/factory/types"
)

// GRPCMiddleware is an autogenerated mock type for the GRPCMiddleware type
type GRPCMiddleware struct {
	mock.Mock
}

// GRPCBasicAuth provides a mock function with given fields: ctx
func (_m *GRPCMiddleware) GRPCBasicAuth(ctx context.Context) context.Context {
	ret := _m.Called(ctx)

	var r0 context.Context
	if rf, ok := ret.Get(0).(func(context.Context) context.Context); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	return r0
}

// GRPCBearerAuth provides a mock function with given fields: ctx
func (_m *GRPCMiddleware) GRPCBearerAuth(ctx context.Context) context.Context {
	ret := _m.Called(ctx)

	var r0 context.Context
	if rf, ok := ret.Get(0).(func(context.Context) context.Context); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	return r0
}

// GRPCPermissionACL provides a mock function with given fields: permissionCode
func (_m *GRPCMiddleware) GRPCPermissionACL(permissionCode string) types.MiddlewareFunc {
	ret := _m.Called(permissionCode)

	var r0 types.MiddlewareFunc
	if rf, ok := ret.Get(0).(func(string) types.MiddlewareFunc); ok {
		r0 = rf(permissionCode)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.MiddlewareFunc)
		}
	}

	return r0
}
