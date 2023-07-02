// Code generated by mockery v2.28.2. DO NOT EDIT.

package mocks

import (
	context "context"

	sqlx "github.com/jmoiron/sqlx"
	mock "github.com/stretchr/testify/mock"
	pkg "github.com/clavinjune/lokr/pkg"

	v1 "github.com/clavinjune/lokr/api/lokr/v1"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

type Repository_Expecter struct {
	mock *mock.Mock
}

func (_m *Repository) EXPECT() *Repository_Expecter {
	return &Repository_Expecter{mock: &_m.Mock}
}

// FetchLockTx provides a mock function with given fields: ctx, tx, key
func (_m *Repository) FetchLockTx(ctx context.Context, tx *sqlx.Tx, key string) (*v1.Lock, error) {
	ret := _m.Called(ctx, tx, key)

	var r0 *v1.Lock
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *sqlx.Tx, string) (*v1.Lock, error)); ok {
		return rf(ctx, tx, key)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *sqlx.Tx, string) *v1.Lock); ok {
		r0 = rf(ctx, tx, key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Lock)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *sqlx.Tx, string) error); ok {
		r1 = rf(ctx, tx, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Repository_FetchLockTx_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FetchLockTx'
type Repository_FetchLockTx_Call struct {
	*mock.Call
}

// FetchLockTx is a helper method to define mock.On call
//  - ctx context.Context
//  - tx *sqlx.Tx
//  - key string
func (_e *Repository_Expecter) FetchLockTx(ctx interface{}, tx interface{}, key interface{}) *Repository_FetchLockTx_Call {
	return &Repository_FetchLockTx_Call{Call: _e.mock.On("FetchLockTx", ctx, tx, key)}
}

func (_c *Repository_FetchLockTx_Call) Run(run func(ctx context.Context, tx *sqlx.Tx, key string)) *Repository_FetchLockTx_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*sqlx.Tx), args[2].(string))
	})
	return _c
}

func (_c *Repository_FetchLockTx_Call) Return(_a0 *v1.Lock, _a1 error) *Repository_FetchLockTx_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Repository_FetchLockTx_Call) RunAndReturn(run func(context.Context, *sqlx.Tx, string) (*v1.Lock, error)) *Repository_FetchLockTx_Call {
	_c.Call.Return(run)
	return _c
}

// PatchLockByKey provides a mock function with given fields: ctx, lock
func (_m *Repository) PatchLockByKey(ctx context.Context, lock *v1.Lock) error {
	ret := _m.Called(ctx, lock)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.Lock) error); ok {
		r0 = rf(ctx, lock)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Repository_PatchLockByKey_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PatchLockByKey'
type Repository_PatchLockByKey_Call struct {
	*mock.Call
}

// PatchLockByKey is a helper method to define mock.On call
//  - ctx context.Context
//  - lock *v1.Lock
func (_e *Repository_Expecter) PatchLockByKey(ctx interface{}, lock interface{}) *Repository_PatchLockByKey_Call {
	return &Repository_PatchLockByKey_Call{Call: _e.mock.On("PatchLockByKey", ctx, lock)}
}

func (_c *Repository_PatchLockByKey_Call) Run(run func(ctx context.Context, lock *v1.Lock)) *Repository_PatchLockByKey_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*v1.Lock))
	})
	return _c
}

func (_c *Repository_PatchLockByKey_Call) Return(_a0 error) *Repository_PatchLockByKey_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Repository_PatchLockByKey_Call) RunAndReturn(run func(context.Context, *v1.Lock) error) *Repository_PatchLockByKey_Call {
	_c.Call.Return(run)
	return _c
}

// PatchLockByKeyTx provides a mock function with given fields: ctx, tx, lock
func (_m *Repository) PatchLockByKeyTx(ctx context.Context, tx *sqlx.Tx, lock *v1.Lock) error {
	ret := _m.Called(ctx, tx, lock)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *sqlx.Tx, *v1.Lock) error); ok {
		r0 = rf(ctx, tx, lock)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Repository_PatchLockByKeyTx_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PatchLockByKeyTx'
type Repository_PatchLockByKeyTx_Call struct {
	*mock.Call
}

// PatchLockByKeyTx is a helper method to define mock.On call
//  - ctx context.Context
//  - tx *sqlx.Tx
//  - lock *v1.Lock
func (_e *Repository_Expecter) PatchLockByKeyTx(ctx interface{}, tx interface{}, lock interface{}) *Repository_PatchLockByKeyTx_Call {
	return &Repository_PatchLockByKeyTx_Call{Call: _e.mock.On("PatchLockByKeyTx", ctx, tx, lock)}
}

func (_c *Repository_PatchLockByKeyTx_Call) Run(run func(ctx context.Context, tx *sqlx.Tx, lock *v1.Lock)) *Repository_PatchLockByKeyTx_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*sqlx.Tx), args[2].(*v1.Lock))
	})
	return _c
}

func (_c *Repository_PatchLockByKeyTx_Call) Return(_a0 error) *Repository_PatchLockByKeyTx_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Repository_PatchLockByKeyTx_Call) RunAndReturn(run func(context.Context, *sqlx.Tx, *v1.Lock) error) *Repository_PatchLockByKeyTx_Call {
	_c.Call.Return(run)
	return _c
}

// StoreLock provides a mock function with given fields: ctx, lock
func (_m *Repository) StoreLock(ctx context.Context, lock *v1.Lock) error {
	ret := _m.Called(ctx, lock)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.Lock) error); ok {
		r0 = rf(ctx, lock)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Repository_StoreLock_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StoreLock'
type Repository_StoreLock_Call struct {
	*mock.Call
}

// StoreLock is a helper method to define mock.On call
//  - ctx context.Context
//  - lock *v1.Lock
func (_e *Repository_Expecter) StoreLock(ctx interface{}, lock interface{}) *Repository_StoreLock_Call {
	return &Repository_StoreLock_Call{Call: _e.mock.On("StoreLock", ctx, lock)}
}

func (_c *Repository_StoreLock_Call) Run(run func(ctx context.Context, lock *v1.Lock)) *Repository_StoreLock_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*v1.Lock))
	})
	return _c
}

func (_c *Repository_StoreLock_Call) Return(_a0 error) *Repository_StoreLock_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Repository_StoreLock_Call) RunAndReturn(run func(context.Context, *v1.Lock) error) *Repository_StoreLock_Call {
	_c.Call.Return(run)
	return _c
}

// Tx provides a mock function with given fields: ctx, fn
func (_m *Repository) Tx(ctx context.Context, fn pkg.TxHandler) error {
	ret := _m.Called(ctx, fn)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, pkg.TxHandler) error); ok {
		r0 = rf(ctx, fn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Repository_Tx_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Tx'
type Repository_Tx_Call struct {
	*mock.Call
}

// Tx is a helper method to define mock.On call
//  - ctx context.Context
//  - fn pkg.TxHandler
func (_e *Repository_Expecter) Tx(ctx interface{}, fn interface{}) *Repository_Tx_Call {
	return &Repository_Tx_Call{Call: _e.mock.On("Tx", ctx, fn)}
}

func (_c *Repository_Tx_Call) Run(run func(ctx context.Context, fn pkg.TxHandler)) *Repository_Tx_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(pkg.TxHandler))
	})
	return _c
}

func (_c *Repository_Tx_Call) Return(_a0 error) *Repository_Tx_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Repository_Tx_Call) RunAndReturn(run func(context.Context, pkg.TxHandler) error) *Repository_Tx_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepository(t mockConstructorTestingTNewRepository) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
