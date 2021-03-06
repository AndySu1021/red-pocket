// Code generated by MockGen. DO NOT EDIT.
// Source: util/interface/cache.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockICache is a mock of ICache interface.
type MockICache struct {
	ctrl     *gomock.Controller
	recorder *MockICacheMockRecorder
}

// MockICacheMockRecorder is the mock recorder for MockICache.
type MockICacheMockRecorder struct {
	mock *MockICache
}

// NewMockICache creates a new mock instance.
func NewMockICache(ctrl *gomock.Controller) *MockICache {
	mock := &MockICache{ctrl: ctrl}
	mock.recorder = &MockICacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockICache) EXPECT() *MockICacheMockRecorder {
	return m.recorder
}

// Expire mocks base method.
func (m *MockICache) Expire(ctx context.Context, key string, expireAt time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Expire", ctx, key, expireAt)
	ret0, _ := ret[0].(error)
	return ret0
}

// Expire indicates an expected call of Expire.
func (mr *MockICacheMockRecorder) Expire(ctx, key, expireAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Expire", reflect.TypeOf((*MockICache)(nil).Expire), ctx, key, expireAt)
}

// Get mocks base method.
func (m *MockICache) Get(ctx context.Context, key string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, key)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockICacheMockRecorder) Get(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockICache)(nil).Get), ctx, key)
}

// LPush mocks base method.
func (m *MockICache) LPush(ctx context.Context, key string, values ...interface{}) (int64, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range values {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "LPush", varargs...)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LPush indicates an expected call of LPush.
func (mr *MockICacheMockRecorder) LPush(ctx, key interface{}, values ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, values...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LPush", reflect.TypeOf((*MockICache)(nil).LPush), varargs...)
}

// RPop mocks base method.
func (m *MockICache) RPop(ctx context.Context, key string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RPop", ctx, key)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RPop indicates an expected call of RPop.
func (mr *MockICacheMockRecorder) RPop(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RPop", reflect.TypeOf((*MockICache)(nil).RPop), ctx, key)
}

// SetEX mocks base method.
func (m *MockICache) SetEX(ctx context.Context, key string, data interface{}, expireAt time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetEX", ctx, key, data, expireAt)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetEX indicates an expected call of SetEX.
func (mr *MockICacheMockRecorder) SetEX(ctx, key, data, expireAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetEX", reflect.TypeOf((*MockICache)(nil).SetEX), ctx, key, data, expireAt)
}

// SetNX mocks base method.
func (m *MockICache) SetNX(ctx context.Context, key string, data interface{}, expireAt time.Duration) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetNX", ctx, key, data, expireAt)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetNX indicates an expected call of SetNX.
func (mr *MockICacheMockRecorder) SetNX(ctx, key, data, expireAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetNX", reflect.TypeOf((*MockICache)(nil).SetNX), ctx, key, data, expireAt)
}
