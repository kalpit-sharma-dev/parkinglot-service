// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	models "github.com/kalpit-sharma-dev/parkinglot-service/src/models"
	reflect "reflect"
)

// MockParkingLotService is a mock of ParkingLotService interface
type MockParkingLotService struct {
	ctrl     *gomock.Controller
	recorder *MockParkingLotServiceMockRecorder
}

// MockParkingLotServiceMockRecorder is the mock recorder for MockParkingLotService
type MockParkingLotServiceMockRecorder struct {
	mock *MockParkingLotService
}

// NewMockParkingLotService creates a new mock instance
func NewMockParkingLotService(ctrl *gomock.Controller) *MockParkingLotService {
	mock := &MockParkingLotService{ctrl: ctrl}
	mock.recorder = &MockParkingLotServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockParkingLotService) EXPECT() *MockParkingLotServiceMockRecorder {
	return m.recorder
}

// CreateSlotEvent mocks base method
func (m *MockParkingLotService) CreateSlotEvent(ctx context.Context, req models.Slot) (models.Slot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSlotEvent", ctx, req)
	ret0, _ := ret[0].(models.Slot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSlotEvent indicates an expected call of CreateSlotEvent
func (mr *MockParkingLotServiceMockRecorder) CreateSlotEvent(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSlotEvent", reflect.TypeOf((*MockParkingLotService)(nil).CreateSlotEvent), ctx, req)
}

// GetAllCarsWithColor mocks base method
func (m *MockParkingLotService) GetAllCarsWithColor(ctx context.Context, reqColor string) ([]models.Vehicle, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllCarsWithColor", ctx, reqColor)
	ret0, _ := ret[0].([]models.Vehicle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllCarsWithColor indicates an expected call of GetAllCarsWithColor
func (mr *MockParkingLotServiceMockRecorder) GetAllCarsWithColor(ctx, reqColor interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllCarsWithColor", reflect.TypeOf((*MockParkingLotService)(nil).GetAllCarsWithColor), ctx, reqColor)
}

// GetSlotNumberWithCarID mocks base method
func (m *MockParkingLotService) GetSlotNumberWithCarID(ctx context.Context, reqNumber string) (models.Vehicle, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSlotNumberWithCarID", ctx, reqNumber)
	ret0, _ := ret[0].(models.Vehicle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSlotNumberWithCarID indicates an expected call of GetSlotNumberWithCarID
func (mr *MockParkingLotServiceMockRecorder) GetSlotNumberWithCarID(ctx, reqNumber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSlotNumberWithCarID", reflect.TypeOf((*MockParkingLotService)(nil).GetSlotNumberWithCarID), ctx, reqNumber)
}

// GetAllSlotNumberWithColor mocks base method
func (m *MockParkingLotService) GetAllSlotNumberWithColor(ctx context.Context, reqColor string) ([]models.Vehicle, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllSlotNumberWithColor", ctx, reqColor)
	ret0, _ := ret[0].([]models.Vehicle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllSlotNumberWithColor indicates an expected call of GetAllSlotNumberWithColor
func (mr *MockParkingLotServiceMockRecorder) GetAllSlotNumberWithColor(ctx, reqColor interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllSlotNumberWithColor", reflect.TypeOf((*MockParkingLotService)(nil).GetAllSlotNumberWithColor), ctx, reqColor)
}

// CreateParkEvent mocks base method
func (m *MockParkingLotService) CreateParkEvent(ctx context.Context, req models.Vehicle) (models.Vehicle, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateParkEvent", ctx, req)
	ret0, _ := ret[0].(models.Vehicle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateParkEvent indicates an expected call of CreateParkEvent
func (mr *MockParkingLotServiceMockRecorder) CreateParkEvent(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateParkEvent", reflect.TypeOf((*MockParkingLotService)(nil).CreateParkEvent), ctx, req)
}

// ExitParkEvent mocks base method
func (m *MockParkingLotService) ExitParkEvent(ctx context.Context, req models.Vehicle) (models.Vehicle, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExitParkEvent", ctx, req)
	ret0, _ := ret[0].(models.Vehicle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExitParkEvent indicates an expected call of ExitParkEvent
func (mr *MockParkingLotServiceMockRecorder) ExitParkEvent(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExitParkEvent", reflect.TypeOf((*MockParkingLotService)(nil).ExitParkEvent), ctx, req)
}
