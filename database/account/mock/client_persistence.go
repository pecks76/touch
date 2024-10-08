// Code generated by MockGen. DO NOT EDIT.
// Source: client_persistence.go

// Package mock_account is a generated GoMock package.
package mock_account

import (
	reflect "reflect"
	account "restservice/domain/account"

	gomock "github.com/golang/mock/gomock"
)

// MockClientRepository is a mock of ClientRepository interface.
type MockClientRepository struct {
	ctrl     *gomock.Controller
	recorder *MockClientRepositoryMockRecorder
}

// MockClientRepositoryMockRecorder is the mock recorder for MockClientRepository.
type MockClientRepositoryMockRecorder struct {
	mock *MockClientRepository
}

// NewMockClientRepository creates a new mock instance.
func NewMockClientRepository(ctrl *gomock.Controller) *MockClientRepository {
	mock := &MockClientRepository{ctrl: ctrl}
	mock.recorder = &MockClientRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClientRepository) EXPECT() *MockClientRepositoryMockRecorder {
	return m.recorder
}

// InsertClient mocks base method.
func (m *MockClientRepository) InsertClient() (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertClient")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertClient indicates an expected call of InsertClient.
func (mr *MockClientRepositoryMockRecorder) InsertClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertClient", reflect.TypeOf((*MockClientRepository)(nil).InsertClient))
}

// ReadClient mocks base method.
func (m *MockClientRepository) ReadClient(id int) account.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadClient", id)
	ret0, _ := ret[0].(account.Client)
	return ret0
}

// ReadClient indicates an expected call of ReadClient.
func (mr *MockClientRepositoryMockRecorder) ReadClient(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadClient", reflect.TypeOf((*MockClientRepository)(nil).ReadClient), id)
}
