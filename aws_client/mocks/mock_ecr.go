package mocks

import (
	context "context"
	reflect "reflect"

	ecr "github.com/aws/aws-sdk-go-v2/service/ecr"
	gomock "github.com/golang/mock/gomock"
)

type MockEcrClient struct {
	ctrl		*gomock.Controller
	recorder	*MockEcrClientMockRecorder
}

type MockEcrClientMockRecorder struct {
	mock *MockEcrClient
}

func NewMockEcrClient(ctrl *gomock.Controller) *MockEcrClient {
	mock := &MockEcrClient{ctrl: ctrl}
	mock.recorder = &MockEcrClientMockRecorder{mock}
	return mock
}

func (m *MockEcrClient) EXPECT() *MockEcrClientMockRecorder {
	return m.recorder
}

func (m *MockEcrClient) DescribeImages(arg0 context.Context, arg1 *ecr.DescribeImagesInput, arg2 ...func(*ecr.Options)) (*ecr.DescribeImagesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeImages", varargs...)
	ret0, _ := ret[0].(*ecr.DescribeImagesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEcrClientMockRecorder) DescribeImages(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeImages", reflect.TypeOf((*MockEcrClient)(nil).DescribeImages), varargs...)
}

func (m *MockEcrClient) DescribeRegistry(arg0 context.Context, arg1 *ecr.DescribeRegistryInput, arg2 ...func(*ecr.Options)) (*ecr.DescribeRegistryOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeRegistry", varargs...)
	ret0, _ := ret[0].(*ecr.DescribeRegistryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEcrClientMockRecorder) DescribeRegistry(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeRegistry", reflect.TypeOf((*MockEcrClient)(nil).DescribeRegistry), varargs...)
}

func (m *MockEcrClient) DescribeRepositories(arg0 context.Context, arg1 *ecr.DescribeRepositoriesInput, arg2 ...func(*ecr.Options)) (*ecr.DescribeRepositoriesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeRepositories", varargs...)
	ret0, _ := ret[0].(*ecr.DescribeRepositoriesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEcrClientMockRecorder) DescribeRepositories(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeRepositories", reflect.TypeOf((*MockEcrClient)(nil).DescribeRepositories), varargs...)
}

func (m *MockEcrClient) GetRegistryPolicy(arg0 context.Context, arg1 *ecr.GetRegistryPolicyInput, arg2 ...func(*ecr.Options)) (*ecr.GetRegistryPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRegistryPolicy", varargs...)
	ret0, _ := ret[0].(*ecr.GetRegistryPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEcrClientMockRecorder) GetRegistryPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRegistryPolicy", reflect.TypeOf((*MockEcrClient)(nil).GetRegistryPolicy), varargs...)
}

func (m *MockEcrClient) ListTagsForResource(arg0 context.Context, arg1 *ecr.ListTagsForResourceInput, arg2 ...func(*ecr.Options)) (*ecr.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTagsForResource", varargs...)
	ret0, _ := ret[0].(*ecr.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEcrClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockEcrClient)(nil).ListTagsForResource), varargs...)
}
