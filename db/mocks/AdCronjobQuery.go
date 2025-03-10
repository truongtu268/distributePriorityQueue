// Code generated by mockery v2.42.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	query "github.com/truongtu268/distributePriorityQueue/db/query"
)

// AdCronjobQuery is an autogenerated mock type for the AdCronjobQuery type
type AdCronjobQuery struct {
	mock.Mock
}

// GetAdByID provides a mock function with given fields: ctx, id
func (_m *AdCronjobQuery) GetAdByID(ctx context.Context, id string) (query.Ad, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetAdByID")
	}

	var r0 query.Ad
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (query.Ad, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) query.Ad); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(query.Ad)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateAdAnalysis provides a mock function with given fields: ctx, arg
func (_m *AdCronjobQuery) UpdateAdAnalysis(ctx context.Context, arg query.UpdateAdAnalysisParams) error {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for UpdateAdAnalysis")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, query.UpdateAdAnalysisParams) error); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateAdRetry provides a mock function with given fields: ctx, arg
func (_m *AdCronjobQuery) UpdateAdRetry(ctx context.Context, arg query.UpdateAdRetryParams) error {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for UpdateAdRetry")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, query.UpdateAdRetryParams) error); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateAdStatus provides a mock function with given fields: ctx, arg
func (_m *AdCronjobQuery) UpdateAdStatus(ctx context.Context, arg query.UpdateAdStatusParams) error {
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

// NewAdCronjobQuery creates a new instance of AdCronjobQuery. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAdCronjobQuery(t interface {
	mock.TestingT
	Cleanup(func())
}) *AdCronjobQuery {
	mock := &AdCronjobQuery{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
