// Code generated by MockGen. DO NOT EDIT.
// Source: Go_cinema_reconstructed/model (interfaces: MovieRepository)

// Package mock_model is a generated GoMock package.
package mock_model

import (
	model "Go_cinema_reconstructed/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockMovieRepository is a mock of MovieRepository interface.
type MockMovieRepository struct {
	ctrl     *gomock.Controller
	recorder *MockMovieRepositoryMockRecorder
}

// MockMovieRepositoryMockRecorder is the mock recorder for MockMovieRepository.
type MockMovieRepositoryMockRecorder struct {
	mock *MockMovieRepository
}

// NewMockMovieRepository creates a new mock instance.
func NewMockMovieRepository(ctrl *gomock.Controller) *MockMovieRepository {
	mock := &MockMovieRepository{ctrl: ctrl}
	mock.recorder = &MockMovieRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMovieRepository) EXPECT() *MockMovieRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockMovieRepository) Create(arg0 *model.MovieReq) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockMovieRepositoryMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockMovieRepository)(nil).Create), arg0)
}

// Delete mocks base method.
func (m *MockMovieRepository) Delete(arg0 string) (int32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(int32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockMovieRepositoryMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockMovieRepository)(nil).Delete), arg0)
}

// GetAll mocks base method.
func (m *MockMovieRepository) GetAll() ([]*model.MovieRes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]*model.MovieRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockMovieRepositoryMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockMovieRepository)(nil).GetAll))
}

// GetByID mocks base method.
func (m *MockMovieRepository) GetByID(arg0 string) (*model.MovieRes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", arg0)
	ret0, _ := ret[0].(*model.MovieRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockMovieRepositoryMockRecorder) GetByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockMovieRepository)(nil).GetByID), arg0)
}

// GetWatchedMovies mocks base method.
func (m *MockMovieRepository) GetWatchedMovies(arg0 []string) ([]*model.MovieRes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWatchedMovies", arg0)
	ret0, _ := ret[0].([]*model.MovieRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWatchedMovies indicates an expected call of GetWatchedMovies.
func (mr *MockMovieRepositoryMockRecorder) GetWatchedMovies(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWatchedMovies", reflect.TypeOf((*MockMovieRepository)(nil).GetWatchedMovies), arg0)
}

// Update mocks base method.
func (m *MockMovieRepository) Update(arg0 string, arg1 *model.MovieReq) (int32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(int32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockMovieRepositoryMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockMovieRepository)(nil).Update), arg0, arg1)
}
