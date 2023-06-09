// Code generated by MockGen. DO NOT EDIT.
// Source: usecase/usecase.go

// Package mocks is a generated GoMock package.
package mocks

import (
	models "nexter-assignment/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// GetCashRegister mocks base method.
func (m *MockRepository) GetCashRegister() (models.CashRegister, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCashRegister")
	ret0, _ := ret[0].(models.CashRegister)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCashRegister indicates an expected call of GetCashRegister.
func (mr *MockRepositoryMockRecorder) GetCashRegister() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCashRegister", reflect.TypeOf((*MockRepository)(nil).GetCashRegister))
}

// SaveCashRegister mocks base method.
func (m *MockRepository) SaveCashRegister(arg0 models.CashRegister) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveCashRegister", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveCashRegister indicates an expected call of SaveCashRegister.
func (mr *MockRepositoryMockRecorder) SaveCashRegister(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveCashRegister", reflect.TypeOf((*MockRepository)(nil).SaveCashRegister), arg0)
}
