// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// InsertOneResultAPI is an autogenerated mock type for the InsertOneResultAPI type
type InsertOneResultAPI struct {
	mock.Mock
}

type mockConstructorTestingTNewInsertOneResultAPI interface {
	mock.TestingT
	Cleanup(func())
}

// NewInsertOneResultAPI creates a new instance of InsertOneResultAPI. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewInsertOneResultAPI(t mockConstructorTestingTNewInsertOneResultAPI) *InsertOneResultAPI {
	mock := &InsertOneResultAPI{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}