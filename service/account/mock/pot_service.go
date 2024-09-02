// Code generated by MockGen. DO NOT EDIT.
// Source: pot_service.go

// Package mock_account is a generated GoMock package.
package mock_account

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockPotService is a mock of PotService interface.
type MockPotService struct {
	ctrl     *gomock.Controller
	recorder *MockPotServiceMockRecorder
}

// MockPotServiceMockRecorder is the mock recorder for MockPotService.
type MockPotServiceMockRecorder struct {
	mock *MockPotService
}

// NewMockPotService creates a new mock instance.
func NewMockPotService(ctrl *gomock.Controller) *MockPotService {
	mock := &MockPotService{ctrl: ctrl}
	mock.recorder = &MockPotServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPotService) EXPECT() *MockPotServiceMockRecorder {
	return m.recorder
}

// GetOrCreatePot mocks base method.
func (m *MockPotService) GetOrCreatePot(id int, name string, clientId, depositId int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrCreatePot", id, name, clientId, depositId)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrCreatePot indicates an expected call of GetOrCreatePot.
func (mr *MockPotServiceMockRecorder) GetOrCreatePot(id, name, clientId, depositId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrCreatePot", reflect.TypeOf((*MockPotService)(nil).GetOrCreatePot), id, name, clientId, depositId)
}
