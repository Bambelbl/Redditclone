// Code generated by MockGen. DO NOT EDIT.
// Source: session.go

// Package session is a generated GoMock package.
package session

import (
	user "redditclone/pkg/user"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockSessionsRepo is a mock of SessionsRepo interface.
type MockSessionsRepo struct {
	ctrl     *gomock.Controller
	recorder *MockSessionsRepoMockRecorder
}

// MockSessionsRepoMockRecorder is the mock recorder for MockSessionsRepo.
type MockSessionsRepoMockRecorder struct {
	mock *MockSessionsRepo
}

// NewMockSessionsRepo creates a new mock instance.
func NewMockSessionsRepo(ctrl *gomock.Controller) *MockSessionsRepo {
	mock := &MockSessionsRepo{ctrl: ctrl}
	mock.recorder = &MockSessionsRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSessionsRepo) EXPECT() *MockSessionsRepoMockRecorder {
	return m.recorder
}

// Check mocks base method.
func (m *MockSessionsRepo) Check(arg0 string) (*Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Check", arg0)
	ret0, _ := ret[0].(*Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Check indicates an expected call of Check.
func (mr *MockSessionsRepoMockRecorder) Check(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Check", reflect.TypeOf((*MockSessionsRepo)(nil).Check), arg0)
}

// Create mocks base method.
func (m *MockSessionsRepo) Create(arg0 user.User) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockSessionsRepoMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockSessionsRepo)(nil).Create), arg0)
}
