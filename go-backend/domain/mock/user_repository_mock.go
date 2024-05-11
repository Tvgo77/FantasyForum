// Code generated by MockGen. DO NOT EDIT.
// Source: user_repository_domain.go

// Package mock_domain is a generated GoMock package.
package mock_domain

import (
	context "context"
	domain "go-backend/domain"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// CheckExistByEmail mocks base method.
func (m *MockUserRepository) CheckExistByEmail(ctx context.Context, email string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckExistByEmail", ctx, email)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckExistByEmail indicates an expected call of CheckExistByEmail.
func (mr *MockUserRepositoryMockRecorder) CheckExistByEmail(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckExistByEmail", reflect.TypeOf((*MockUserRepository)(nil).CheckExistByEmail), ctx, email)
}

// Create mocks base method.
func (m *MockUserRepository) Create(arg0 context.Context, arg1 *domain.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockUserRepositoryMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserRepository)(nil).Create), arg0, arg1)
}

// Fetch mocks base method.
func (m *MockUserRepository) Fetch(ctx context.Context, conds *domain.User) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fetch", ctx, conds)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Fetch indicates an expected call of Fetch.
func (mr *MockUserRepositoryMockRecorder) Fetch(ctx, conds interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fetch", reflect.TypeOf((*MockUserRepository)(nil).Fetch), ctx, conds)
}
