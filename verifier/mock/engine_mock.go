// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/iost-official/go-iost/vm (interfaces: Engine)

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	tx "github.com/iost-official/go-iost/core/tx"
	reflect "reflect"
	time "time"
)

// MockEngine is a mock of Engine interface
type MockEngine struct {
	ctrl     *gomock.Controller
	recorder *MockEngineMockRecorder
}

// MockEngineMockRecorder is the mock recorder for MockEngine
type MockEngineMockRecorder struct {
	mock *MockEngine
}

// NewMockEngine creates a new mock instance
func NewMockEngine(ctrl *gomock.Controller) *MockEngine {
	mock := &MockEngine{ctrl: ctrl}
	mock.recorder = &MockEngineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEngine) EXPECT() *MockEngineMockRecorder {
	return m.recorder
}

// Exec mocks base method
func (m *MockEngine) Exec(arg0 *tx.Tx, arg1 time.Duration) (*tx.TxReceipt, error) {
	ret := m.ctrl.Call(m, "Exec", arg0, arg1)
	ret0, _ := ret[0].(*tx.TxReceipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exec indicates an expected call of Exec
func (mr *MockEngineMockRecorder) Exec(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exec", reflect.TypeOf((*MockEngine)(nil).Exec), arg0, arg1)
}

// GC mocks base method
func (m *MockEngine) GC() {
	m.ctrl.Call(m, "GC")
}

// GC indicates an expected call of GC
func (mr *MockEngineMockRecorder) GC() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GC", reflect.TypeOf((*MockEngine)(nil).GC))
}

// SetUp mocks base method
func (m *MockEngine) SetUp(arg0, arg1 string) error {
	ret := m.ctrl.Call(m, "SetUp", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetUp indicates an expected call of SetUp
func (mr *MockEngineMockRecorder) SetUp(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUp", reflect.TypeOf((*MockEngine)(nil).SetUp), arg0, arg1)
}
