// Code generated by MockGen. DO NOT EDIT.
// Source: route256/loms/pkg/loms_v1 (interfaces: LomsV1Client)

// Package mock_loms_v1 is a generated GoMock package.
package mock_loms_v1

import (
	context "context"
	reflect "reflect"
	loms "route256/loms/pkg/loms_v1"

	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// MockLomsV1Client is a mock of LomsV1Client interface.
type MockLomsV1Client struct {
	ctrl     *gomock.Controller
	recorder *MockLomsV1ClientMockRecorder
}

// MockLomsV1ClientMockRecorder is the mock recorder for MockLomsV1Client.
type MockLomsV1ClientMockRecorder struct {
	mock *MockLomsV1Client
}

// NewMockLomsV1Client creates a new mock instance.
func NewMockLomsV1Client(ctrl *gomock.Controller) *MockLomsV1Client {
	mock := &MockLomsV1Client{ctrl: ctrl}
	mock.recorder = &MockLomsV1ClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLomsV1Client) EXPECT() *MockLomsV1ClientMockRecorder {
	return m.recorder
}

// CancelOrder mocks base method.
func (m *MockLomsV1Client) CancelOrder(arg0 context.Context, arg1 *loms.CancelOrderRequest, arg2 ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CancelOrder", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CancelOrder indicates an expected call of CancelOrder.
func (mr *MockLomsV1ClientMockRecorder) CancelOrder(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelOrder", reflect.TypeOf((*MockLomsV1Client)(nil).CancelOrder), varargs...)
}

// CreateOrder mocks base method.
func (m *MockLomsV1Client) CreateOrder(arg0 context.Context, arg1 *loms.CreateOrderRequest, arg2 ...grpc.CallOption) (*loms.CreateOrderResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateOrder", varargs...)
	ret0, _ := ret[0].(*loms.CreateOrderResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrder indicates an expected call of CreateOrder.
func (mr *MockLomsV1ClientMockRecorder) CreateOrder(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrder", reflect.TypeOf((*MockLomsV1Client)(nil).CreateOrder), varargs...)
}

// ListOrder mocks base method.
func (m *MockLomsV1Client) ListOrder(arg0 context.Context, arg1 *loms.ListOrderRequest, arg2 ...grpc.CallOption) (*loms.ListOrderResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListOrder", varargs...)
	ret0, _ := ret[0].(*loms.ListOrderResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOrder indicates an expected call of ListOrder.
func (mr *MockLomsV1ClientMockRecorder) ListOrder(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOrder", reflect.TypeOf((*MockLomsV1Client)(nil).ListOrder), varargs...)
}

// OrderPayed mocks base method.
func (m *MockLomsV1Client) OrderPayed(arg0 context.Context, arg1 *loms.OrderPayedRequest, arg2 ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "OrderPayed", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OrderPayed indicates an expected call of OrderPayed.
func (mr *MockLomsV1ClientMockRecorder) OrderPayed(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrderPayed", reflect.TypeOf((*MockLomsV1Client)(nil).OrderPayed), varargs...)
}

// Stocks mocks base method.
func (m *MockLomsV1Client) Stocks(arg0 context.Context, arg1 *loms.StocksRequest, arg2 ...grpc.CallOption) (*loms.StocksResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Stocks", varargs...)
	ret0, _ := ret[0].(*loms.StocksResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Stocks indicates an expected call of Stocks.
func (mr *MockLomsV1ClientMockRecorder) Stocks(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stocks", reflect.TypeOf((*MockLomsV1Client)(nil).Stocks), varargs...)
}
