// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	database "github.com/jchapman63/neo-gkmn/internal/database"
	mock "github.com/stretchr/testify/mock"
)

// MockQuerier is an autogenerated mock type for the Querier type
type MockQuerier struct {
	mock.Mock
}

type MockQuerier_Expecter struct {
	mock *mock.Mock
}

func (_m *MockQuerier) EXPECT() *MockQuerier_Expecter {
	return &MockQuerier_Expecter{mock: &_m.Mock}
}

// FetchMonster provides a mock function with given fields: ctx, id
func (_m *MockQuerier) FetchMonster(ctx context.Context, id string) (database.Monster, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for FetchMonster")
	}

	var r0 database.Monster
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (database.Monster, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) database.Monster); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(database.Monster)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQuerier_FetchMonster_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FetchMonster'
type MockQuerier_FetchMonster_Call struct {
	*mock.Call
}

// FetchMonster is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *MockQuerier_Expecter) FetchMonster(ctx interface{}, id interface{}) *MockQuerier_FetchMonster_Call {
	return &MockQuerier_FetchMonster_Call{Call: _e.mock.On("FetchMonster", ctx, id)}
}

func (_c *MockQuerier_FetchMonster_Call) Run(run func(ctx context.Context, id string)) *MockQuerier_FetchMonster_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockQuerier_FetchMonster_Call) Return(_a0 database.Monster, _a1 error) *MockQuerier_FetchMonster_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockQuerier_FetchMonster_Call) RunAndReturn(run func(context.Context, string) (database.Monster, error)) *MockQuerier_FetchMonster_Call {
	_c.Call.Return(run)
	return _c
}

// FetchMove provides a mock function with given fields: ctx, id
func (_m *MockQuerier) FetchMove(ctx context.Context, id string) (database.Move, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for FetchMove")
	}

	var r0 database.Move
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (database.Move, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) database.Move); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(database.Move)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQuerier_FetchMove_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FetchMove'
type MockQuerier_FetchMove_Call struct {
	*mock.Call
}

// FetchMove is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *MockQuerier_Expecter) FetchMove(ctx interface{}, id interface{}) *MockQuerier_FetchMove_Call {
	return &MockQuerier_FetchMove_Call{Call: _e.mock.On("FetchMove", ctx, id)}
}

func (_c *MockQuerier_FetchMove_Call) Run(run func(ctx context.Context, id string)) *MockQuerier_FetchMove_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockQuerier_FetchMove_Call) Return(_a0 database.Move, _a1 error) *MockQuerier_FetchMove_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockQuerier_FetchMove_Call) RunAndReturn(run func(context.Context, string) (database.Move, error)) *MockQuerier_FetchMove_Call {
	_c.Call.Return(run)
	return _c
}

// FetchMovesForMon provides a mock function with given fields: ctx, monsterid
func (_m *MockQuerier) FetchMovesForMon(ctx context.Context, monsterid string) ([]database.Movemap, error) {
	ret := _m.Called(ctx, monsterid)

	if len(ret) == 0 {
		panic("no return value specified for FetchMovesForMon")
	}

	var r0 []database.Movemap
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]database.Movemap, error)); ok {
		return rf(ctx, monsterid)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []database.Movemap); ok {
		r0 = rf(ctx, monsterid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]database.Movemap)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, monsterid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQuerier_FetchMovesForMon_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FetchMovesForMon'
type MockQuerier_FetchMovesForMon_Call struct {
	*mock.Call
}

// FetchMovesForMon is a helper method to define mock.On call
//   - ctx context.Context
//   - monsterid string
func (_e *MockQuerier_Expecter) FetchMovesForMon(ctx interface{}, monsterid interface{}) *MockQuerier_FetchMovesForMon_Call {
	return &MockQuerier_FetchMovesForMon_Call{Call: _e.mock.On("FetchMovesForMon", ctx, monsterid)}
}

func (_c *MockQuerier_FetchMovesForMon_Call) Run(run func(ctx context.Context, monsterid string)) *MockQuerier_FetchMovesForMon_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockQuerier_FetchMovesForMon_Call) Return(_a0 []database.Movemap, _a1 error) *MockQuerier_FetchMovesForMon_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockQuerier_FetchMovesForMon_Call) RunAndReturn(run func(context.Context, string) ([]database.Movemap, error)) *MockQuerier_FetchMovesForMon_Call {
	_c.Call.Return(run)
	return _c
}

// FetchStat provides a mock function with given fields: ctx, arg
func (_m *MockQuerier) FetchStat(ctx context.Context, arg database.FetchStatParams) (database.Stat, error) {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for FetchStat")
	}

	var r0 database.Stat
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, database.FetchStatParams) (database.Stat, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, database.FetchStatParams) database.Stat); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(database.Stat)
	}

	if rf, ok := ret.Get(1).(func(context.Context, database.FetchStatParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQuerier_FetchStat_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FetchStat'
type MockQuerier_FetchStat_Call struct {
	*mock.Call
}

// FetchStat is a helper method to define mock.On call
//   - ctx context.Context
//   - arg database.FetchStatParams
func (_e *MockQuerier_Expecter) FetchStat(ctx interface{}, arg interface{}) *MockQuerier_FetchStat_Call {
	return &MockQuerier_FetchStat_Call{Call: _e.mock.On("FetchStat", ctx, arg)}
}

func (_c *MockQuerier_FetchStat_Call) Run(run func(ctx context.Context, arg database.FetchStatParams)) *MockQuerier_FetchStat_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(database.FetchStatParams))
	})
	return _c
}

func (_c *MockQuerier_FetchStat_Call) Return(_a0 database.Stat, _a1 error) *MockQuerier_FetchStat_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockQuerier_FetchStat_Call) RunAndReturn(run func(context.Context, database.FetchStatParams) (database.Stat, error)) *MockQuerier_FetchStat_Call {
	_c.Call.Return(run)
	return _c
}

// ListMonsters provides a mock function with given fields: ctx
func (_m *MockQuerier) ListMonsters(ctx context.Context) ([]database.Monster, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for ListMonsters")
	}

	var r0 []database.Monster
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]database.Monster, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []database.Monster); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]database.Monster)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQuerier_ListMonsters_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListMonsters'
type MockQuerier_ListMonsters_Call struct {
	*mock.Call
}

// ListMonsters is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockQuerier_Expecter) ListMonsters(ctx interface{}) *MockQuerier_ListMonsters_Call {
	return &MockQuerier_ListMonsters_Call{Call: _e.mock.On("ListMonsters", ctx)}
}

func (_c *MockQuerier_ListMonsters_Call) Run(run func(ctx context.Context)) *MockQuerier_ListMonsters_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockQuerier_ListMonsters_Call) Return(_a0 []database.Monster, _a1 error) *MockQuerier_ListMonsters_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockQuerier_ListMonsters_Call) RunAndReturn(run func(context.Context) ([]database.Monster, error)) *MockQuerier_ListMonsters_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockQuerier creates a new instance of MockQuerier. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockQuerier(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockQuerier {
	mock := &MockQuerier{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
