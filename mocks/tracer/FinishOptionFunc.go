// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	tracer "github.com/golangid/candi/tracer"
	mock "github.com/stretchr/testify/mock"
)

// FinishOptionFunc is an autogenerated mock type for the FinishOptionFunc type
type FinishOptionFunc struct {
	mock.Mock
}

// Execute provides a mock function with given fields: _a0
func (_m *FinishOptionFunc) Execute(_a0 *tracer.FinishOption) {
	_m.Called(_a0)
}

type mockConstructorTestingTNewFinishOptionFunc interface {
	mock.TestingT
	Cleanup(func())
}

// NewFinishOptionFunc creates a new instance of FinishOptionFunc. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFinishOptionFunc(t mockConstructorTestingTNewFinishOptionFunc) *FinishOptionFunc {
	mock := &FinishOptionFunc{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
