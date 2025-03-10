// Code generated by mockery v2.42.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	query "github.com/truongtu268/distributePriorityQueue/db/query"
)

// AdQueueQuery is an autogenerated mock type for the AdQueueQuery type
type AdQueueQuery struct {
	mock.Mock
}

// UpdateAdStatus provides a mock function with given fields: ctx, arg
func (_m *AdQueueQuery) UpdateAdStatus(ctx context.Context, arg query.UpdateAdStatusParams) error {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for UpdateAdStatus")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, query.UpdateAdStatusParams) error); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewAdQueueQuery creates a new instance of AdQueueQuery. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAdQueueQuery(t interface {
	mock.TestingT
	Cleanup(func())
}) *AdQueueQuery {
	mock := &AdQueueQuery{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
