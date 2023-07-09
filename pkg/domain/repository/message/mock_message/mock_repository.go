// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mock_message is a generated GoMock package.
package mock_message

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	entity "github.com/sivchari/chat-example/pkg/domain/entity"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockRepository) Insert(ctx context.Context, message *entity.Message) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", ctx, message)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockRepositoryMockRecorder) Insert(ctx, message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockRepository)(nil).Insert), ctx, message)
}

// SelectByRoomID mocks base method.
func (m *MockRepository) SelectByRoomID(ctx context.Context, roomID string) ([]*entity.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectByRoomID", ctx, roomID)
	ret0, _ := ret[0].([]*entity.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectByRoomID indicates an expected call of SelectByRoomID.
func (mr *MockRepositoryMockRecorder) SelectByRoomID(ctx, roomID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectByRoomID", reflect.TypeOf((*MockRepository)(nil).SelectByRoomID), ctx, roomID)
}