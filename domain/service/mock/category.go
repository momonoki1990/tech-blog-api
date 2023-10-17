// Code generated by MockGen. DO NOT EDIT.
// Source: ./domain/service/category.go
//
// Generated by this command:
//
//	mockgen -source=./domain/service/category.go -destination=./domain/service/mock/category.go
//
// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"

	model "github.com/momonoki1990/tech-blog-api/domain/model"
	gomock "go.uber.org/mock/gomock"
)

// MockCategoryCreator is a mock of CategoryCreator interface.
type MockCategoryCreator struct {
	ctrl     *gomock.Controller
	recorder *MockCategoryCreatorMockRecorder
}

// MockCategoryCreatorMockRecorder is the mock recorder for MockCategoryCreator.
type MockCategoryCreatorMockRecorder struct {
	mock *MockCategoryCreator
}

// NewMockCategoryCreator creates a new mock instance.
func NewMockCategoryCreator(ctrl *gomock.Controller) *MockCategoryCreator {
	mock := &MockCategoryCreator{ctrl: ctrl}
	mock.recorder = &MockCategoryCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCategoryCreator) EXPECT() *MockCategoryCreatorMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockCategoryCreator) Create(name string, displayOrder int) (*model.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", name, displayOrder)
	ret0, _ := ret[0].(*model.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockCategoryCreatorMockRecorder) Create(name, displayOrder any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCategoryCreator)(nil).Create), name, displayOrder)
}
