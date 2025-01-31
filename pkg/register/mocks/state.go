// /*
// Copyright © 2022 - 2025 SUSE LLC
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
// Source: github.com/rancher/elemental-operator/pkg/register (interfaces: StateHandler)
//
// Generated by this command:
//
//	mockgen-v0.4.0 -copyright_file=scripts/boilerplate.go.txt -destination=pkg/register/mocks/state.go -package=mocks github.com/rancher/elemental-operator/pkg/register StateHandler
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	register "github.com/rancher/elemental-operator/pkg/register"
	gomock "go.uber.org/mock/gomock"
)

// MockStateHandler is a mock of StateHandler interface.
type MockStateHandler struct {
	ctrl     *gomock.Controller
	recorder *MockStateHandlerMockRecorder
}

// MockStateHandlerMockRecorder is the mock recorder for MockStateHandler.
type MockStateHandlerMockRecorder struct {
	mock *MockStateHandler
}

// NewMockStateHandler creates a new mock instance.
func NewMockStateHandler(ctrl *gomock.Controller) *MockStateHandler {
	mock := &MockStateHandler{ctrl: ctrl}
	mock.recorder = &MockStateHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStateHandler) EXPECT() *MockStateHandlerMockRecorder {
	return m.recorder
}

// Init mocks base method.
func (m *MockStateHandler) Init(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init.
func (mr *MockStateHandlerMockRecorder) Init(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockStateHandler)(nil).Init), arg0)
}

// Load mocks base method.
func (m *MockStateHandler) Load() (register.State, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Load")
	ret0, _ := ret[0].(register.State)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Load indicates an expected call of Load.
func (mr *MockStateHandlerMockRecorder) Load() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Load", reflect.TypeOf((*MockStateHandler)(nil).Load))
}

// Save mocks base method.
func (m *MockStateHandler) Save(arg0 register.State) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockStateHandlerMockRecorder) Save(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockStateHandler)(nil).Save), arg0)
}
