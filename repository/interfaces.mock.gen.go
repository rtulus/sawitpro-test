// Code generated by MockGen. DO NOT EDIT.
// Source: repository/interfaces.go

// Package repository is a generated GoMock package.
package repository

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRepositoryInterface is a mock of RepositoryInterface interface.
type MockRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryInterfaceMockRecorder
}

// MockRepositoryInterfaceMockRecorder is the mock recorder for MockRepositoryInterface.
type MockRepositoryInterfaceMockRecorder struct {
	mock *MockRepositoryInterface
}

// NewMockRepositoryInterface creates a new mock instance.
func NewMockRepositoryInterface(ctrl *gomock.Controller) *MockRepositoryInterface {
	mock := &MockRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepositoryInterface) EXPECT() *MockRepositoryInterfaceMockRecorder {
	return m.recorder
}

// AddNewUser mocks base method.
func (m *MockRepositoryInterface) AddNewUser(ctx context.Context, input AddNewUserInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddNewUser", ctx, input)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddNewUser indicates an expected call of AddNewUser.
func (mr *MockRepositoryInterfaceMockRecorder) AddNewUser(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddNewUser", reflect.TypeOf((*MockRepositoryInterface)(nil).AddNewUser), ctx, input)
}

// GetTestById mocks base method.
func (m *MockRepositoryInterface) GetTestById(ctx context.Context, input GetTestByIdInput) (GetTestByIdOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTestById", ctx, input)
	ret0, _ := ret[0].(GetTestByIdOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTestById indicates an expected call of GetTestById.
func (mr *MockRepositoryInterfaceMockRecorder) GetTestById(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTestById", reflect.TypeOf((*MockRepositoryInterface)(nil).GetTestById), ctx, input)
}

// IncrementUserSuccessfulLogin mocks base method.
func (m *MockRepositoryInterface) IncrementUserSuccessfulLogin(ctx context.Context, input IncrementUserSuccessfulLoginInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IncrementUserSuccessfulLogin", ctx, input)
	ret0, _ := ret[0].(error)
	return ret0
}

// IncrementUserSuccessfulLogin indicates an expected call of IncrementUserSuccessfulLogin.
func (mr *MockRepositoryInterfaceMockRecorder) IncrementUserSuccessfulLogin(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncrementUserSuccessfulLogin", reflect.TypeOf((*MockRepositoryInterface)(nil).IncrementUserSuccessfulLogin), ctx, input)
}

// SelectUserByPhoneNumber mocks base method.
func (m *MockRepositoryInterface) SelectUserByPhoneNumber(ctx context.Context, input SelectUserByPhoneNumberInput) (Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectUserByPhoneNumber", ctx, input)
	ret0, _ := ret[0].(Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectUserByPhoneNumber indicates an expected call of SelectUserByPhoneNumber.
func (mr *MockRepositoryInterfaceMockRecorder) SelectUserByPhoneNumber(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectUserByPhoneNumber", reflect.TypeOf((*MockRepositoryInterface)(nil).SelectUserByPhoneNumber), ctx, input)
}

// UpdateFullName mocks base method.
func (m *MockRepositoryInterface) UpdateFullName(ctx context.Context, input UpdateFullNameInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFullName", ctx, input)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateFullName indicates an expected call of UpdateFullName.
func (mr *MockRepositoryInterfaceMockRecorder) UpdateFullName(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFullName", reflect.TypeOf((*MockRepositoryInterface)(nil).UpdateFullName), ctx, input)
}

// UpdatePhoneNumber mocks base method.
func (m *MockRepositoryInterface) UpdatePhoneNumber(ctx context.Context, input UpdatePhoneNumberInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePhoneNumber", ctx, input)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePhoneNumber indicates an expected call of UpdatePhoneNumber.
func (mr *MockRepositoryInterfaceMockRecorder) UpdatePhoneNumber(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePhoneNumber", reflect.TypeOf((*MockRepositoryInterface)(nil).UpdatePhoneNumber), ctx, input)
}