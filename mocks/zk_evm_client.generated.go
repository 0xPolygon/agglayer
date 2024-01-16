// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	context "context"
	big "math/big"

	mock "github.com/stretchr/testify/mock"

	types "github.com/0xPolygonHermez/zkevm-node/jsonrpc/types"
)

// ZkEVMClientMock is an autogenerated mock type for the IZkEVMClient type
type ZkEVMClientMock struct {
	mock.Mock
}

// BatchByNumber provides a mock function with given fields: ctx, number
func (_m *ZkEVMClientMock) BatchByNumber(ctx context.Context, number *big.Int) (*types.Batch, error) {
	ret := _m.Called(ctx, number)

	if len(ret) == 0 {
		panic("no return value specified for BatchByNumber")
	}

	var r0 *types.Batch
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *big.Int) (*types.Batch, error)); ok {
		return rf(ctx, number)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *big.Int) *types.Batch); ok {
		r0 = rf(ctx, number)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Batch)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *big.Int) error); ok {
		r1 = rf(ctx, number)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewZkEVMClientMock creates a new instance of ZkEVMClientMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewZkEVMClientMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *ZkEVMClientMock {
	mock := &ZkEVMClientMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}