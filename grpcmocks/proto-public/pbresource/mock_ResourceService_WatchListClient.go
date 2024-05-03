// Code generated by mockery v2.41.0. DO NOT EDIT.

package mockpbresource

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	metadata "google.golang.org/grpc/metadata"

	pbresource "github.com/hashicorp/consul/proto-public/pbresource"
)

// ResourceService_WatchListClient is an autogenerated mock type for the ResourceService_WatchListClient type
type ResourceService_WatchListClient struct {
	mock.Mock
}

type ResourceService_WatchListClient_Expecter struct {
	mock *mock.Mock
}

func (_m *ResourceService_WatchListClient) EXPECT() *ResourceService_WatchListClient_Expecter {
	return &ResourceService_WatchListClient_Expecter{mock: &_m.Mock}
}

// CloseSend provides a mock function with given fields:
func (_m *ResourceService_WatchListClient) CloseSend() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for CloseSend")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ResourceService_WatchListClient_CloseSend_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CloseSend'
type ResourceService_WatchListClient_CloseSend_Call struct {
	*mock.Call
}

// CloseSend is a helper method to define mock.On call
func (_e *ResourceService_WatchListClient_Expecter) CloseSend() *ResourceService_WatchListClient_CloseSend_Call {
	return &ResourceService_WatchListClient_CloseSend_Call{Call: _e.mock.On("CloseSend")}
}

func (_c *ResourceService_WatchListClient_CloseSend_Call) Run(run func()) *ResourceService_WatchListClient_CloseSend_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ResourceService_WatchListClient_CloseSend_Call) Return(_a0 error) *ResourceService_WatchListClient_CloseSend_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ResourceService_WatchListClient_CloseSend_Call) RunAndReturn(run func() error) *ResourceService_WatchListClient_CloseSend_Call {
	_c.Call.Return(run)
	return _c
}

// Context provides a mock function with given fields:
func (_m *ResourceService_WatchListClient) Context() context.Context {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Context")
	}

	var r0 context.Context
	if rf, ok := ret.Get(0).(func() context.Context); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	return r0
}

// ResourceService_WatchListClient_Context_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Context'
type ResourceService_WatchListClient_Context_Call struct {
	*mock.Call
}

// Context is a helper method to define mock.On call
func (_e *ResourceService_WatchListClient_Expecter) Context() *ResourceService_WatchListClient_Context_Call {
	return &ResourceService_WatchListClient_Context_Call{Call: _e.mock.On("Context")}
}

func (_c *ResourceService_WatchListClient_Context_Call) Run(run func()) *ResourceService_WatchListClient_Context_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ResourceService_WatchListClient_Context_Call) Return(_a0 context.Context) *ResourceService_WatchListClient_Context_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ResourceService_WatchListClient_Context_Call) RunAndReturn(run func() context.Context) *ResourceService_WatchListClient_Context_Call {
	_c.Call.Return(run)
	return _c
}

// Header provides a mock function with given fields:
func (_m *ResourceService_WatchListClient) Header() (metadata.MD, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Header")
	}

	var r0 metadata.MD
	var r1 error
	if rf, ok := ret.Get(0).(func() (metadata.MD, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() metadata.MD); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(metadata.MD)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResourceService_WatchListClient_Header_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Header'
type ResourceService_WatchListClient_Header_Call struct {
	*mock.Call
}

// Header is a helper method to define mock.On call
func (_e *ResourceService_WatchListClient_Expecter) Header() *ResourceService_WatchListClient_Header_Call {
	return &ResourceService_WatchListClient_Header_Call{Call: _e.mock.On("Header")}
}

func (_c *ResourceService_WatchListClient_Header_Call) Run(run func()) *ResourceService_WatchListClient_Header_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ResourceService_WatchListClient_Header_Call) Return(_a0 metadata.MD, _a1 error) *ResourceService_WatchListClient_Header_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ResourceService_WatchListClient_Header_Call) RunAndReturn(run func() (metadata.MD, error)) *ResourceService_WatchListClient_Header_Call {
	_c.Call.Return(run)
	return _c
}

// Recv provides a mock function with given fields:
func (_m *ResourceService_WatchListClient) Recv() (*pbresource.WatchEvent, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Recv")
	}

	var r0 *pbresource.WatchEvent
	var r1 error
	if rf, ok := ret.Get(0).(func() (*pbresource.WatchEvent, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *pbresource.WatchEvent); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pbresource.WatchEvent)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResourceService_WatchListClient_Recv_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Recv'
type ResourceService_WatchListClient_Recv_Call struct {
	*mock.Call
}

