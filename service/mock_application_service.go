// Code generated by mockery v1.0.0. DO NOT EDIT.

package service

import context "context"
import mock "github.com/stretchr/testify/mock"
import model "bitbucket.org/accezz-io/sac-operator/model"
import uuid "github.com/google/uuid"

// MockApplicationService is an autogenerated mock type for the ApplicationService type
type MockApplicationService struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, applicationToCreate
func (_m *MockApplicationService) Create(ctx context.Context, applicationToCreate *model.Application) (*model.Application, error) {
	ret := _m.Called(ctx, applicationToCreate)

	var r0 *model.Application
	if rf, ok := ret.Get(0).(func(context.Context, *model.Application) *model.Application); ok {
		r0 = rf(ctx, applicationToCreate)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Application)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *model.Application) error); ok {
		r1 = rf(ctx, applicationToCreate)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *MockApplicationService) Delete(ctx context.Context, id uuid.UUID) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: ctx, updatedApplication
func (_m *MockApplicationService) Update(ctx context.Context, updatedApplication *model.Application) (*model.Application, error) {
	ret := _m.Called(ctx, updatedApplication)

	var r0 *model.Application
	if rf, ok := ret.Get(0).(func(context.Context, *model.Application) *model.Application); ok {
		r0 = rf(ctx, updatedApplication)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Application)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *model.Application) error); ok {
		r1 = rf(ctx, updatedApplication)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
