// Code generated by MockGen. DO NOT EDIT.
// Source: manager.go

// Package mock_manager is a generated GoMock package.
package mock_manager

import (
	context "context"
	reflect "reflect"

	q "g.hz.netease.com/horizon/lib/q"
	models "g.hz.netease.com/horizon/pkg/cluster/models"
	models0 "g.hz.netease.com/horizon/pkg/tag/models"
	gomock "github.com/golang/mock/gomock"
)

// MockManager is a mock of Manager interface.
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
}

// MockManagerMockRecorder is the mock recorder for MockManager.
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance.
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// CheckClusterExists mocks base method.
func (m *MockManager) CheckClusterExists(ctx context.Context, cluster string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckClusterExists", ctx, cluster)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckClusterExists indicates an expected call of CheckClusterExists.
func (mr *MockManagerMockRecorder) CheckClusterExists(ctx, cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckClusterExists", reflect.TypeOf((*MockManager)(nil).CheckClusterExists), ctx, cluster)
}

// Create mocks base method.
func (m *MockManager) Create(ctx context.Context, cluster *models.Cluster, tags []*models0.Tag, extraMembers map[string]string) (*models.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, cluster, tags, extraMembers)
	ret0, _ := ret[0].(*models.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockManagerMockRecorder) Create(ctx, cluster, tags, extraMembers interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockManager)(nil).Create), ctx, cluster, tags, extraMembers)
}

// DeleteByID mocks base method.
func (m *MockManager) DeleteByID(ctx context.Context, id uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID.
func (mr *MockManagerMockRecorder) DeleteByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockManager)(nil).DeleteByID), ctx, id)
}

// GetByID mocks base method.
func (m *MockManager) GetByID(ctx context.Context, id uint) (*models.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, id)
	ret0, _ := ret[0].(*models.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockManagerMockRecorder) GetByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockManager)(nil).GetByID), ctx, id)
}

// GetByIDIncludeSoftDelete mocks base method.
func (m *MockManager) GetByIDIncludeSoftDelete(ctx context.Context, id uint) (*models.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByIDIncludeSoftDelete", ctx, id)
	ret0, _ := ret[0].(*models.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByIDIncludeSoftDelete indicates an expected call of GetByIDIncludeSoftDelete.
func (mr *MockManagerMockRecorder) GetByIDIncludeSoftDelete(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByIDIncludeSoftDelete", reflect.TypeOf((*MockManager)(nil).GetByIDIncludeSoftDelete), ctx, id)
}

// GetByName mocks base method.
func (m *MockManager) GetByName(ctx context.Context, clusterName string) (*models.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByName", ctx, clusterName)
	ret0, _ := ret[0].(*models.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByName indicates an expected call of GetByName.
func (mr *MockManagerMockRecorder) GetByName(ctx, clusterName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByName", reflect.TypeOf((*MockManager)(nil).GetByName), ctx, clusterName)
}

// List mocks base method.
func (m *MockManager) List(ctx context.Context, query *q.Query, appIDs ...uint) (int, []*models.ClusterWithRegion, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, query}
	for _, a := range appIDs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].([]*models.ClusterWithRegion)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List.
func (mr *MockManagerMockRecorder) List(ctx, query interface{}, appIDs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, query}, appIDs...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockManager)(nil).List), varargs...)
}

// ListByApplicationID mocks base method.
func (m *MockManager) ListByApplicationID(ctx context.Context, applicationID uint) (int, []*models.ClusterWithRegion, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByApplicationID", ctx, applicationID)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].([]*models.ClusterWithRegion)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListByApplicationID indicates an expected call of ListByApplicationID.
func (mr *MockManagerMockRecorder) ListByApplicationID(ctx, applicationID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByApplicationID", reflect.TypeOf((*MockManager)(nil).ListByApplicationID), ctx, applicationID)
}

// ListClusterWithExpiry mocks base method.
func (m *MockManager) ListClusterWithExpiry(ctx context.Context, query *q.Query) ([]*models.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListClusterWithExpiry", ctx, query)
	ret0, _ := ret[0].([]*models.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListClusterWithExpiry indicates an expected call of ListClusterWithExpiry.
func (mr *MockManagerMockRecorder) ListClusterWithExpiry(ctx, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListClusterWithExpiry", reflect.TypeOf((*MockManager)(nil).ListClusterWithExpiry), ctx, query)
}

// UpdateByID mocks base method.
func (m *MockManager) UpdateByID(ctx context.Context, id uint, cluster *models.Cluster) (*models.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateByID", ctx, id, cluster)
	ret0, _ := ret[0].(*models.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateByID indicates an expected call of UpdateByID.
func (mr *MockManagerMockRecorder) UpdateByID(ctx, id, cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateByID", reflect.TypeOf((*MockManager)(nil).UpdateByID), ctx, id, cluster)
}
