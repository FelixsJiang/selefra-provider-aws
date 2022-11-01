package mocks

import (
	context "context"
	reflect "reflect"

	databasemigrationservice "github.com/aws/aws-sdk-go-v2/service/databasemigrationservice"
	gomock "github.com/golang/mock/gomock"
)

type MockDatabasemigrationserviceClient struct {
	ctrl		*gomock.Controller
	recorder	*MockDatabasemigrationserviceClientMockRecorder
}

type MockDatabasemigrationserviceClientMockRecorder struct {
	mock *MockDatabasemigrationserviceClient
}

func NewMockDatabasemigrationserviceClient(ctrl *gomock.Controller) *MockDatabasemigrationserviceClient {
	mock := &MockDatabasemigrationserviceClient{ctrl: ctrl}
	mock.recorder = &MockDatabasemigrationserviceClientMockRecorder{mock}
	return mock
}

func (m *MockDatabasemigrationserviceClient) EXPECT() *MockDatabasemigrationserviceClientMockRecorder {
	return m.recorder
}

func (m *MockDatabasemigrationserviceClient) DescribeReplicationInstances(arg0 context.Context, arg1 *databasemigrationservice.DescribeReplicationInstancesInput, arg2 ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeReplicationInstancesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeReplicationInstances", varargs...)
	ret0, _ := ret[0].(*databasemigrationservice.DescribeReplicationInstancesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDatabasemigrationserviceClientMockRecorder) DescribeReplicationInstances(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeReplicationInstances", reflect.TypeOf((*MockDatabasemigrationserviceClient)(nil).DescribeReplicationInstances), varargs...)
}

func (m *MockDatabasemigrationserviceClient) ListTagsForResource(arg0 context.Context, arg1 *databasemigrationservice.ListTagsForResourceInput, arg2 ...func(*databasemigrationservice.Options)) (*databasemigrationservice.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTagsForResource", varargs...)
	ret0, _ := ret[0].(*databasemigrationservice.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDatabasemigrationserviceClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockDatabasemigrationserviceClient)(nil).ListTagsForResource), varargs...)
}
