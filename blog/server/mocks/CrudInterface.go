// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	emptypb "google.golang.org/protobuf/types/known/emptypb"

	proto "github.com/MelquesPaiva/grpc-course/blog/proto"
)

// CrudInterface is an autogenerated mock type for the CrudInterface type
type CrudInterface struct {
	mock.Mock
}

// CreateBlog provides a mock function with given fields: ctx, blog
func (_m *CrudInterface) CreateBlog(ctx context.Context, blog *proto.Blog) (*proto.BlogId, error) {
	ret := _m.Called(ctx, blog)

	var r0 *proto.BlogId
	if rf, ok := ret.Get(0).(func(context.Context, *proto.Blog) *proto.BlogId); ok {
		r0 = rf(ctx, blog)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.BlogId)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *proto.Blog) error); ok {
		r1 = rf(ctx, blog)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteBlog provides a mock function with given fields: ctx, blogId
func (_m *CrudInterface) DeleteBlog(ctx context.Context, blogId *proto.BlogId) (*emptypb.Empty, error) {
	ret := _m.Called(ctx, blogId)

	var r0 *emptypb.Empty
	if rf, ok := ret.Get(0).(func(context.Context, *proto.BlogId) *emptypb.Empty); ok {
		r0 = rf(ctx, blogId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*emptypb.Empty)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *proto.BlogId) error); ok {
		r1 = rf(ctx, blogId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewCrudInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewCrudInterface creates a new instance of CrudInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCrudInterface(t mockConstructorTestingTNewCrudInterface) *CrudInterface {
	mock := &CrudInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}