// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/martinomburajr/pexels/pexels (interfaces: Pexeler,GetRandomPexeler)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockPexeler is a mock of Pexeler interface
type MockPexeler struct {
	ctrl     *gomock.Controller
	recorder *MockPexelerMockRecorder
}

// MockPexelerMockRecorder is the mock recorder for MockPexeler
type MockPexelerMockRecorder struct {
	mock *MockPexeler
}

// NewMockPexeler creates a new mock instance
func NewMockPexeler(ctrl *gomock.Controller) *MockPexeler {
	mock := &MockPexeler{ctrl: ctrl}
	mock.recorder = &MockPexelerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPexeler) EXPECT() *MockPexelerMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockPexeler) Get(arg0 int, arg1 string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockPexelerMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockPexeler)(nil).Get), arg0, arg1)
}

// GetBySize mocks base method
func (m *MockPexeler) GetBySize(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBySize", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetBySize indicates an expected call of GetBySize
func (mr *MockPexelerMockRecorder) GetBySize(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBySize", reflect.TypeOf((*MockPexeler)(nil).GetBySize), arg0)
}

// GetRandomImage mocks base method
func (m *MockPexeler) GetRandomImage(arg0 string) (int, []byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRandomImage", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].([]byte)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetRandomImage indicates an expected call of GetRandomImage
func (mr *MockPexelerMockRecorder) GetRandomImage(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRandomImage", reflect.TypeOf((*MockPexeler)(nil).GetRandomImage), arg0)
}

// MockGetRandomPexeler is a mock of GetRandomPexeler interface
type MockGetRandomPexeler struct {
	ctrl     *gomock.Controller
	recorder *MockGetRandomPexelerMockRecorder
}

// MockGetRandomPexelerMockRecorder is the mock recorder for MockGetRandomPexeler
type MockGetRandomPexelerMockRecorder struct {
	mock *MockGetRandomPexeler
}

// NewMockGetRandomPexeler creates a new mock instance
func NewMockGetRandomPexeler(ctrl *gomock.Controller) *MockGetRandomPexeler {
	mock := &MockGetRandomPexeler{ctrl: ctrl}
	mock.recorder = &MockGetRandomPexelerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGetRandomPexeler) EXPECT() *MockGetRandomPexelerMockRecorder {
	return m.recorder
}

// GetRandomImage mocks base method
func (m *MockGetRandomPexeler) GetRandomImage(arg0 string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRandomImage", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRandomImage indicates an expected call of GetRandomImage
func (mr *MockGetRandomPexelerMockRecorder) GetRandomImage(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRandomImage", reflect.TypeOf((*MockGetRandomPexeler)(nil).GetRandomImage), arg0)
}
