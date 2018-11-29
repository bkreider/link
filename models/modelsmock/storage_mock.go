// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Scalingo/link/models (interfaces: Storage)

// Package modelsmock is a generated GoMock package.
package modelsmock

import (
	context "context"
	reflect "reflect"

	models "github.com/Scalingo/link/models"
	gomock "github.com/golang/mock/gomock"
)

// MockStorage is a mock of Storage interface
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *MockStorageMockRecorder
}

// MockStorageMockRecorder is the mock recorder for MockStorage
type MockStorageMockRecorder struct {
	mock *MockStorage
}

// NewMockStorage creates a new mock instance
func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &MockStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStorage) EXPECT() *MockStorageMockRecorder {
	return m.recorder
}

// AddIP mocks base method
func (m *MockStorage) AddIP(arg0 context.Context, arg1 models.IP) (models.IP, error) {
	ret := m.ctrl.Call(m, "AddIP", arg0, arg1)
	ret0, _ := ret[0].(models.IP)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddIP indicates an expected call of AddIP
func (mr *MockStorageMockRecorder) AddIP(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddIP", reflect.TypeOf((*MockStorage)(nil).AddIP), arg0, arg1)
}

// GetIPs mocks base method
func (m *MockStorage) GetIPs(arg0 context.Context) ([]models.IP, error) {
	ret := m.ctrl.Call(m, "GetIPs", arg0)
	ret0, _ := ret[0].([]models.IP)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIPs indicates an expected call of GetIPs
func (mr *MockStorageMockRecorder) GetIPs(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIPs", reflect.TypeOf((*MockStorage)(nil).GetIPs), arg0)
}

// RemoveIP mocks base method
func (m *MockStorage) RemoveIP(arg0 context.Context, arg1 string) error {
	ret := m.ctrl.Call(m, "RemoveIP", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveIP indicates an expected call of RemoveIP
func (mr *MockStorageMockRecorder) RemoveIP(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveIP", reflect.TypeOf((*MockStorage)(nil).RemoveIP), arg0, arg1)
}