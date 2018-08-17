// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/google/ts-bridge/stackdriver (interfaces: MetricClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	metric "google.golang.org/genproto/googleapis/api/metric"
	v3 "google.golang.org/genproto/googleapis/monitoring/v3"
	reflect "reflect"
)

// MockMetricClient is a mock of MetricClient interface
type MockMetricClient struct {
	ctrl     *gomock.Controller
	recorder *MockMetricClientMockRecorder
}

// MockMetricClientMockRecorder is the mock recorder for MockMetricClient
type MockMetricClientMockRecorder struct {
	mock *MockMetricClient
}

// NewMockMetricClient creates a new mock instance
func NewMockMetricClient(ctrl *gomock.Controller) *MockMetricClient {
	mock := &MockMetricClient{ctrl: ctrl}
	mock.recorder = &MockMetricClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMetricClient) EXPECT() *MockMetricClientMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockMetricClient) Close() error {
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockMetricClientMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockMetricClient)(nil).Close))
}

// CreateMetricDescriptor mocks base method
func (m *MockMetricClient) CreateMetricDescriptor(arg0 context.Context, arg1 *v3.CreateMetricDescriptorRequest) (*metric.MetricDescriptor, error) {
	ret := m.ctrl.Call(m, "CreateMetricDescriptor", arg0, arg1)
	ret0, _ := ret[0].(*metric.MetricDescriptor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMetricDescriptor indicates an expected call of CreateMetricDescriptor
func (mr *MockMetricClientMockRecorder) CreateMetricDescriptor(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMetricDescriptor", reflect.TypeOf((*MockMetricClient)(nil).CreateMetricDescriptor), arg0, arg1)
}

// CreateTimeSeries mocks base method
func (m *MockMetricClient) CreateTimeSeries(arg0 context.Context, arg1 *v3.CreateTimeSeriesRequest) error {
	ret := m.ctrl.Call(m, "CreateTimeSeries", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateTimeSeries indicates an expected call of CreateTimeSeries
func (mr *MockMetricClientMockRecorder) CreateTimeSeries(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTimeSeries", reflect.TypeOf((*MockMetricClient)(nil).CreateTimeSeries), arg0, arg1)
}

// DeleteMetricDescriptor mocks base method
func (m *MockMetricClient) DeleteMetricDescriptor(arg0 context.Context, arg1 *v3.DeleteMetricDescriptorRequest) error {
	ret := m.ctrl.Call(m, "DeleteMetricDescriptor", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMetricDescriptor indicates an expected call of DeleteMetricDescriptor
func (mr *MockMetricClientMockRecorder) DeleteMetricDescriptor(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMetricDescriptor", reflect.TypeOf((*MockMetricClient)(nil).DeleteMetricDescriptor), arg0, arg1)
}

// GetMetricDescriptor mocks base method
func (m *MockMetricClient) GetMetricDescriptor(arg0 context.Context, arg1 *v3.GetMetricDescriptorRequest) (*metric.MetricDescriptor, error) {
	ret := m.ctrl.Call(m, "GetMetricDescriptor", arg0, arg1)
	ret0, _ := ret[0].(*metric.MetricDescriptor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMetricDescriptor indicates an expected call of GetMetricDescriptor
func (mr *MockMetricClientMockRecorder) GetMetricDescriptor(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMetricDescriptor", reflect.TypeOf((*MockMetricClient)(nil).GetMetricDescriptor), arg0, arg1)
}

// ListTimeSeries mocks base method
func (m *MockMetricClient) ListTimeSeries(arg0 context.Context, arg1 *v3.ListTimeSeriesRequest) ([]*v3.TimeSeries, error) {
	ret := m.ctrl.Call(m, "ListTimeSeries", arg0, arg1)
	ret0, _ := ret[0].([]*v3.TimeSeries)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTimeSeries indicates an expected call of ListTimeSeries
func (mr *MockMetricClientMockRecorder) ListTimeSeries(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTimeSeries", reflect.TypeOf((*MockMetricClient)(nil).ListTimeSeries), arg0, arg1)
}
