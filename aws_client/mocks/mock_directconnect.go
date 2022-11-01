package mocks

import (
	context "context"
	reflect "reflect"

	directconnect "github.com/aws/aws-sdk-go-v2/service/directconnect"
	gomock "github.com/golang/mock/gomock"
)

type MockDirectconnectClient struct {
	ctrl		*gomock.Controller
	recorder	*MockDirectconnectClientMockRecorder
}

type MockDirectconnectClientMockRecorder struct {
	mock *MockDirectconnectClient
}

func NewMockDirectconnectClient(ctrl *gomock.Controller) *MockDirectconnectClient {
	mock := &MockDirectconnectClient{ctrl: ctrl}
	mock.recorder = &MockDirectconnectClientMockRecorder{mock}
	return mock
}

func (m *MockDirectconnectClient) EXPECT() *MockDirectconnectClientMockRecorder {
	return m.recorder
}

func (m *MockDirectconnectClient) DescribeConnections(arg0 context.Context, arg1 *directconnect.DescribeConnectionsInput, arg2 ...func(*directconnect.Options)) (*directconnect.DescribeConnectionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeConnections", varargs...)
	ret0, _ := ret[0].(*directconnect.DescribeConnectionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDirectconnectClientMockRecorder) DescribeConnections(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeConnections", reflect.TypeOf((*MockDirectconnectClient)(nil).DescribeConnections), varargs...)
}

func (m *MockDirectconnectClient) DescribeDirectConnectGatewayAssociations(arg0 context.Context, arg1 *directconnect.DescribeDirectConnectGatewayAssociationsInput, arg2 ...func(*directconnect.Options)) (*directconnect.DescribeDirectConnectGatewayAssociationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeDirectConnectGatewayAssociations", varargs...)
	ret0, _ := ret[0].(*directconnect.DescribeDirectConnectGatewayAssociationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDirectconnectClientMockRecorder) DescribeDirectConnectGatewayAssociations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDirectConnectGatewayAssociations", reflect.TypeOf((*MockDirectconnectClient)(nil).DescribeDirectConnectGatewayAssociations), varargs...)
}

func (m *MockDirectconnectClient) DescribeDirectConnectGatewayAttachments(arg0 context.Context, arg1 *directconnect.DescribeDirectConnectGatewayAttachmentsInput, arg2 ...func(*directconnect.Options)) (*directconnect.DescribeDirectConnectGatewayAttachmentsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeDirectConnectGatewayAttachments", varargs...)
	ret0, _ := ret[0].(*directconnect.DescribeDirectConnectGatewayAttachmentsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDirectconnectClientMockRecorder) DescribeDirectConnectGatewayAttachments(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDirectConnectGatewayAttachments", reflect.TypeOf((*MockDirectconnectClient)(nil).DescribeDirectConnectGatewayAttachments), varargs...)
}

func (m *MockDirectconnectClient) DescribeDirectConnectGateways(arg0 context.Context, arg1 *directconnect.DescribeDirectConnectGatewaysInput, arg2 ...func(*directconnect.Options)) (*directconnect.DescribeDirectConnectGatewaysOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeDirectConnectGateways", varargs...)
	ret0, _ := ret[0].(*directconnect.DescribeDirectConnectGatewaysOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDirectconnectClientMockRecorder) DescribeDirectConnectGateways(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDirectConnectGateways", reflect.TypeOf((*MockDirectconnectClient)(nil).DescribeDirectConnectGateways), varargs...)
}

func (m *MockDirectconnectClient) DescribeLags(arg0 context.Context, arg1 *directconnect.DescribeLagsInput, arg2 ...func(*directconnect.Options)) (*directconnect.DescribeLagsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeLags", varargs...)
	ret0, _ := ret[0].(*directconnect.DescribeLagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDirectconnectClientMockRecorder) DescribeLags(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeLags", reflect.TypeOf((*MockDirectconnectClient)(nil).DescribeLags), varargs...)
}

func (m *MockDirectconnectClient) DescribeVirtualGateways(arg0 context.Context, arg1 *directconnect.DescribeVirtualGatewaysInput, arg2 ...func(*directconnect.Options)) (*directconnect.DescribeVirtualGatewaysOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeVirtualGateways", varargs...)
	ret0, _ := ret[0].(*directconnect.DescribeVirtualGatewaysOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDirectconnectClientMockRecorder) DescribeVirtualGateways(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeVirtualGateways", reflect.TypeOf((*MockDirectconnectClient)(nil).DescribeVirtualGateways), varargs...)
}

func (m *MockDirectconnectClient) DescribeVirtualInterfaces(arg0 context.Context, arg1 *directconnect.DescribeVirtualInterfacesInput, arg2 ...func(*directconnect.Options)) (*directconnect.DescribeVirtualInterfacesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeVirtualInterfaces", varargs...)
	ret0, _ := ret[0].(*directconnect.DescribeVirtualInterfacesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDirectconnectClientMockRecorder) DescribeVirtualInterfaces(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeVirtualInterfaces", reflect.TypeOf((*MockDirectconnectClient)(nil).DescribeVirtualInterfaces), varargs...)
}
