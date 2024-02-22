// Code generated by mockery v2.39.0. DO NOT EDIT.

package mocks

import (
	context "context"
	big "math/big"

	common "github.com/ethereum/go-ethereum/common"

	ethtxmanager "github.com/0xPolygonHermez/zkevm-node/ethtxmanager"

	mock "github.com/stretchr/testify/mock"

	pgx "github.com/jackc/pgx/v4"
)

// EthTxManagerMock is an autogenerated mock type for the IEthTxManager type
type EthTxManagerMock struct {
	mock.Mock
}

type EthTxManagerMock_Expecter struct {
	mock *mock.Mock
}

func (_m *EthTxManagerMock) EXPECT() *EthTxManagerMock_Expecter {
	return &EthTxManagerMock_Expecter{mock: &_m.Mock}
}

// Add provides a mock function with given fields: ctx, owner, id, from, to, value, data, gasOffset, dbTx
func (_m *EthTxManagerMock) Add(ctx context.Context, owner string, id string, from common.Address, to *common.Address, value *big.Int, data []byte, gasOffset uint64, dbTx pgx.Tx) error {
	ret := _m.Called(ctx, owner, id, from, to, value, data, gasOffset, dbTx)

	if len(ret) == 0 {
		panic("no return value specified for Add")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, common.Address, *common.Address, *big.Int, []byte, uint64, pgx.Tx) error); ok {
		r0 = rf(ctx, owner, id, from, to, value, data, gasOffset, dbTx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EthTxManagerMock_Add_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Add'
type EthTxManagerMock_Add_Call struct {
	*mock.Call
}

// Add is a helper method to define mock.On call
//   - ctx context.Context
//   - owner string
//   - id string
//   - from common.Address
//   - to *common.Address
//   - value *big.Int
//   - data []byte
//   - gasOffset uint64
//   - dbTx pgx.Tx
func (_e *EthTxManagerMock_Expecter) Add(ctx interface{}, owner interface{}, id interface{}, from interface{}, to interface{}, value interface{}, data interface{}, gasOffset interface{}, dbTx interface{}) *EthTxManagerMock_Add_Call {
	return &EthTxManagerMock_Add_Call{Call: _e.mock.On("Add", ctx, owner, id, from, to, value, data, gasOffset, dbTx)}
}

func (_c *EthTxManagerMock_Add_Call) Run(run func(ctx context.Context, owner string, id string, from common.Address, to *common.Address, value *big.Int, data []byte, gasOffset uint64, dbTx pgx.Tx)) *EthTxManagerMock_Add_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(common.Address), args[4].(*common.Address), args[5].(*big.Int), args[6].([]byte), args[7].(uint64), args[8].(pgx.Tx))
	})
	return _c
}

func (_c *EthTxManagerMock_Add_Call) Return(_a0 error) *EthTxManagerMock_Add_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *EthTxManagerMock_Add_Call) RunAndReturn(run func(context.Context, string, string, common.Address, *common.Address, *big.Int, []byte, uint64, pgx.Tx) error) *EthTxManagerMock_Add_Call {
	_c.Call.Return(run)
	return _c
}

// ProcessPendingMonitoredTxs provides a mock function with given fields: ctx, owner, failedResultHandler, dbTx
func (_m *EthTxManagerMock) ProcessPendingMonitoredTxs(ctx context.Context, owner string, failedResultHandler ethtxmanager.ResultHandler, dbTx pgx.Tx) {
	_m.Called(ctx, owner, failedResultHandler, dbTx)
}

// EthTxManagerMock_ProcessPendingMonitoredTxs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ProcessPendingMonitoredTxs'
type EthTxManagerMock_ProcessPendingMonitoredTxs_Call struct {
	*mock.Call
}

// ProcessPendingMonitoredTxs is a helper method to define mock.On call
//   - ctx context.Context
//   - owner string
//   - failedResultHandler ethtxmanager.ResultHandler
//   - dbTx pgx.Tx
func (_e *EthTxManagerMock_Expecter) ProcessPendingMonitoredTxs(ctx interface{}, owner interface{}, failedResultHandler interface{}, dbTx interface{}) *EthTxManagerMock_ProcessPendingMonitoredTxs_Call {
	return &EthTxManagerMock_ProcessPendingMonitoredTxs_Call{Call: _e.mock.On("ProcessPendingMonitoredTxs", ctx, owner, failedResultHandler, dbTx)}
}

func (_c *EthTxManagerMock_ProcessPendingMonitoredTxs_Call) Run(run func(ctx context.Context, owner string, failedResultHandler ethtxmanager.ResultHandler, dbTx pgx.Tx)) *EthTxManagerMock_ProcessPendingMonitoredTxs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(ethtxmanager.ResultHandler), args[3].(pgx.Tx))
	})
	return _c
}

func (_c *EthTxManagerMock_ProcessPendingMonitoredTxs_Call) Return() *EthTxManagerMock_ProcessPendingMonitoredTxs_Call {
	_c.Call.Return()
	return _c
}

func (_c *EthTxManagerMock_ProcessPendingMonitoredTxs_Call) RunAndReturn(run func(context.Context, string, ethtxmanager.ResultHandler, pgx.Tx)) *EthTxManagerMock_ProcessPendingMonitoredTxs_Call {
	_c.Call.Return(run)
	return _c
}

