// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// UpdateResultAPI is an autogenerated mock type for the UpdateResultAPI type
type UpdateResultAPI struct {
	mock.Mock
}

type mockConstructorTestingTNewUpdateResultAPI interface {
	mock.TestingT
	Cleanup(func())
}

// NewUpdateResultAPI creates a new instance of UpdateResultAPI. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUpdateResultAPI(t mockConstructorTestingTNewUpdateResultAPI) *UpdateResultAPI {
	mock := &UpdateResultAPI{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
