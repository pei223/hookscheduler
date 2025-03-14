// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/pei223/hook-scheduler/internal/usecase (interfaces: HookServiceIF)

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
	hook "github.com/pei223/hook-scheduler/internal/domain/hook"
	models "github.com/pei223/hook-scheduler/internal/models"
)

// MockHookServiceIF is a mock of HookServiceIF interface.
type MockHookServiceIF struct {
	ctrl     *gomock.Controller
	recorder *MockHookServiceIFMockRecorder
}

// MockHookServiceIFMockRecorder is the mock recorder for MockHookServiceIF.
type MockHookServiceIFMockRecorder struct {
	mock *MockHookServiceIF
}

// NewMockHookServiceIF creates a new mock instance.
func NewMockHookServiceIF(ctrl *gomock.Controller) *MockHookServiceIF {
	mock := &MockHookServiceIF{ctrl: ctrl}
	mock.recorder = &MockHookServiceIFMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHookServiceIF) EXPECT() *MockHookServiceIFMockRecorder {
	return m.recorder
}

// CreateHook mocks base method.
func (m *MockHookServiceIF) CreateHook(arg0 context.Context, arg1 *hook.HookCreateParams) (*models.Hook, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateHook", arg0, arg1)
	ret0, _ := ret[0].(*models.Hook)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateHook indicates an expected call of CreateHook.
func (mr *MockHookServiceIFMockRecorder) CreateHook(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateHook", reflect.TypeOf((*MockHookServiceIF)(nil).CreateHook), arg0, arg1)
}

// DeleteHook mocks base method.
func (m *MockHookServiceIF) DeleteHook(arg0 context.Context, arg1 uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteHook", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteHook indicates an expected call of DeleteHook.
func (mr *MockHookServiceIFMockRecorder) DeleteHook(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteHook", reflect.TypeOf((*MockHookServiceIF)(nil).DeleteHook), arg0, arg1)
}

// GetAllHooks mocks base method.
func (m *MockHookServiceIF) GetAllHooks(arg0 context.Context, arg1, arg2 int) (models.HookSlice, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllHooks", arg0, arg1, arg2)
	ret0, _ := ret[0].(models.HookSlice)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetAllHooks indicates an expected call of GetAllHooks.
func (mr *MockHookServiceIFMockRecorder) GetAllHooks(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllHooks", reflect.TypeOf((*MockHookServiceIF)(nil).GetAllHooks), arg0, arg1, arg2)
}

// GetHook mocks base method.
func (m *MockHookServiceIF) GetHook(arg0 context.Context, arg1 uuid.UUID) (*models.Hook, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHook", arg0, arg1)
	ret0, _ := ret[0].(*models.Hook)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHook indicates an expected call of GetHook.
func (mr *MockHookServiceIFMockRecorder) GetHook(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHook", reflect.TypeOf((*MockHookServiceIF)(nil).GetHook), arg0, arg1)
}
