// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	context "context"
	model "warehouse/internal/domain/model"

	mock "github.com/stretchr/testify/mock"

	primitive "go.mongodb.org/mongo-driver/bson/primitive"

	request "warehouse/internal/domain/net/request"
)

// Category is an autogenerated mock type for the Category type
type Category struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, category
func (_m *Category) Create(ctx context.Context, category model.Category) (primitive.ObjectID, error) {
	ret := _m.Called(ctx, category)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 primitive.ObjectID
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.Category) (primitive.ObjectID, error)); ok {
		return rf(ctx, category)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.Category) primitive.ObjectID); ok {
		r0 = rf(ctx, category)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(primitive.ObjectID)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.Category) error); ok {
		r1 = rf(ctx, category)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *Category) Delete(ctx context.Context, id primitive.ObjectID) (primitive.ObjectID, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 primitive.ObjectID
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID) (primitive.ObjectID, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID) primitive.ObjectID); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(primitive.ObjectID)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, primitive.ObjectID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields: ctx, req
func (_m *Category) GetAll(ctx context.Context, req request.GetCategories) ([]model.Category, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
	}

	var r0 []model.Category
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, request.GetCategories) ([]model.Category, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, request.GetCategories) []model.Category); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Category)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, request.GetCategories) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByName provides a mock function with given fields: ctx, name
func (_m *Category) GetByName(ctx context.Context, name string) (model.Category, error) {
	ret := _m.Called(ctx, name)

	if len(ret) == 0 {
		panic("no return value specified for GetByName")
	}

	var r0 model.Category
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (model.Category, error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) model.Category); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Get(0).(model.Category)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewCategory creates a new instance of Category. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCategory(t interface {
	mock.TestingT
	Cleanup(func())
}) *Category {
	mock := &Category{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}