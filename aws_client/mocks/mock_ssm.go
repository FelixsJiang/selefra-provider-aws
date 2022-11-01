package mocks

import (
	context "context"
	reflect "reflect"

	ssm "github.com/aws/aws-sdk-go-v2/service/ssm"
	gomock "github.com/golang/mock/gomock"
)

type MockSSMClient struct {
	ctrl		*gomock.Controller
	recorder	*MockSSMClientMockRecorder
}

type MockSSMClientMockRecorder struct {
	mock *MockSSMClient
}

func NewMockSSMClient(ctrl *gomock.Controller) *MockSSMClient {
	mock := &MockSSMClient{ctrl: ctrl}
	mock.recorder = &MockSSMClientMockRecorder{mock}
	return mock
}

func (m *MockSSMClient) EXPECT() *MockSSMClientMockRecorder {
	return m.recorder
}

func (m *MockSSMClient) DescribeDocument(arg0 context.Context, arg1 *ssm.DescribeDocumentInput, arg2 ...func(*ssm.Options)) (*ssm.DescribeDocumentOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeDocument", varargs...)
	ret0, _ := ret[0].(*ssm.DescribeDocumentOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSSMClientMockRecorder) DescribeDocument(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDocument", reflect.TypeOf((*MockSSMClient)(nil).DescribeDocument), varargs...)
}

func (m *MockSSMClient) DescribeDocumentPermission(arg0 context.Context, arg1 *ssm.DescribeDocumentPermissionInput, arg2 ...func(*ssm.Options)) (*ssm.DescribeDocumentPermissionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeDocumentPermission", varargs...)
	ret0, _ := ret[0].(*ssm.DescribeDocumentPermissionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSSMClientMockRecorder) DescribeDocumentPermission(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDocumentPermission", reflect.TypeOf((*MockSSMClient)(nil).DescribeDocumentPermission), varargs...)
}

func (m *MockSSMClient) DescribeInstanceInformation(arg0 context.Context, arg1 *ssm.DescribeInstanceInformationInput, arg2 ...func(*ssm.Options)) (*ssm.DescribeInstanceInformationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeInstanceInformation", varargs...)
	ret0, _ := ret[0].(*ssm.DescribeInstanceInformationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSSMClientMockRecorder) DescribeInstanceInformation(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeInstanceInformation", reflect.TypeOf((*MockSSMClient)(nil).DescribeInstanceInformation), varargs...)
}

func (m *MockSSMClient) DescribeParameters(arg0 context.Context, arg1 *ssm.DescribeParametersInput, arg2 ...func(*ssm.Options)) (*ssm.DescribeParametersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeParameters", varargs...)
	ret0, _ := ret[0].(*ssm.DescribeParametersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSSMClientMockRecorder) DescribeParameters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeParameters", reflect.TypeOf((*MockSSMClient)(nil).DescribeParameters), varargs...)
}

func (m *MockSSMClient) ListComplianceItems(arg0 context.Context, arg1 *ssm.ListComplianceItemsInput, arg2 ...func(*ssm.Options)) (*ssm.ListComplianceItemsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListComplianceItems", varargs...)
	ret0, _ := ret[0].(*ssm.ListComplianceItemsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSSMClientMockRecorder) ListComplianceItems(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListComplianceItems", reflect.TypeOf((*MockSSMClient)(nil).ListComplianceItems), varargs...)
}

func (m *MockSSMClient) ListDocuments(arg0 context.Context, arg1 *ssm.ListDocumentsInput, arg2 ...func(*ssm.Options)) (*ssm.ListDocumentsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListDocuments", varargs...)
	ret0, _ := ret[0].(*ssm.ListDocumentsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSSMClientMockRecorder) ListDocuments(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDocuments", reflect.TypeOf((*MockSSMClient)(nil).ListDocuments), varargs...)
}
