// Code generated by mockery v2.40.3. DO NOT EDIT.

package mocks

import (
	cache "github.com/flyteorg/flyte/flytestdlib/cache"
	mock "github.com/stretchr/testify/mock"
)

// ItemWrapper is an autogenerated mock type for the ItemWrapper type
type ItemWrapper struct {
	mock.Mock
}

type ItemWrapper_Expecter struct {
	mock *mock.Mock
}

func (_m *ItemWrapper) EXPECT() *ItemWrapper_Expecter {
	return &ItemWrapper_Expecter{mock: &_m.Mock}
}

// GetID provides a mock function with given fields:
func (_m *ItemWrapper) GetID() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetID")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ItemWrapper_GetID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetID'
type ItemWrapper_GetID_Call struct {
	*mock.Call
}

// GetID is a helper method to define mock.On call
func (_e *ItemWrapper_Expecter) GetID() *ItemWrapper_GetID_Call {
	return &ItemWrapper_GetID_Call{Call: _e.mock.On("GetID")}
}

func (_c *ItemWrapper_GetID_Call) Run(run func()) *ItemWrapper_GetID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ItemWrapper_GetID_Call) Return(_a0 string) *ItemWrapper_GetID_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ItemWrapper_GetID_Call) RunAndReturn(run func() string) *ItemWrapper_GetID_Call {
	_c.Call.Return(run)
	return _c
}

// GetItem provides a mock function with given fields:
func (_m *ItemWrapper) GetItem() cache.Item {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetItem")
	}

	var r0 cache.Item
	if rf, ok := ret.Get(0).(func() cache.Item); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(cache.Item)
		}
	}

	return r0
}

// ItemWrapper_GetItem_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetItem'
type ItemWrapper_GetItem_Call struct {
	*mock.Call
}

// GetItem is a helper method to define mock.On call
func (_e *ItemWrapper_Expecter) GetItem() *ItemWrapper_GetItem_Call {
	return &ItemWrapper_GetItem_Call{Call: _e.mock.On("GetItem")}
}

func (_c *ItemWrapper_GetItem_Call) Run(run func()) *ItemWrapper_GetItem_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ItemWrapper_GetItem_Call) Return(_a0 cache.Item) *ItemWrapper_GetItem_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ItemWrapper_GetItem_Call) RunAndReturn(run func() cache.Item) *ItemWrapper_GetItem_Call {
	_c.Call.Return(run)
	return _c
}

// NewItemWrapper creates a new instance of ItemWrapper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewItemWrapper(t interface {
	mock.TestingT
	Cleanup(func())
}) *ItemWrapper {
	mock := &ItemWrapper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
