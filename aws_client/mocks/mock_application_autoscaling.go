package mocks

import (
	context "context"
	reflect "reflect"

	applicationautoscaling "github.com/aws/aws-sdk-go-v2/service/applicationautoscaling"
	gomock "github.com/golang/mock/gomock"
)

type MockApplicationAutoscalingClient struct {
	ctrl		*gomock.Controller
	recorder	*MockApplicationAutoscalingClientMockRecorder
}

type MockApplicationAutoscalingClientMockRecorder struct {
	mock *MockApplicationAutoscalingClient
}

func NewMockApplicationAutoscalingClient(ctrl *gomock.Controller) *MockApplicationAutoscalingClient {
	mock := &MockApplicationAutoscalingClient{ctrl: ctrl}
	mock.recorder = &MockApplicationAutoscalingClientMockRecorder{mock}
	return mock
}

func (m *MockApplicationAutoscalingClient) EXPECT() *MockApplicationAutoscalingClientMockRecorder {
	return m.recorder
}

func (m *MockApplicationAutoscalingClient) DescribeScalingPolicies(arg0 context.Context, arg1 *applicationautoscaling.DescribeScalingPoliciesInput, arg2 ...func(*applicationautoscaling.Options)) (*applicationautoscaling.DescribeScalingPoliciesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeScalingPolicies", varargs...)
	ret0, _ := ret[0].(*applicationautoscaling.DescribeScalingPoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApplicationAutoscalingClientMockRecorder) DescribeScalingPolicies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeScalingPolicies", reflect.TypeOf((*MockApplicationAutoscalingClient)(nil).DescribeScalingPolicies), varargs...)
}
