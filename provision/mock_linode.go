// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/provision/linode.go

// Package mock is a generated GoMock package.
package provision

import (
	gomock "github.com/golang/mock/gomock"
	linodego "github.com/linode/linodego"
	reflect "reflect"
)

// mockLinodeInterface is a mock of LinodeInterface interface
type mockLinodeInterface struct {
	ctrl     *gomock.Controller
	recorder *mockLinodeInterfaceMockRecorder
}

// mockLinodeInterfaceMockRecorder is the mock recorder for mockLinodeInterface
type mockLinodeInterfaceMockRecorder struct {
	mock *mockLinodeInterface
}

// newMockLinodeInterface creates a new mock instance
func newMockLinodeInterface(ctrl *gomock.Controller) *mockLinodeInterface {
	mock := &mockLinodeInterface{ctrl: ctrl}
	mock.recorder = &mockLinodeInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *mockLinodeInterface) EXPECT() *mockLinodeInterfaceMockRecorder {
	return m.recorder
}

// CreateStackscript mocks base method
func (m *mockLinodeInterface) CreateStackscript(createOpts linodego.StackscriptCreateOptions) (*linodego.Stackscript, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateStackscript", createOpts)
	ret0, _ := ret[0].(*linodego.Stackscript)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateStackscript indicates an expected call of CreateStackscript
func (mr *mockLinodeInterfaceMockRecorder) CreateStackscript(createOpts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateStackscript", reflect.TypeOf((*mockLinodeInterface)(nil).CreateStackscript), createOpts)
}

// CreateInstance mocks base method
func (m *mockLinodeInterface) CreateInstance(instance linodego.InstanceCreateOptions) (*linodego.Instance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateInstance", instance)
	ret0, _ := ret[0].(*linodego.Instance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateInstance indicates an expected call of CreateInstance
func (mr *mockLinodeInterfaceMockRecorder) CreateInstance(instance interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateInstance", reflect.TypeOf((*mockLinodeInterface)(nil).CreateInstance), instance)
}

// GetInstance mocks base method
func (m *mockLinodeInterface) GetInstance(linodeID int) (*linodego.Instance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInstance", linodeID)
	ret0, _ := ret[0].(*linodego.Instance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInstance indicates an expected call of GetInstance
func (mr *mockLinodeInterfaceMockRecorder) GetInstance(linodeID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInstance", reflect.TypeOf((*mockLinodeInterface)(nil).GetInstance), linodeID)
}

// DeleteInstance mocks base method
func (m *mockLinodeInterface) DeleteInstance(linodeID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteInstance", linodeID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteInstance indicates an expected call of DeleteInstance
func (mr *mockLinodeInterfaceMockRecorder) DeleteInstance(linodeID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteInstance", reflect.TypeOf((*mockLinodeInterface)(nil).DeleteInstance), linodeID)
}

// DeleteStackscript mocks base method
func (m *mockLinodeInterface) DeleteStackscript(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteStackscript", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteStackscript indicates an expected call of DeleteStackscript
func (mr *mockLinodeInterfaceMockRecorder) DeleteStackscript(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStackscript", reflect.TypeOf((*mockLinodeInterface)(nil).DeleteStackscript), id)
}