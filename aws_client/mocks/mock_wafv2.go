package mocks

import (
	context "context"
	reflect "reflect"

	wafv2 "github.com/aws/aws-sdk-go-v2/service/wafv2"
	gomock "github.com/golang/mock/gomock"
)

type MockWafV2Client struct {
	ctrl		*gomock.Controller
	recorder	*MockWafV2ClientMockRecorder
}

type MockWafV2ClientMockRecorder struct {
	mock *MockWafV2Client
}

func NewMockWafV2Client(ctrl *gomock.Controller) *MockWafV2Client {
	mock := &MockWafV2Client{ctrl: ctrl}
	mock.recorder = &MockWafV2ClientMockRecorder{mock}
	return mock
}

func (m *MockWafV2Client) EXPECT() *MockWafV2ClientMockRecorder {
	return m.recorder
}

func (m *MockWafV2Client) DescribeManagedRuleGroup(arg0 context.Context, arg1 *wafv2.DescribeManagedRuleGroupInput, arg2 ...func(*wafv2.Options)) (*wafv2.DescribeManagedRuleGroupOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeManagedRuleGroup", varargs...)
	ret0, _ := ret[0].(*wafv2.DescribeManagedRuleGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafV2ClientMockRecorder) DescribeManagedRuleGroup(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeManagedRuleGroup", reflect.TypeOf((*MockWafV2Client)(nil).DescribeManagedRuleGroup), varargs...)
}

func (m *MockWafV2Client) GetIPSet(arg0 context.Context, arg1 *wafv2.GetIPSetInput, arg2 ...func(*wafv2.Options)) (*wafv2.GetIPSetOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetIPSet", varargs...)
	ret0, _ := ret[0].(*wafv2.GetIPSetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafV2ClientMockRecorder) GetIPSet(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIPSet", reflect.TypeOf((*MockWafV2Client)(nil).GetIPSet), varargs...)
}

func (m *MockWafV2Client) GetLoggingConfiguration(arg0 context.Context, arg1 *wafv2.GetLoggingConfigurationInput, arg2 ...func(*wafv2.Options)) (*wafv2.GetLoggingConfigurationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetLoggingConfiguration", varargs...)
	ret0, _ := ret[0].(*wafv2.GetLoggingConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafV2ClientMockRecorder) GetLoggingConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLoggingConfiguration", reflect.TypeOf((*MockWafV2Client)(nil).GetLoggingConfiguration), varargs...)
}

func (m *MockWafV2Client) GetPermissionPolicy(arg0 context.Context, arg1 *wafv2.GetPermissionPolicyInput, arg2 ...func(*wafv2.Options)) (*wafv2.GetPermissionPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetPermissionPolicy", varargs...)
	ret0, _ := ret[0].(*wafv2.GetPermissionPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafV2ClientMockRecorder) GetPermissionPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPermissionPolicy", reflect.TypeOf((*MockWafV2Client)(nil).GetPermissionPolicy), varargs...)
}

func (m *MockWafV2Client) GetRegexPatternSet(arg0 context.Context, arg1 *wafv2.GetRegexPatternSetInput, arg2 ...func(*wafv2.Options)) (*wafv2.GetRegexPatternSetOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRegexPatternSet", varargs...)
	ret0, _ := ret[0].(*wafv2.GetRegexPatternSetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafV2ClientMockRecorder) GetRegexPatternSet(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRegexPatternSet", reflect.TypeOf((*MockWafV2Client)(nil).GetRegexPatternSet), varargs...)
}

func (m *MockWafV2Client) GetRuleGroup(arg0 context.Context, arg1 *wafv2.GetRuleGroupInput, arg2 ...func(*wafv2.Options)) (*wafv2.GetRuleGroupOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRuleGroup", varargs...)
	ret0, _ := ret[0].(*wafv2.GetRuleGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafV2ClientMockRecorder) GetRuleGroup(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRuleGroup", reflect.TypeOf((*MockWafV2Client)(nil).GetRuleGroup), varargs...)
}

func (m *MockWafV2Client) GetWebACL(arg0 context.Context, arg1 *wafv2.GetWebACLInput, arg2 ...func(*wafv2.Options)) (*wafv2.GetWebACLOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetWebACL", varargs...)
	ret0, _ := ret[0].(*wafv2.GetWebACLOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafV2ClientMockRecorder) GetWebACL(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWebACL", reflect.TypeOf((*MockWafV2Client)(nil).GetWebACL), varargs...)
}

func (m *MockWafV2Client) GetWebACLForResource(arg0 context.Context, arg1 *wafv2.GetWebACLForResourceInput, arg2 ...func(*wafv2.Options)) (*wafv2.GetWebACLForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetWebACLForResource", varargs...)
	ret0, _ := ret[0].(*wafv2.GetWebACLForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafV2ClientMockRecorder) GetWebACLForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWebACLForResource", reflect.TypeOf((*MockWafV2Client)(nil).GetWebACLForResource), varargs...)
}

func (m *MockWafV2Client) ListAvailableManagedRuleGroups(arg0 context.Context, arg1 *wafv2.ListAvailableManagedRuleGroupsInput, arg2 ...func(*wafv2.Options)) (*wafv2.ListAvailableManagedRuleGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAvailableManagedRuleGroups", varargs...)
	ret0, _ := ret[0].(*wafv2.ListAvailableManagedRuleGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafV2ClientMockRecorder) ListAvailableManagedRuleGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAvailableManagedRuleGroups", reflect.TypeOf((*MockWafV2Client)(nil).ListAvailableManagedRuleGroups), varargs...)
}

func (m *MockWafV2Client) ListIPSets(arg0 context.Context, arg1 *wafv2.ListIPSetsInput, arg2 ...func(*wafv2.Options)) (*wafv2.ListIPSetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListIPSets", varargs...)
	ret0, _ := ret[0].(*wafv2.ListIPSetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafV2ClientMockRecorder) ListIPSets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListIPSets", reflect.TypeOf((*MockWafV2Client)(nil).ListIPSets), varargs...)
}

func (m *MockWafV2Client) ListRegexPatternSets(arg0 context.Context, arg1 *wafv2.ListRegexPatternSetsInput, arg2 ...func(*wafv2.Options)) (*wafv2.ListRegexPatternSetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListRegexPatternSets", varargs...)
	ret0, _ := ret[0].(*wafv2.ListRegexPatternSetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafV2ClientMockRecorder) ListRegexPatternSets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRegexPatternSets", reflect.TypeOf((*MockWafV2Client)(nil).ListRegexPatternSets), varargs...)
}

func (m *MockWafV2Client) ListResourcesForWebACL(arg0 context.Context, arg1 *wafv2.ListResourcesForWebACLInput, arg2 ...func(*wafv2.Options)) (*wafv2.ListResourcesForWebACLOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListResourcesForWebACL", varargs...)
	ret0, _ := ret[0].(*wafv2.ListResourcesForWebACLOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafV2ClientMockRecorder) ListResourcesForWebACL(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListResourcesForWebACL", reflect.TypeOf((*MockWafV2Client)(nil).ListResourcesForWebACL), varargs...)
}

func (m *MockWafV2Client) ListRuleGroups(arg0 context.Context, arg1 *wafv2.ListRuleGroupsInput, arg2 ...func(*wafv2.Options)) (*wafv2.ListRuleGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListRuleGroups", varargs...)
	ret0, _ := ret[0].(*wafv2.ListRuleGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafV2ClientMockRecorder) ListRuleGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRuleGroups", reflect.TypeOf((*MockWafV2Client)(nil).ListRuleGroups), varargs...)
}

func (m *MockWafV2Client) ListTagsForResource(arg0 context.Context, arg1 *wafv2.ListTagsForResourceInput, arg2 ...func(*wafv2.Options)) (*wafv2.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTagsForResource", varargs...)
	ret0, _ := ret[0].(*wafv2.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafV2ClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockWafV2Client)(nil).ListTagsForResource), varargs...)
}

func (m *MockWafV2Client) ListWebACLs(arg0 context.Context, arg1 *wafv2.ListWebACLsInput, arg2 ...func(*wafv2.Options)) (*wafv2.ListWebACLsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListWebACLs", varargs...)
	ret0, _ := ret[0].(*wafv2.ListWebACLsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafV2ClientMockRecorder) ListWebACLs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListWebACLs", reflect.TypeOf((*MockWafV2Client)(nil).ListWebACLs), varargs...)
}
