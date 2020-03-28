// The MIT License (MIT)
//
// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: taskPriorityAssigner.go

// Package history is a generated GoMock package.
package history

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MocktaskPriorityAssigner is a mock of taskPriorityAssigner interface.
type MocktaskPriorityAssigner struct {
	ctrl     *gomock.Controller
	recorder *MocktaskPriorityAssignerMockRecorder
}

// MocktaskPriorityAssignerMockRecorder is the mock recorder for MocktaskPriorityAssigner.
type MocktaskPriorityAssignerMockRecorder struct {
	mock *MocktaskPriorityAssigner
}

// NewMocktaskPriorityAssigner creates a new mock instance.
func NewMocktaskPriorityAssigner(ctrl *gomock.Controller) *MocktaskPriorityAssigner {
	mock := &MocktaskPriorityAssigner{ctrl: ctrl}
	mock.recorder = &MocktaskPriorityAssignerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MocktaskPriorityAssigner) EXPECT() *MocktaskPriorityAssignerMockRecorder {
	return m.recorder
}

// Assign mocks base method.
func (m *MocktaskPriorityAssigner) Assign(arg0 queueTask) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Assign", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Assign indicates an expected call of Assign.
func (mr *MocktaskPriorityAssignerMockRecorder) Assign(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Assign", reflect.TypeOf((*MocktaskPriorityAssigner)(nil).Assign), arg0)
}