// Result provides a mock function with given fields: ctx, owner, id, dbTx
func (_m *EthTxManagerMock) Result(ctx context.Context, owner string, id string, dbTx pgx.Tx) (ethtxmanager.MonitoredTxResult, error) {
	ret := _m.Called(ctx, owner, id, dbTx)

	if len(ret) == 0 {
		panic("no return value specified for Result")
	}

	var r0 ethtxmanager.MonitoredTxResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, pgx.Tx) (ethtxmanager.MonitoredTxResult, error)); ok {
		return rf(ctx, owner, id, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, pgx.Tx) ethtxmanager.MonitoredTxResult); ok {
		r0 = rf(ctx, owner, id, dbTx)
	} else {
		r0 = ret.Get(0).(ethtxmanager.MonitoredTxResult)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, pgx.Tx) error); ok {
		r1 = rf(ctx, owner, id, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EthTxManagerMock_Result_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Result'
type EthTxManagerMock_Result_Call struct {
	*mock.Call
}

// Result is a helper method to define mock.On call
//   - ctx context.Context
//   - owner string
//   - id string
//   - dbTx pgx.Tx
func (_e *EthTxManagerMock_Expecter) Result(ctx interface{}, owner interface{}, id interface{}, dbTx interface{}) *EthTxManagerMock_Result_Call {
	return &EthTxManagerMock_Result_Call{Call: _e.mock.On("Result", ctx, owner, id, dbTx)}
}

func (_c *EthTxManagerMock_Result_Call) Run(run func(ctx context.Context, owner string, id string, dbTx pgx.Tx)) *EthTxManagerMock_Result_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(pgx.Tx))
	})
	return _c
}

func (_c *EthTxManagerMock_Result_Call) Return(_a0 ethtxmanager.MonitoredTxResult, _a1 error) *EthTxManagerMock_Result_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *EthTxManagerMock_Result_Call) RunAndReturn(run func(context.Context, string, string, pgx.Tx) (ethtxmanager.MonitoredTxResult, error)) *EthTxManagerMock_Result_Call {
	_c.Call.Return(run)
	return _c
}

// ResultsByStatus provides a mock function with given fields: ctx, owner, statuses, dbTx
func (_m *EthTxManagerMock) ResultsByStatus(ctx context.Context, owner string, statuses []ethtxmanager.MonitoredTxStatus, dbTx pgx.Tx) ([]ethtxmanager.MonitoredTxResult, error) {
	ret := _m.Called(ctx, owner, statuses, dbTx)

	if len(ret) == 0 {
		panic("no return value specified for ResultsByStatus")
	}

	var r0 []ethtxmanager.MonitoredTxResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, []ethtxmanager.MonitoredTxStatus, pgx.Tx) ([]ethtxmanager.MonitoredTxResult, error)); ok {
		return rf(ctx, owner, statuses, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, []ethtxmanager.MonitoredTxStatus, pgx.Tx) []ethtxmanager.MonitoredTxResult); ok {
		r0 = rf(ctx, owner, statuses, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]ethtxmanager.MonitoredTxResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, []ethtxmanager.MonitoredTxStatus, pgx.Tx) error); ok {
		r1 = rf(ctx, owner, statuses, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EthTxManagerMock_ResultsByStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ResultsByStatus'
type EthTxManagerMock_ResultsByStatus_Call struct {
	*mock.Call
}

// ResultsByStatus is a helper method to define mock.On call
//   - ctx context.Context
//   - owner string
//   - statuses []ethtxmanager.MonitoredTxStatus
//   - dbTx pgx.Tx
func (_e *EthTxManagerMock_Expecter) ResultsByStatus(ctx interface{}, owner interface{}, statuses interface{}, dbTx interface{}) *EthTxManagerMock_ResultsByStatus_Call {
	return &EthTxManagerMock_ResultsByStatus_Call{Call: _e.mock.On("ResultsByStatus", ctx, owner, statuses, dbTx)}
}

func (_c *EthTxManagerMock_ResultsByStatus_Call) Run(run func(ctx context.Context, owner string, statuses []ethtxmanager.MonitoredTxStatus, dbTx pgx.Tx)) *EthTxManagerMock_ResultsByStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].([]ethtxmanager.MonitoredTxStatus), args[3].(pgx.Tx))
	})
	return _c
}

func (_c *EthTxManagerMock_ResultsByStatus_Call) Return(_a0 []ethtxmanager.MonitoredTxResult, _a1 error) *EthTxManagerMock_ResultsByStatus_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *EthTxManagerMock_ResultsByStatus_Call) RunAndReturn(run func(context.Context, string, []ethtxmanager.MonitoredTxStatus, pgx.Tx) ([]ethtxmanager.MonitoredTxResult, error)) *EthTxManagerMock_ResultsByStatus_Call {
	_c.Call.Return(run)
	return _c
}

// NewEthTxManagerMock creates a new instance of EthTxManagerMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEthTxManagerMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *EthTxManagerMock {
	mock := &EthTxManagerMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
