package mocks

import (
	context "context"
	reflect "reflect"

	emr "github.com/aws/aws-sdk-go-v2/service/emr"
	gomock "github.com/golang/mock/gomock"
)

type MockEmrClient struct {
	ctrl		*gomock.Controller
	recorder	*MockEmrClientMockRecorder
}

type MockEmrClientMockRecorder struct {
	mock *MockEmrClient
}

func NewMockEmrClient(ctrl *gomock.Controller) *MockEmrClient {
	mock := &MockEmrClient{ctrl: ctrl}
	mock.recorder = &MockEmrClientMockRecorder{mock}
	return mock
}

func (m *MockEmrClient) EXPECT() *MockEmrClientMockRecorder {
	return m.recorder
}

func (m *MockEmrClient) DescribeCluster(arg0 context.Context, arg1 *emr.DescribeClusterInput, arg2 ...func(*emr.Options)) (*emr.DescribeClusterOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeCluster", varargs...)
	ret0, _ := ret[0].(*emr.DescribeClusterOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEmrClientMockRecorder) DescribeCluster(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCluster", reflect.TypeOf((*MockEmrClient)(nil).DescribeCluster), varargs...)
}

func (m *MockEmrClient) GetBlockPublicAccessConfiguration(arg0 context.Context, arg1 *emr.GetBlockPublicAccessConfigurationInput, arg2 ...func(*emr.Options)) (*emr.GetBlockPublicAccessConfigurationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBlockPublicAccessConfiguration", varargs...)
	ret0, _ := ret[0].(*emr.GetBlockPublicAccessConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEmrClientMockRecorder) GetBlockPublicAccessConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockPublicAccessConfiguration", reflect.TypeOf((*MockEmrClient)(nil).GetBlockPublicAccessConfiguration), varargs...)
}

func (m *MockEmrClient) ListClusters(arg0 context.Context, arg1 *emr.ListClustersInput, arg2 ...func(*emr.Options)) (*emr.ListClustersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListClusters", varargs...)
	ret0, _ := ret[0].(*emr.ListClustersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEmrClientMockRecorder) ListClusters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListClusters", reflect.TypeOf((*MockEmrClient)(nil).ListClusters), varargs...)
}
