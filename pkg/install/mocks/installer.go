// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/rancher/elemental-operator/pkg/install (interfaces: Installer)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	v1beta1 "github.com/rancher/elemental-operator/api/v1beta1"
	gomock "go.uber.org/mock/gomock"
)

// MockInstaller is a mock of Installer interface.
type MockInstaller struct {
	ctrl     *gomock.Controller
	recorder *MockInstallerMockRecorder
}

// MockInstallerMockRecorder is the mock recorder for MockInstaller.
type MockInstallerMockRecorder struct {
	mock *MockInstaller
}

// NewMockInstaller creates a new mock instance.
func NewMockInstaller(ctrl *gomock.Controller) *MockInstaller {
	mock := &MockInstaller{ctrl: ctrl}
	mock.recorder = &MockInstallerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInstaller) EXPECT() *MockInstallerMockRecorder {
	return m.recorder
}

// InstallElemental mocks base method.
func (m *MockInstaller) InstallElemental(arg0 v1beta1.Config) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstallElemental", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// InstallElemental indicates an expected call of InstallElemental.
func (mr *MockInstallerMockRecorder) InstallElemental(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstallElemental", reflect.TypeOf((*MockInstaller)(nil).InstallElemental), arg0)
}

// IsSystemInstalled mocks base method.
func (m *MockInstaller) IsSystemInstalled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsSystemInstalled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsSystemInstalled indicates an expected call of IsSystemInstalled.
func (mr *MockInstallerMockRecorder) IsSystemInstalled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsSystemInstalled", reflect.TypeOf((*MockInstaller)(nil).IsSystemInstalled))
}
