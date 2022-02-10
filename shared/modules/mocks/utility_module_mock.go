// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/shared/modules/utility_module.go

// Package mock_modules is a generated GoMock package.
package mock_modules

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	context "pocket/shared/context"
	modules "pocket/shared/modules"
	typespb "pocket/consensus/pkg/types/typespb"
)

// MockUtilityModule is a mock of UtilityModule interface.
type MockUtilityModule struct {
	ctrl     *gomock.Controller
	recorder *MockUtilityModuleMockRecorder
}

// MockUtilityModuleMockRecorder is the mock recorder for MockUtilityModule.
type MockUtilityModuleMockRecorder struct {
	mock *MockUtilityModule
}

// NewMockUtilityModule creates a new mock instance.
func NewMockUtilityModule(ctrl *gomock.Controller) *MockUtilityModule {
	mock := &MockUtilityModule{ctrl: ctrl}
	mock.recorder = &MockUtilityModuleMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUtilityModule) EXPECT() *MockUtilityModuleMockRecorder {
	return m.recorder
}

// BeginBlock mocks base method.
func (m *MockUtilityModule) BeginBlock(arg0 *context.PocketContext) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeginBlock", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// BeginBlock indicates an expected call of BeginBlock.
func (mr *MockUtilityModuleMockRecorder) BeginBlock(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginBlock", reflect.TypeOf((*MockUtilityModule)(nil).BeginBlock), arg0)
}

// DeliverTx mocks base method.
func (m *MockUtilityModule) DeliverTx(arg0 *context.PocketContext, arg1 *typespb.Transaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeliverTx", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeliverTx indicates an expected call of DeliverTx.
func (mr *MockUtilityModuleMockRecorder) DeliverTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeliverTx", reflect.TypeOf((*MockUtilityModule)(nil).DeliverTx), arg0, arg1)
}

// EndBlock mocks base method.
func (m *MockUtilityModule) EndBlock(arg0 *context.PocketContext) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EndBlock", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// EndBlock indicates an expected call of EndBlock.
func (mr *MockUtilityModuleMockRecorder) EndBlock(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EndBlock", reflect.TypeOf((*MockUtilityModule)(nil).EndBlock), arg0)
}

// GetPocketBusMod mocks base method.
func (m *MockUtilityModule) GetPocketBusMod() modules.PocketBusModule {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPocketBusMod")
	ret0, _ := ret[0].(modules.PocketBusModule)
	return ret0
}

// GetPocketBusMod indicates an expected call of GetPocketBusMod.
func (mr *MockUtilityModuleMockRecorder) GetPocketBusMod() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPocketBusMod", reflect.TypeOf((*MockUtilityModule)(nil).GetPocketBusMod))
}

// HandleEvidence mocks base method.
func (m *MockUtilityModule) HandleEvidence(arg0 *context.PocketContext, arg1 *typespb.Evidence) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleEvidence", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleEvidence indicates an expected call of HandleEvidence.
func (mr *MockUtilityModuleMockRecorder) HandleEvidence(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleEvidence", reflect.TypeOf((*MockUtilityModule)(nil).HandleEvidence), arg0, arg1)
}

// HandleTransaction mocks base method.
func (m *MockUtilityModule) HandleTransaction(arg0 *context.PocketContext, arg1 *typespb.Transaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleTransaction", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleTransaction indicates an expected call of HandleTransaction.
func (mr *MockUtilityModuleMockRecorder) HandleTransaction(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleTransaction", reflect.TypeOf((*MockUtilityModule)(nil).HandleTransaction), arg0, arg1)
}

// ReapMempool mocks base method.
func (m *MockUtilityModule) ReapMempool(arg0 *context.PocketContext) ([]*typespb.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReapMempool", arg0)
	ret0, _ := ret[0].([]*typespb.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReapMempool indicates an expected call of ReapMempool.
func (mr *MockUtilityModuleMockRecorder) ReapMempool(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReapMempool", reflect.TypeOf((*MockUtilityModule)(nil).ReapMempool), arg0)
}

// SetPocketBusMod mocks base method.
func (m *MockUtilityModule) SetPocketBusMod(arg0 modules.PocketBusModule) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetPocketBusMod", arg0)
}

// SetPocketBusMod indicates an expected call of SetPocketBusMod.
func (mr *MockUtilityModuleMockRecorder) SetPocketBusMod(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPocketBusMod", reflect.TypeOf((*MockUtilityModule)(nil).SetPocketBusMod), arg0)
}

// Start mocks base method.
func (m *MockUtilityModule) Start(arg0 *context.PocketContext) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockUtilityModuleMockRecorder) Start(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockUtilityModule)(nil).Start), arg0)
}

// Stop mocks base method.
func (m *MockUtilityModule) Stop(arg0 *context.PocketContext) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop.
func (mr *MockUtilityModuleMockRecorder) Stop(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockUtilityModule)(nil).Stop), arg0)
}