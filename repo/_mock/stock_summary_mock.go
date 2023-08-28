// Code generated by MockGen. DO NOT EDIT.
// Source: ./init.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	redis "github.com/go-redis/redis/v8"
	gomock "github.com/golang/mock/gomock"
)

// MockRedisClient is a mock of RedisClient interface.
type MockRedisClient struct {
	ctrl     *gomock.Controller
	recorder *MockRedisClientMockRecorder
}

// MockRedisClientMockRecorder is the mock recorder for MockRedisClient.
type MockRedisClientMockRecorder struct {
	mock *MockRedisClient
}

// NewMockRedisClient creates a new mock instance.
func NewMockRedisClient(ctrl *gomock.Controller) *MockRedisClient {
	mock := &MockRedisClient{ctrl: ctrl}
	mock.recorder = &MockRedisClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRedisClient) EXPECT() *MockRedisClientMockRecorder {
	return m.recorder
}

// ZAdd mocks base method.
func (m *MockRedisClient) ZAdd(ctx context.Context, key string, members ...*redis.Z) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range members {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ZAdd", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ZAdd indicates an expected call of ZAdd.
func (mr *MockRedisClientMockRecorder) ZAdd(ctx, key interface{}, members ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, members...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZAdd", reflect.TypeOf((*MockRedisClient)(nil).ZAdd), varargs...)
}

// ZRangeByScore mocks base method.
func (m *MockRedisClient) ZRangeByScore(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZRangeByScore", ctx, key, opt)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// ZRangeByScore indicates an expected call of ZRangeByScore.
func (mr *MockRedisClientMockRecorder) ZRangeByScore(ctx, key, opt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRangeByScore", reflect.TypeOf((*MockRedisClient)(nil).ZRangeByScore), ctx, key, opt)
}

// ZRem mocks base method.
func (m *MockRedisClient) ZRem(ctx context.Context, key string, members ...interface{}) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range members {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ZRem", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ZRem indicates an expected call of ZRem.
func (mr *MockRedisClientMockRecorder) ZRem(ctx, key interface{}, members ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, members...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRem", reflect.TypeOf((*MockRedisClient)(nil).ZRem), varargs...)
}
