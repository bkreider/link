// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Scalingo/go-utils/nsqproducer (interfaces: Producer)

// Package nsqproducermock is a generated GoMock package.
package nsqproducermock

import (
	context "context"
	nsqproducer "github.com/Scalingo/go-utils/nsqproducer"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockProducer is a mock of Producer interface
type MockProducer struct {
	ctrl     *gomock.Controller
	recorder *MockProducerMockRecorder
}

// MockProducerMockRecorder is the mock recorder for MockProducer
type MockProducerMockRecorder struct {
	mock *MockProducer
}

// NewMockProducer creates a new mock instance
func NewMockProducer(ctrl *gomock.Controller) *MockProducer {
	mock := &MockProducer{ctrl: ctrl}
	mock.recorder = &MockProducerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProducer) EXPECT() *MockProducerMockRecorder {
	return m.recorder
}

// DeferredPublish mocks base method
func (m *MockProducer) DeferredPublish(arg0 context.Context, arg1 string, arg2 int64, arg3 nsqproducer.NsqMessageSerialize) error {
	ret := m.ctrl.Call(m, "DeferredPublish", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeferredPublish indicates an expected call of DeferredPublish
func (mr *MockProducerMockRecorder) DeferredPublish(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeferredPublish", reflect.TypeOf((*MockProducer)(nil).DeferredPublish), arg0, arg1, arg2, arg3)
}

// Publish mocks base method
func (m *MockProducer) Publish(arg0 context.Context, arg1 string, arg2 nsqproducer.NsqMessageSerialize) error {
	ret := m.ctrl.Call(m, "Publish", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Publish indicates an expected call of Publish
func (mr *MockProducerMockRecorder) Publish(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockProducer)(nil).Publish), arg0, arg1, arg2)
}

// Stop mocks base method
func (m *MockProducer) Stop() {
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop
func (mr *MockProducerMockRecorder) Stop() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockProducer)(nil).Stop))
}
