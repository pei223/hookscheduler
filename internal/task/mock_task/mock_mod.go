// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/pei223/hook-scheduler/internal/task (interfaces: TaskMod)

// Package mock_task is a generated GoMock package.
package mock_task

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
	models "github.com/pei223/hook-scheduler/internal/models"
	task "github.com/pei223/hook-scheduler/internal/task"
)

// MockTaskMod is a mock of TaskMod interface.
type MockTaskMod struct {
	ctrl     *gomock.Controller
	recorder *MockTaskModMockRecorder
}

// MockTaskModMockRecorder is the mock recorder for MockTaskMod.
type MockTaskModMockRecorder struct {
	mock *MockTaskMod
}

// NewMockTaskMod creates a new mock instance.
func NewMockTaskMod(ctrl *gomock.Controller) *MockTaskMod {
	mock := &MockTaskMod{ctrl: ctrl}
	mock.recorder = &MockTaskModMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTaskMod) EXPECT() *MockTaskModMockRecorder {
	return m.recorder
}

// CreateTask mocks base method.
func (m *MockTaskMod) CreateTask(arg0 context.Context, arg1 *task.TaskCreateParams) (*models.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTask", arg0, arg1)
	ret0, _ := ret[0].(*models.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTask indicates an expected call of CreateTask.
func (mr *MockTaskModMockRecorder) CreateTask(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTask", reflect.TypeOf((*MockTaskMod)(nil).CreateTask), arg0, arg1)
}

// GetTask mocks base method.
func (m *MockTaskMod) GetTask(arg0 context.Context, arg1 uuid.UUID) (*models.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTask", arg0, arg1)
	ret0, _ := ret[0].(*models.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTask indicates an expected call of GetTask.
func (mr *MockTaskModMockRecorder) GetTask(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTask", reflect.TypeOf((*MockTaskMod)(nil).GetTask), arg0, arg1)
}
