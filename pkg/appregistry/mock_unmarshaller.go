// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/appregistry/unmarshaller.go

// Package appregistry is a generated GoMock package.
package appregistry

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockblobUnmarshaller is a mock of blobUnmarshaller interface
type MockblobUnmarshaller struct {
	ctrl     *gomock.Controller
	recorder *MockblobUnmarshallerMockRecorder
}

// MockblobUnmarshallerMockRecorder is the mock recorder for MockblobUnmarshaller
type MockblobUnmarshallerMockRecorder struct {
	mock *MockblobUnmarshaller
}

// NewMockblobUnmarshaller creates a new mock instance
func NewMockblobUnmarshaller(ctrl *gomock.Controller) *MockblobUnmarshaller {
	mock := &MockblobUnmarshaller{ctrl: ctrl}
	mock.recorder = &MockblobUnmarshallerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockblobUnmarshaller) EXPECT() *MockblobUnmarshallerMockRecorder {
	return m.recorder
}

// Unmarshal mocks base method
func (m *MockblobUnmarshaller) Unmarshal(in []byte) (*Manifest, error) {
	ret := m.ctrl.Call(m, "Unmarshal", in)
	ret0, _ := ret[0].(*Manifest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Unmarshal indicates an expected call of Unmarshal
func (mr *MockblobUnmarshallerMockRecorder) Unmarshal(in interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unmarshal", reflect.TypeOf((*MockblobUnmarshaller)(nil).Unmarshal), in)
}
