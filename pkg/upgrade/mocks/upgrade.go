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
// Source: github.com/rancher/elemental-operator/pkg/upgrade (interfaces: Upgrader)
//
// Generated by this command:
//
//	mockgen-v0.4.0 -copyright_file=scripts/boilerplate.go.txt -destination=pkg/upgrade/mocks/upgrade.go -package=mocks github.com/rancher/elemental-operator/pkg/upgrade Upgrader
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	upgrade "github.com/rancher/elemental-operator/pkg/upgrade"
	gomock "go.uber.org/mock/gomock"
)

// MockUpgrader is a mock of Upgrader interface.
type MockUpgrader struct {
	ctrl     *gomock.Controller
	recorder *MockUpgraderMockRecorder
}

// MockUpgraderMockRecorder is the mock recorder for MockUpgrader.
type MockUpgraderMockRecorder struct {
	mock *MockUpgrader
}

// NewMockUpgrader creates a new mock instance.
func NewMockUpgrader(ctrl *gomock.Controller) *MockUpgrader {
	mock := &MockUpgrader{ctrl: ctrl}
	mock.recorder = &MockUpgraderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUpgrader) EXPECT() *MockUpgraderMockRecorder {
	return m.recorder
}

// UpgradeElemental mocks base method.
func (m *MockUpgrader) UpgradeElemental(arg0 upgrade.Environment) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpgradeElemental", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpgradeElemental indicates an expected call of UpgradeElemental.
func (mr *MockUpgraderMockRecorder) UpgradeElemental(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpgradeElemental", reflect.TypeOf((*MockUpgrader)(nil).UpgradeElemental), arg0)
}