// Recv is a helper method to define mock.On call
func (_e *ResourceService_WatchListClient_Expecter) Recv() *ResourceService_WatchListClient_Recv_Call {
	return &ResourceService_WatchListClient_Recv_Call{Call: _e.mock.On("Recv")}
}

func (_c *ResourceService_WatchListClient_Recv_Call) Run(run func()) *ResourceService_WatchListClient_Recv_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ResourceService_WatchListClient_Recv_Call) Return(_a0 *pbresource.WatchEvent, _a1 error) *ResourceService_WatchListClient_Recv_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ResourceService_WatchListClient_Recv_Call) RunAndReturn(run func() (*pbresource.WatchEvent, error)) *ResourceService_WatchListClient_Recv_Call {
	_c.Call.Return(run)
	return _c
}

// RecvMsg provides a mock function with given fields: m
func (_m *ResourceService_WatchListClient) RecvMsg(m interface{}) error {
	ret := _m.Called(m)

	if len(ret) == 0 {
		panic("no return value specified for RecvMsg")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ResourceService_WatchListClient_RecvMsg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RecvMsg'
type ResourceService_WatchListClient_RecvMsg_Call struct {
	*mock.Call
}

// RecvMsg is a helper method to define mock.On call
//   - m interface{}
func (_e *ResourceService_WatchListClient_Expecter) RecvMsg(m interface{}) *ResourceService_WatchListClient_RecvMsg_Call {
	return &ResourceService_WatchListClient_RecvMsg_Call{Call: _e.mock.On("RecvMsg", m)}
}

func (_c *ResourceService_WatchListClient_RecvMsg_Call) Run(run func(m interface{})) *ResourceService_WatchListClient_RecvMsg_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *ResourceService_WatchListClient_RecvMsg_Call) Return(_a0 error) *ResourceService_WatchListClient_RecvMsg_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ResourceService_WatchListClient_RecvMsg_Call) RunAndReturn(run func(interface{}) error) *ResourceService_WatchListClient_RecvMsg_Call {
	_c.Call.Return(run)
	return _c
}

// SendMsg provides a mock function with given fields: m
func (_m *ResourceService_WatchListClient) SendMsg(m interface{}) error {
	ret := _m.Called(m)

	if len(ret) == 0 {
		panic("no return value specified for SendMsg")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ResourceService_WatchListClient_SendMsg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendMsg'
type ResourceService_WatchListClient_SendMsg_Call struct {
	*mock.Call
}

// SendMsg is a helper method to define mock.On call
//   - m interface{}
func (_e *ResourceService_WatchListClient_Expecter) SendMsg(m interface{}) *ResourceService_WatchListClient_SendMsg_Call {
	return &ResourceService_WatchListClient_SendMsg_Call{Call: _e.mock.On("SendMsg", m)}
}

func (_c *ResourceService_WatchListClient_SendMsg_Call) Run(run func(m interface{})) *ResourceService_WatchListClient_SendMsg_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *ResourceService_WatchListClient_SendMsg_Call) Return(_a0 error) *ResourceService_WatchListClient_SendMsg_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ResourceService_WatchListClient_SendMsg_Call) RunAndReturn(run func(interface{}) error) *ResourceService_WatchListClient_SendMsg_Call {
	_c.Call.Return(run)
	return _c
}

// Trailer provides a mock function with given fields:
func (_m *ResourceService_WatchListClient) Trailer() metadata.MD {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Trailer")
	}

	var r0 metadata.MD
	if rf, ok := ret.Get(0).(func() metadata.MD); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(metadata.MD)
		}
	}

	return r0
}

// ResourceService_WatchListClient_Trailer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Trailer'
type ResourceService_WatchListClient_Trailer_Call struct {
	*mock.Call
}

// Trailer is a helper method to define mock.On call
func (_e *ResourceService_WatchListClient_Expecter) Trailer() *ResourceService_WatchListClient_Trailer_Call {
	return &ResourceService_WatchListClient_Trailer_Call{Call: _e.mock.On("Trailer")}
}

func (_c *ResourceService_WatchListClient_Trailer_Call) Run(run func()) *ResourceService_WatchListClient_Trailer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ResourceService_WatchListClient_Trailer_Call) Return(_a0 metadata.MD) *ResourceService_WatchListClient_Trailer_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ResourceService_WatchListClient_Trailer_Call) RunAndReturn(run func() metadata.MD) *ResourceService_WatchListClient_Trailer_Call {
	_c.Call.Return(run)
	return _c
}

// NewResourceService_WatchListClient creates a new instance of ResourceService_WatchListClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewResourceService_WatchListClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *ResourceService_WatchListClient {
	mock := &ResourceService_WatchListClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}