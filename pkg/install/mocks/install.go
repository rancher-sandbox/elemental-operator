// /*
// Copyright © 2022 - 2024 SUSE LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// */
//
//

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/rancher/elemental-operator/pkg/install (interfaces: Installer)
//
// Generated by this command:
//
//	mockgen -copyright_file=scripts/boilerplate.go.txt -destination=pkg/install/mocks/install.go -package=mocks github.com/rancher/elemental-operator/pkg/install Installer
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	v1beta1 "github.com/rancher/elemental-operator/api/v1beta1"
	register "github.com/rancher/elemental-operator/pkg/register"
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
func (m *MockInstaller) InstallElemental(arg0 v1beta1.Config, arg1 register.State) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstallElemental", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// InstallElemental indicates an expected call of InstallElemental.
func (mr *MockInstallerMockRecorder) InstallElemental(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstallElemental", reflect.TypeOf((*MockInstaller)(nil).InstallElemental), arg0, arg1)
}

// ResetElemental mocks base method.
func (m *MockInstaller) ResetElemental(arg0 v1beta1.Config, arg1 register.State) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResetElemental", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResetElemental indicates an expected call of ResetElemental.
func (mr *MockInstallerMockRecorder) ResetElemental(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetElemental", reflect.TypeOf((*MockInstaller)(nil).ResetElemental), arg0, arg1)
}

// ResetNetwork mocks base method.
func (m *MockInstaller) ResetNetwork() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResetNetwork")
	ret0, _ := ret[0].(error)
	return ret0
}

// ResetNetwork indicates an expected call of ResetNetwork.
func (mr *MockInstallerMockRecorder) ResetNetwork() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetNetwork", reflect.TypeOf((*MockInstaller)(nil).ResetNetwork))
}

// WriteLocalSystemAgentConfig mocks base method.
func (m *MockInstaller) WriteLocalSystemAgentConfig(arg0 v1beta1.Elemental) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteLocalSystemAgentConfig", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteLocalSystemAgentConfig indicates an expected call of WriteLocalSystemAgentConfig.
func (mr *MockInstallerMockRecorder) WriteLocalSystemAgentConfig(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteLocalSystemAgentConfig", reflect.TypeOf((*MockInstaller)(nil).WriteLocalSystemAgentConfig), arg0)
}
