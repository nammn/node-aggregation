// Code generated by MockGen. DO NOT EDIT.
// Source: database/redis.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	database "github.com/nammn/node-aggregation/database"
	reflect "reflect"
)

// MockRedisConnection is a mock of RedisConnection interface
type MockRedisConnection struct {
	ctrl     *gomock.Controller
	recorder *MockRedisConnectionMockRecorder
}

// MockRedisConnectionMockRecorder is the mock recorder for MockRedisConnection
type MockRedisConnectionMockRecorder struct {
	mock *MockRedisConnection
}

// NewMockRedisConnection creates a new mock instance
func NewMockRedisConnection(ctrl *gomock.Controller) *MockRedisConnection {
	mock := &MockRedisConnection{ctrl: ctrl}
	mock.recorder = &MockRedisConnectionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRedisConnection) EXPECT() *MockRedisConnectionMockRecorder {
	return m.recorder
}

// SaveNodeStatValue mocks base method
func (m *MockRedisConnection) SaveNodeStatValue(key database.NodeStat) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveNodeStatValue", key)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveNodeStatValue indicates an expected call of SaveNodeStatValue
func (mr *MockRedisConnectionMockRecorder) SaveNodeStatValue(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveNodeStatValue", reflect.TypeOf((*MockRedisConnection)(nil).SaveNodeStatValue), key)
}

// Ping mocks base method
func (m *MockRedisConnection) Ping() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping")
	ret0, _ := ret[0].(error)
	return ret0
}

// Ping indicates an expected call of Ping
func (mr *MockRedisConnectionMockRecorder) Ping() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockRedisConnection)(nil).Ping))
}
