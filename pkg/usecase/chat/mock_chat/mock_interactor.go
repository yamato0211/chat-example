// Code generated by MockGen. DO NOT EDIT.
// Source: interactor.go

// Package mock_chat is a generated GoMock package.
package mock_chat

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	entity "github.com/sivchari/chat-example/pkg/domain/entity"
)

// MockInteractor is a mock of Interactor interface.
type MockInteractor struct {
	ctrl     *gomock.Controller
	recorder *MockInteractorMockRecorder
}

// MockInteractorMockRecorder is the mock recorder for MockInteractor.
type MockInteractorMockRecorder struct {
	mock *MockInteractor
}

// NewMockInteractor creates a new mock instance.
func NewMockInteractor(ctrl *gomock.Controller) *MockInteractor {
	mock := &MockInteractor{ctrl: ctrl}
	mock.recorder = &MockInteractorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInteractor) EXPECT() *MockInteractorMockRecorder {
	return m.recorder
}

// CreateRoom mocks base method.
func (m *MockInteractor) CreateRoom(ctx context.Context, name string) (*entity.Room, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRoom", ctx, name)
	ret0, _ := ret[0].(*entity.Room)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRoom indicates an expected call of CreateRoom.
func (mr *MockInteractorMockRecorder) CreateRoom(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRoom", reflect.TypeOf((*MockInteractor)(nil).CreateRoom), ctx, name)
}

// GetPass mocks base method.
func (m *MockInteractor) GetPass(ctx context.Context) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPass", ctx)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPass indicates an expected call of GetPass.
func (mr *MockInteractorMockRecorder) GetPass(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPass", reflect.TypeOf((*MockInteractor)(nil).GetPass), ctx)
}

// GetRoom mocks base method.
func (m *MockInteractor) GetRoom(ctx context.Context, id string) (*entity.Room, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoom", ctx, id)
	ret0, _ := ret[0].(*entity.Room)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoom indicates an expected call of GetRoom.
func (mr *MockInteractorMockRecorder) GetRoom(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoom", reflect.TypeOf((*MockInteractor)(nil).GetRoom), ctx, id)
}

// ListMessage mocks base method.
func (m *MockInteractor) ListMessage(ctx context.Context, roomID string) ([]*entity.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListMessage", ctx, roomID)
	ret0, _ := ret[0].([]*entity.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListMessage indicates an expected call of ListMessage.
func (mr *MockInteractorMockRecorder) ListMessage(ctx, roomID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMessage", reflect.TypeOf((*MockInteractor)(nil).ListMessage), ctx, roomID)
}

// ListRoom mocks base method.
func (m *MockInteractor) ListRoom(ctx context.Context) ([]*entity.Room, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRoom", ctx)
	ret0, _ := ret[0].([]*entity.Room)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRoom indicates an expected call of ListRoom.
func (mr *MockInteractorMockRecorder) ListRoom(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRoom", reflect.TypeOf((*MockInteractor)(nil).ListRoom), ctx)
}

// SendMessage mocks base method.
func (m *MockInteractor) SendMessage(ctx context.Context, roomID, text string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMessage", ctx, roomID, text)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMessage indicates an expected call of SendMessage.
func (mr *MockInteractorMockRecorder) SendMessage(ctx, roomID, text interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMessage", reflect.TypeOf((*MockInteractor)(nil).SendMessage), ctx, roomID, text)
}
