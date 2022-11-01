package mocks

import (
	context "context"
	reflect "reflect"

	ec2 "github.com/aws/aws-sdk-go-v2/service/ec2"
	gomock "github.com/golang/mock/gomock"
)

type MockEc2Client struct {
	ctrl		*gomock.Controller
	recorder	*MockEc2ClientMockRecorder
}

type MockEc2ClientMockRecorder struct {
	mock *MockEc2Client
}

func NewMockEc2Client(ctrl *gomock.Controller) *MockEc2Client {
	mock := &MockEc2Client{ctrl: ctrl}
	mock.recorder = &MockEc2ClientMockRecorder{mock}
	return mock
}

func (m *MockEc2Client) EXPECT() *MockEc2ClientMockRecorder {
	return m.recorder
}

func (m *MockEc2Client) DescribeAddresses(arg0 context.Context, arg1 *ec2.DescribeAddressesInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeAddressesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeAddresses", varargs...)
	ret0, _ := ret[0].(*ec2.DescribeAddressesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeAddresses(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeAddresses", reflect.TypeOf((*MockEc2Client)(nil).DescribeAddresses), varargs...)
}

func (m *MockEc2Client) DescribeByoipCidrs(arg0 context.Context, arg1 *ec2.DescribeByoipCidrsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeByoipCidrsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeByoipCidrs", varargs...)
	ret0, _ := ret[0].(*ec2.DescribeByoipCidrsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeByoipCidrs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeByoipCidrs", reflect.TypeOf((*MockEc2Client)(nil).DescribeByoipCidrs), varargs...)
}

func (m *MockEc2Client) DescribeCustomerGateways(arg0 context.Context, arg1 *ec2.DescribeCustomerGatewaysInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeCustomerGatewaysOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeCustomerGateways", varargs...)
	ret0, _ := ret[0].(*ec2.DescribeCustomerGatewaysOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeCustomerGateways(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCustomerGateways", reflect.TypeOf((*MockEc2Client)(nil).DescribeCustomerGateways), varargs...)
}

func (m *MockEc2Client) DescribeEgressOnlyInternetGateways(arg0 context.Context, arg1 *ec2.DescribeEgressOnlyInternetGatewaysInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeEgressOnlyInternetGatewaysOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeEgressOnlyInternetGateways", varargs...)
	ret0, _ := ret[0].(*ec2.DescribeEgressOnlyInternetGatewaysOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeEgressOnlyInternetGateways(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeEgressOnlyInternetGateways", reflect.TypeOf((*MockEc2Client)(nil).DescribeEgressOnlyInternetGateways), varargs...)
}

func (m *MockEc2Client) DescribeFlowLogs(arg0 context.Context, arg1 *ec2.DescribeFlowLogsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeFlowLogsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeFlowLogs", varargs...)
	ret0, _ := ret[0].(*ec2.DescribeFlowLogsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeFlowLogs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeFlowLogs", reflect.TypeOf((*MockEc2Client)(nil).DescribeFlowLogs), varargs...)
}

func (m *MockEc2Client) DescribeHosts(arg0 context.Context, arg1 *ec2.DescribeHostsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeHostsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeHosts", varargs...)
	ret0, _ := ret[0].(*ec2.DescribeHostsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeHosts(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeHosts", reflect.TypeOf((*MockEc2Client)(nil).DescribeHosts), varargs...)
}

func (m *MockEc2Client) DescribeImageAttribute(arg0 context.Context, arg1 *ec2.DescribeImageAttributeInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeImageAttributeOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeImageAttribute", varargs...)
	ret0, _ := ret[0].(*ec2.DescribeImageAttributeOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeImageAttribute(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeImageAttribute", reflect.TypeOf((*MockEc2Client)(nil).DescribeImageAttribute), varargs...)
}

func (m *MockEc2Client) DescribeImages(arg0 context.Context, arg1 *ec2.DescribeImagesInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeImagesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeImages", varargs...)
	ret0, _ := ret[0].(*ec2.DescribeImagesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeImages(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeImages", reflect.TypeOf((*MockEc2Client)(nil).DescribeImages), varargs...)
}

func (m *MockEc2Client) DescribeInstanceStatus(arg0 context.Context, arg1 *ec2.DescribeInstanceStatusInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeInstanceStatusOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeInstanceStatus", varargs...)
	ret0, _ := ret[0].(*ec2.DescribeInstanceStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeInstanceStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeInstanceStatus", reflect.TypeOf((*MockEc2Client)(nil).DescribeInstanceStatus), varargs...)
}

func (m *MockEc2Client) DescribeInstanceTypes(arg0 context.Context, arg1 *ec2.DescribeInstanceTypesInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeInstanceTypesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeInstanceTypes", varargs...)
	ret0, _ := ret[0].(*ec2.DescribeInstanceTypesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeInstanceTypes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeInstanceTypes", reflect.TypeOf((*MockEc2Client)(nil).DescribeInstanceTypes), varargs...)
}

func (m *MockEc2Client) DescribeInstances(arg0 context.Context, arg1 *ec2.DescribeInstancesInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeInstancesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeInstances", varargs...)
	ret0, _ := ret[0].(*ec2.DescribeInstancesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeInstances(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeInstances", reflect.TypeOf((*MockEc2Client)(nil).DescribeInstances), varargs...)
}

func (m *MockEc2Client) DescribeInternetGateways(arg0 context.Context, arg1 *ec2.DescribeInternetGatewaysInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeInternetGatewaysOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeInternetGateways", varargs...)
	ret0, _ := ret[0].(*ec2.DescribeInternetGatewaysOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeInternetGateways(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeInternetGateways", reflect.TypeOf((*MockEc2Client)(nil).DescribeInternetGateways), varargs...)
}

func (m *MockEc2Client) DescribeKeyPairs(arg0 context.Context, arg1 *ec2.DescribeKeyPairsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeKeyPairsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeKeyPairs", varargs...)
	ret0, _ := ret[0].(*ec2.DescribeKeyPairsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeKeyPairs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeKeyPairs", reflect.TypeOf((*MockEc2Client)(nil).DescribeKeyPairs), varargs...)
}

func (m *MockEc2Client) DescribeNatGateways(arg0 context.Context, arg1 *ec2.DescribeNatGatewaysInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeNatGatewaysOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeNatGateways", varargs...)
	ret0, _ := ret[0].(*ec2.DescribeNatGatewaysOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeNatGateways(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeNatGateways", reflect.TypeOf((*MockEc2Client)(nil).DescribeNatGateways), varargs...)
}

func (m *MockEc2Client) DescribeNetworkAcls(arg0 context.Context, arg1 *ec2.DescribeNetworkAclsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeNetworkAclsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeNetworkAcls", varargs...)
	ret0, _ := ret[0].(*ec2.DescribeNetworkAclsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeNetworkAcls(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeNetworkAcls", reflect.TypeOf((*MockEc2Client)(nil).DescribeNetworkAcls), varargs...)
}

func (m *MockEc2Client) DescribeNetworkInterfaces(arg0 context.Context, arg1 *ec2.DescribeNetworkInterfacesInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeNetworkInterfacesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeNetworkInterfaces", varargs...)
	ret0, _ := ret[0].(*ec2.DescribeNetworkInterfacesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeNetworkInterfaces(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeNetworkInterfaces", reflect.TypeOf((*MockEc2Client)(nil).DescribeNetworkInterfaces), varargs...)
}

func (m *MockEc2Client) DescribeRegions(arg0 context.Context, arg1 *ec2.DescribeRegionsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeRegionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeRegions", varargs...)
	ret0, _ := ret[0].(*ec2.DescribeRegionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeRegions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeRegions", reflect.TypeOf((*MockEc2Client)(nil).DescribeRegions), varargs...)
}

func (m *MockEc2Client) DescribeReservedInstances(arg0 context.Context, arg1 *ec2.DescribeReservedInstancesInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeReservedInstancesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeReservedInstances", varargs...)
	ret0, _ := ret[0].(*ec2.DescribeReservedInstancesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeReservedInstances(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeReservedInstances", reflect.TypeOf((*MockEc2Client)(nil).DescribeReservedInstances), varargs...)
}

func (m *MockEc2Client) DescribeRouteTables(arg0 context.Context, arg1 *ec2.DescribeRouteTablesInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeRouteTablesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeRouteTables", varargs...)
	ret0, _ := ret[0].(*ec2.DescribeRouteTablesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeRouteTables(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeRouteTables", reflect.TypeOf((*MockEc2Client)(nil).DescribeRouteTables), varargs...)
}

func (m *MockEc2Client) DescribeSecurityGroups(arg0 context.Context, arg1 *ec2.DescribeSecurityGroupsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeSecurityGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeSecurityGroups", varargs...)
	ret0, _ := ret[0].(*ec2.DescribeSecurityGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeSecurityGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeSecurityGroups", reflect.TypeOf((*MockEc2Client)(nil).DescribeSecurityGroups), varargs...)
}

func (m *MockEc2Client) DescribeSnapshotAttribute(arg0 context.Context, arg1 *ec2.DescribeSnapshotAttributeInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeSnapshotAttributeOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeSnapshotAttribute", varargs...)
	ret0, _ := ret[0].(*ec2.DescribeSnapshotAttributeOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeSnapshotAttribute(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeSnapshotAttribute", reflect.TypeOf((*MockEc2Client)(nil).DescribeSnapshotAttribute), varargs...)
}

func (m *MockEc2Client) DescribeSnapshots(arg0 context.Context, arg1 *ec2.DescribeSnapshotsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeSnapshotsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeSnapshots", varargs...)
	ret0, _ := ret[0].(*ec2.DescribeSnapshotsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeSnapshots(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeSnapshots", reflect.TypeOf((*MockEc2Client)(nil).DescribeSnapshots), varargs...)
}

func (m *MockEc2Client) DescribeSubnets(arg0 context.Context, arg1 *ec2.DescribeSubnetsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeSubnetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeSubnets", varargs...)
	ret0, _ := ret[0].(*ec2.DescribeSubnetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeSubnets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeSubnets", reflect.TypeOf((*MockEc2Client)(nil).DescribeSubnets), varargs...)
}

func (m *MockEc2Client) DescribeTransitGatewayAttachments(arg0 context.Context, arg1 *ec2.DescribeTransitGatewayAttachmentsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeTransitGatewayAttachmentsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeTransitGatewayAttachments", varargs...)
	ret0, _ := ret[0].(*ec2.DescribeTransitGatewayAttachmentsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeTransitGatewayAttachments(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeTransitGatewayAttachments", reflect.TypeOf((*MockEc2Client)(nil).DescribeTransitGatewayAttachments), varargs...)
}

func (m *MockEc2Client) DescribeTransitGatewayMulticastDomains(arg0 context.Context, arg1 *ec2.DescribeTransitGatewayMulticastDomainsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeTransitGatewayMulticastDomainsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeTransitGatewayMulticastDomains", varargs...)
	ret0, _ := ret[0].(*ec2.DescribeTransitGatewayMulticastDomainsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeTransitGatewayMulticastDomains(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeTransitGatewayMulticastDomains", reflect.TypeOf((*MockEc2Client)(nil).DescribeTransitGatewayMulticastDomains), varargs...)
}

func (m *MockEc2Client) DescribeTransitGatewayPeeringAttachments(arg0 context.Context, arg1 *ec2.DescribeTransitGatewayPeeringAttachmentsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeTransitGatewayPeeringAttachmentsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeTransitGatewayPeeringAttachments", varargs...)
	ret0, _ := ret[0].(*ec2.DescribeTransitGatewayPeeringAttachmentsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeTransitGatewayPeeringAttachments(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeTransitGatewayPeeringAttachments", reflect.TypeOf((*MockEc2Client)(nil).DescribeTransitGatewayPeeringAttachments), varargs...)
}

func (m *MockEc2Client) DescribeTransitGatewayRouteTables(arg0 context.Context, arg1 *ec2.DescribeTransitGatewayRouteTablesInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeTransitGatewayRouteTablesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeTransitGatewayRouteTables", varargs...)
	ret0, _ := ret[0].(*ec2.DescribeTransitGatewayRouteTablesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeTransitGatewayRouteTables(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeTransitGatewayRouteTables", reflect.TypeOf((*MockEc2Client)(nil).DescribeTransitGatewayRouteTables), varargs...)
}

func (m *MockEc2Client) DescribeTransitGatewayVpcAttachments(arg0 context.Context, arg1 *ec2.DescribeTransitGatewayVpcAttachmentsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeTransitGatewayVpcAttachmentsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeTransitGatewayVpcAttachments", varargs...)
	ret0, _ := ret[0].(*ec2.DescribeTransitGatewayVpcAttachmentsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeTransitGatewayVpcAttachments(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeTransitGatewayVpcAttachments", reflect.TypeOf((*MockEc2Client)(nil).DescribeTransitGatewayVpcAttachments), varargs...)
}

func (m *MockEc2Client) DescribeTransitGateways(arg0 context.Context, arg1 *ec2.DescribeTransitGatewaysInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeTransitGatewaysOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeTransitGateways", varargs...)
	ret0, _ := ret[0].(*ec2.DescribeTransitGatewaysOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeTransitGateways(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeTransitGateways", reflect.TypeOf((*MockEc2Client)(nil).DescribeTransitGateways), varargs...)
}

func (m *MockEc2Client) DescribeVolumes(arg0 context.Context, arg1 *ec2.DescribeVolumesInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeVolumesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeVolumes", varargs...)
	ret0, _ := ret[0].(*ec2.DescribeVolumesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeVolumes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeVolumes", reflect.TypeOf((*MockEc2Client)(nil).DescribeVolumes), varargs...)
}

func (m *MockEc2Client) DescribeVpcEndpointServiceConfigurations(arg0 context.Context, arg1 *ec2.DescribeVpcEndpointServiceConfigurationsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeVpcEndpointServiceConfigurationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeVpcEndpointServiceConfigurations", varargs...)
	ret0, _ := ret[0].(*ec2.DescribeVpcEndpointServiceConfigurationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeVpcEndpointServiceConfigurations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeVpcEndpointServiceConfigurations", reflect.TypeOf((*MockEc2Client)(nil).DescribeVpcEndpointServiceConfigurations), varargs...)
}

func (m *MockEc2Client) DescribeVpcEndpointServices(arg0 context.Context, arg1 *ec2.DescribeVpcEndpointServicesInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeVpcEndpointServicesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeVpcEndpointServices", varargs...)
	ret0, _ := ret[0].(*ec2.DescribeVpcEndpointServicesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeVpcEndpointServices(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeVpcEndpointServices", reflect.TypeOf((*MockEc2Client)(nil).DescribeVpcEndpointServices), varargs...)
}

func (m *MockEc2Client) DescribeVpcEndpoints(arg0 context.Context, arg1 *ec2.DescribeVpcEndpointsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeVpcEndpointsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeVpcEndpoints", varargs...)
	ret0, _ := ret[0].(*ec2.DescribeVpcEndpointsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeVpcEndpoints(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeVpcEndpoints", reflect.TypeOf((*MockEc2Client)(nil).DescribeVpcEndpoints), varargs...)
}

func (m *MockEc2Client) DescribeVpcPeeringConnections(arg0 context.Context, arg1 *ec2.DescribeVpcPeeringConnectionsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeVpcPeeringConnectionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeVpcPeeringConnections", varargs...)
	ret0, _ := ret[0].(*ec2.DescribeVpcPeeringConnectionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeVpcPeeringConnections(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeVpcPeeringConnections", reflect.TypeOf((*MockEc2Client)(nil).DescribeVpcPeeringConnections), varargs...)
}

func (m *MockEc2Client) DescribeVpcs(arg0 context.Context, arg1 *ec2.DescribeVpcsInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeVpcsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeVpcs", varargs...)
	ret0, _ := ret[0].(*ec2.DescribeVpcsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeVpcs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeVpcs", reflect.TypeOf((*MockEc2Client)(nil).DescribeVpcs), varargs...)
}

func (m *MockEc2Client) DescribeVpnGateways(arg0 context.Context, arg1 *ec2.DescribeVpnGatewaysInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeVpnGatewaysOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeVpnGateways", varargs...)
	ret0, _ := ret[0].(*ec2.DescribeVpnGatewaysOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) DescribeVpnGateways(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeVpnGateways", reflect.TypeOf((*MockEc2Client)(nil).DescribeVpnGateways), varargs...)
}

func (m *MockEc2Client) GetEbsDefaultKmsKeyId(arg0 context.Context, arg1 *ec2.GetEbsDefaultKmsKeyIdInput, arg2 ...func(*ec2.Options)) (*ec2.GetEbsDefaultKmsKeyIdOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetEbsDefaultKmsKeyId", varargs...)
	ret0, _ := ret[0].(*ec2.GetEbsDefaultKmsKeyIdOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) GetEbsDefaultKmsKeyId(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEbsDefaultKmsKeyId", reflect.TypeOf((*MockEc2Client)(nil).GetEbsDefaultKmsKeyId), varargs...)
}

func (m *MockEc2Client) GetEbsEncryptionByDefault(arg0 context.Context, arg1 *ec2.GetEbsEncryptionByDefaultInput, arg2 ...func(*ec2.Options)) (*ec2.GetEbsEncryptionByDefaultOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetEbsEncryptionByDefault", varargs...)
	ret0, _ := ret[0].(*ec2.GetEbsEncryptionByDefaultOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEc2ClientMockRecorder) GetEbsEncryptionByDefault(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEbsEncryptionByDefault", reflect.TypeOf((*MockEc2Client)(nil).GetEbsEncryptionByDefault), varargs...)
}
