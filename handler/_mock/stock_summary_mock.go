// Code generated by MockGen. DO NOT EDIT.
// Source: ./init.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	model "stock/model"
	gomock "github.com/golang/mock/gomock"
)

// MockStockUsecase is a mock of StockUsecase interface.
type MockStockUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockStockUsecaseMockRecorder
}

// MockStockUsecaseMockRecorder is the mock recorder for MockStockUsecase.
type MockStockUsecaseMockRecorder struct {
	mock *MockStockUsecase
}

// NewMockStockUsecase creates a new mock instance.
func NewMockStockUsecase(ctrl *gomock.Controller) *MockStockUsecase {
	mock := &MockStockUsecase{ctrl: ctrl}
	mock.recorder = &MockStockUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStockUsecase) EXPECT() *MockStockUsecaseMockRecorder {
	return m.recorder
}

// GetStockSummary mocks base method.
func (m *MockStockUsecase) GetStockSummary(ctx context.Context, request model.GetStockSummaryRequest) ([]model.Summary, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStockSummary", ctx, request)
	ret0, _ := ret[0].([]model.Summary)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStockSummary indicates an expected call of GetStockSummary.
func (mr *MockStockUsecaseMockRecorder) GetStockSummary(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStockSummary", reflect.TypeOf((*MockStockUsecase)(nil).GetStockSummary), ctx, request)
}

// UpdateStockSummary mocks base method.
func (m *MockStockUsecase) UpdateStockSummary(ctx context.Context, transaction model.Transaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStockSummary", ctx, transaction)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStockSummary indicates an expected call of UpdateStockSummary.
func (mr *MockStockUsecaseMockRecorder) UpdateStockSummary(ctx, transaction interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStockSummary", reflect.TypeOf((*MockStockUsecase)(nil).UpdateStockSummary), ctx, transaction)
}
