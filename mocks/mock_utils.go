// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/martinomburajr/pexels/utils (interfaces: BackgroundChanger,Filer,Rander,Utilizer)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockBackgroundChanger is a mock of BackgroundChanger interface
type MockBackgroundChanger struct {
	ctrl     *gomock.Controller
	recorder *MockBackgroundChangerMockRecorder
}

// MockBackgroundChangerMockRecorder is the mock recorder for MockBackgroundChanger
type MockBackgroundChangerMockRecorder struct {
	mock *MockBackgroundChanger
}

// NewMockBackgroundChanger creates a new mock instance
func NewMockBackgroundChanger(ctrl *gomock.Controller) *MockBackgroundChanger {
	mock := &MockBackgroundChanger{ctrl: ctrl}
	mock.recorder = &MockBackgroundChangerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBackgroundChanger) EXPECT() *MockBackgroundChangerMockRecorder {
	return m.recorder
}

// ChangeUbuntuBackground mocks base method
func (m *MockBackgroundChanger) ChangeUbuntuBackground(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeUbuntuBackground", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeUbuntuBackground indicates an expected call of ChangeUbuntuBackground
func (mr *MockBackgroundChangerMockRecorder) ChangeUbuntuBackground(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeUbuntuBackground", reflect.TypeOf((*MockBackgroundChanger)(nil).ChangeUbuntuBackground), arg0)
}

// MockFiler is a mock of Filer interface
type MockFiler struct {
	ctrl     *gomock.Controller
	recorder *MockFilerMockRecorder
}

// MockFilerMockRecorder is the mock recorder for MockFiler
type MockFilerMockRecorder struct {
	mock *MockFiler
}

// NewMockFiler creates a new mock instance
func NewMockFiler(ctrl *gomock.Controller) *MockFiler {
	mock := &MockFiler{ctrl: ctrl}
	mock.recorder = &MockFilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFiler) EXPECT() *MockFilerMockRecorder {
	return m.recorder
}

// WriteToFile mocks base method
func (m *MockFiler) WriteToFile(arg0 string, arg1 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteToFile", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteToFile indicates an expected call of WriteToFile
func (mr *MockFilerMockRecorder) WriteToFile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteToFile", reflect.TypeOf((*MockFiler)(nil).WriteToFile), arg0, arg1)
}

// MockRander is a mock of Rander interface
type MockRander struct {
	ctrl     *gomock.Controller
	recorder *MockRanderMockRecorder
}

// MockRanderMockRecorder is the mock recorder for MockRander
type MockRanderMockRecorder struct {
	mock *MockRander
}

// NewMockRander creates a new mock instance
func NewMockRander(ctrl *gomock.Controller) *MockRander {
	mock := &MockRander{ctrl: ctrl}
	mock.recorder = &MockRanderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRander) EXPECT() *MockRanderMockRecorder {
	return m.recorder
}

// RandBytes mocks base method
func (m *MockRander) RandBytes(arg0 int) []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RandBytes", arg0)
	ret0, _ := ret[0].([]byte)
	return ret0
}

// RandBytes indicates an expected call of RandBytes
func (mr *MockRanderMockRecorder) RandBytes(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RandBytes", reflect.TypeOf((*MockRander)(nil).RandBytes), arg0)
}

// RandInt mocks base method
func (m *MockRander) RandInt(arg0 int) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RandInt", arg0)
	ret0, _ := ret[0].(int)
	return ret0
}

// RandInt indicates an expected call of RandInt
func (mr *MockRanderMockRecorder) RandInt(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RandInt", reflect.TypeOf((*MockRander)(nil).RandInt), arg0)
}

// MockUtilizer is a mock of Utilizer interface
type MockUtilizer struct {
	ctrl     *gomock.Controller
	recorder *MockUtilizerMockRecorder
}

// MockUtilizerMockRecorder is the mock recorder for MockUtilizer
type MockUtilizerMockRecorder struct {
	mock *MockUtilizer
}

// NewMockUtilizer creates a new mock instance
func NewMockUtilizer(ctrl *gomock.Controller) *MockUtilizer {
	mock := &MockUtilizer{ctrl: ctrl}
	mock.recorder = &MockUtilizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUtilizer) EXPECT() *MockUtilizerMockRecorder {
	return m.recorder
}

// ChangeUbuntuBackground mocks base method
func (m *MockUtilizer) ChangeUbuntuBackground(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeUbuntuBackground", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeUbuntuBackground indicates an expected call of ChangeUbuntuBackground
func (mr *MockUtilizerMockRecorder) ChangeUbuntuBackground(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeUbuntuBackground", reflect.TypeOf((*MockUtilizer)(nil).ChangeUbuntuBackground), arg0)
}

// RandBytes mocks base method
func (m *MockUtilizer) RandBytes(arg0 int) []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RandBytes", arg0)
	ret0, _ := ret[0].([]byte)
	return ret0
}

// RandBytes indicates an expected call of RandBytes
func (mr *MockUtilizerMockRecorder) RandBytes(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RandBytes", reflect.TypeOf((*MockUtilizer)(nil).RandBytes), arg0)
}

// RandInt mocks base method
func (m *MockUtilizer) RandInt(arg0 int) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RandInt", arg0)
	ret0, _ := ret[0].(int)
	return ret0
}

// RandInt indicates an expected call of RandInt
func (mr *MockUtilizerMockRecorder) RandInt(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RandInt", reflect.TypeOf((*MockUtilizer)(nil).RandInt), arg0)
}

// WriteToFile mocks base method
func (m *MockUtilizer) WriteToFile(arg0 string, arg1 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteToFile", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteToFile indicates an expected call of WriteToFile
func (mr *MockUtilizerMockRecorder) WriteToFile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteToFile", reflect.TypeOf((*MockUtilizer)(nil).WriteToFile), arg0, arg1)
}