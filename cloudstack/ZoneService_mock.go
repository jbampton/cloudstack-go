//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: ./cloudstack/ZoneService.go

// Package cloudstack is a generated GoMock package.
package cloudstack

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockZoneServiceIface is a mock of ZoneServiceIface interface.
type MockZoneServiceIface struct {
	ctrl     *gomock.Controller
	recorder *MockZoneServiceIfaceMockRecorder
}

// MockZoneServiceIfaceMockRecorder is the mock recorder for MockZoneServiceIface.
type MockZoneServiceIfaceMockRecorder struct {
	mock *MockZoneServiceIface
}

// NewMockZoneServiceIface creates a new mock instance.
func NewMockZoneServiceIface(ctrl *gomock.Controller) *MockZoneServiceIface {
	mock := &MockZoneServiceIface{ctrl: ctrl}
	mock.recorder = &MockZoneServiceIfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockZoneServiceIface) EXPECT() *MockZoneServiceIfaceMockRecorder {
	return m.recorder
}

// CreateZone mocks base method.
func (m *MockZoneServiceIface) CreateZone(p *CreateZoneParams) (*CreateZoneResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateZone", p)
	ret0, _ := ret[0].(*CreateZoneResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateZone indicates an expected call of CreateZone.
func (mr *MockZoneServiceIfaceMockRecorder) CreateZone(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateZone", reflect.TypeOf((*MockZoneServiceIface)(nil).CreateZone), p)
}

// DedicateZone mocks base method.
func (m *MockZoneServiceIface) DedicateZone(p *DedicateZoneParams) (*DedicateZoneResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DedicateZone", p)
	ret0, _ := ret[0].(*DedicateZoneResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DedicateZone indicates an expected call of DedicateZone.
func (mr *MockZoneServiceIfaceMockRecorder) DedicateZone(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DedicateZone", reflect.TypeOf((*MockZoneServiceIface)(nil).DedicateZone), p)
}

// DeleteZone mocks base method.
func (m *MockZoneServiceIface) DeleteZone(p *DeleteZoneParams) (*DeleteZoneResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteZone", p)
	ret0, _ := ret[0].(*DeleteZoneResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteZone indicates an expected call of DeleteZone.
func (mr *MockZoneServiceIfaceMockRecorder) DeleteZone(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteZone", reflect.TypeOf((*MockZoneServiceIface)(nil).DeleteZone), p)
}

// DisableHAForZone mocks base method.
func (m *MockZoneServiceIface) DisableHAForZone(p *DisableHAForZoneParams) (*DisableHAForZoneResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DisableHAForZone", p)
	ret0, _ := ret[0].(*DisableHAForZoneResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DisableHAForZone indicates an expected call of DisableHAForZone.
func (mr *MockZoneServiceIfaceMockRecorder) DisableHAForZone(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisableHAForZone", reflect.TypeOf((*MockZoneServiceIface)(nil).DisableHAForZone), p)
}

// DisableOutOfBandManagementForZone mocks base method.
func (m *MockZoneServiceIface) DisableOutOfBandManagementForZone(p *DisableOutOfBandManagementForZoneParams) (*DisableOutOfBandManagementForZoneResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DisableOutOfBandManagementForZone", p)
	ret0, _ := ret[0].(*DisableOutOfBandManagementForZoneResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DisableOutOfBandManagementForZone indicates an expected call of DisableOutOfBandManagementForZone.
func (mr *MockZoneServiceIfaceMockRecorder) DisableOutOfBandManagementForZone(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisableOutOfBandManagementForZone", reflect.TypeOf((*MockZoneServiceIface)(nil).DisableOutOfBandManagementForZone), p)
}

// EnableHAForZone mocks base method.
func (m *MockZoneServiceIface) EnableHAForZone(p *EnableHAForZoneParams) (*EnableHAForZoneResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnableHAForZone", p)
	ret0, _ := ret[0].(*EnableHAForZoneResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnableHAForZone indicates an expected call of EnableHAForZone.
func (mr *MockZoneServiceIfaceMockRecorder) EnableHAForZone(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableHAForZone", reflect.TypeOf((*MockZoneServiceIface)(nil).EnableHAForZone), p)
}

// EnableOutOfBandManagementForZone mocks base method.
func (m *MockZoneServiceIface) EnableOutOfBandManagementForZone(p *EnableOutOfBandManagementForZoneParams) (*EnableOutOfBandManagementForZoneResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnableOutOfBandManagementForZone", p)
	ret0, _ := ret[0].(*EnableOutOfBandManagementForZoneResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnableOutOfBandManagementForZone indicates an expected call of EnableOutOfBandManagementForZone.
func (mr *MockZoneServiceIfaceMockRecorder) EnableOutOfBandManagementForZone(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableOutOfBandManagementForZone", reflect.TypeOf((*MockZoneServiceIface)(nil).EnableOutOfBandManagementForZone), p)
}

// GetZoneByID mocks base method.
func (m *MockZoneServiceIface) GetZoneByID(id string, opts ...OptionFunc) (*Zone, int, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{id}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetZoneByID", varargs...)
	ret0, _ := ret[0].(*Zone)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetZoneByID indicates an expected call of GetZoneByID.
func (mr *MockZoneServiceIfaceMockRecorder) GetZoneByID(id interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{id}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetZoneByID", reflect.TypeOf((*MockZoneServiceIface)(nil).GetZoneByID), varargs...)
}

// GetZoneByName mocks base method.
func (m *MockZoneServiceIface) GetZoneByName(name string, opts ...OptionFunc) (*Zone, int, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{name}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetZoneByName", varargs...)
	ret0, _ := ret[0].(*Zone)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetZoneByName indicates an expected call of GetZoneByName.
func (mr *MockZoneServiceIfaceMockRecorder) GetZoneByName(name interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{name}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetZoneByName", reflect.TypeOf((*MockZoneServiceIface)(nil).GetZoneByName), varargs...)
}

// GetZoneID mocks base method.
func (m *MockZoneServiceIface) GetZoneID(name string, opts ...OptionFunc) (string, int, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{name}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetZoneID", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetZoneID indicates an expected call of GetZoneID.
func (mr *MockZoneServiceIfaceMockRecorder) GetZoneID(name interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{name}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetZoneID", reflect.TypeOf((*MockZoneServiceIface)(nil).GetZoneID), varargs...)
}

// ListDedicatedZones mocks base method.
func (m *MockZoneServiceIface) ListDedicatedZones(p *ListDedicatedZonesParams) (*ListDedicatedZonesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListDedicatedZones", p)
	ret0, _ := ret[0].(*ListDedicatedZonesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDedicatedZones indicates an expected call of ListDedicatedZones.
func (mr *MockZoneServiceIfaceMockRecorder) ListDedicatedZones(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDedicatedZones", reflect.TypeOf((*MockZoneServiceIface)(nil).ListDedicatedZones), p)
}

// ListZones mocks base method.
func (m *MockZoneServiceIface) ListZones(p *ListZonesParams) (*ListZonesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListZones", p)
	ret0, _ := ret[0].(*ListZonesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListZones indicates an expected call of ListZones.
func (mr *MockZoneServiceIfaceMockRecorder) ListZones(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListZones", reflect.TypeOf((*MockZoneServiceIface)(nil).ListZones), p)
}

// NewCreateZoneParams mocks base method.
func (m *MockZoneServiceIface) NewCreateZoneParams(dns1, internaldns1, name, networktype string) *CreateZoneParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewCreateZoneParams", dns1, internaldns1, name, networktype)
	ret0, _ := ret[0].(*CreateZoneParams)
	return ret0
}

// NewCreateZoneParams indicates an expected call of NewCreateZoneParams.
func (mr *MockZoneServiceIfaceMockRecorder) NewCreateZoneParams(dns1, internaldns1, name, networktype interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewCreateZoneParams", reflect.TypeOf((*MockZoneServiceIface)(nil).NewCreateZoneParams), dns1, internaldns1, name, networktype)
}

// NewDedicateZoneParams mocks base method.
func (m *MockZoneServiceIface) NewDedicateZoneParams(domainid, zoneid string) *DedicateZoneParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewDedicateZoneParams", domainid, zoneid)
	ret0, _ := ret[0].(*DedicateZoneParams)
	return ret0
}

// NewDedicateZoneParams indicates an expected call of NewDedicateZoneParams.
func (mr *MockZoneServiceIfaceMockRecorder) NewDedicateZoneParams(domainid, zoneid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewDedicateZoneParams", reflect.TypeOf((*MockZoneServiceIface)(nil).NewDedicateZoneParams), domainid, zoneid)
}

// NewDeleteZoneParams mocks base method.
func (m *MockZoneServiceIface) NewDeleteZoneParams(id string) *DeleteZoneParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewDeleteZoneParams", id)
	ret0, _ := ret[0].(*DeleteZoneParams)
	return ret0
}

// NewDeleteZoneParams indicates an expected call of NewDeleteZoneParams.
func (mr *MockZoneServiceIfaceMockRecorder) NewDeleteZoneParams(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewDeleteZoneParams", reflect.TypeOf((*MockZoneServiceIface)(nil).NewDeleteZoneParams), id)
}

// NewDisableHAForZoneParams mocks base method.
func (m *MockZoneServiceIface) NewDisableHAForZoneParams(zoneid string) *DisableHAForZoneParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewDisableHAForZoneParams", zoneid)
	ret0, _ := ret[0].(*DisableHAForZoneParams)
	return ret0
}

// NewDisableHAForZoneParams indicates an expected call of NewDisableHAForZoneParams.
func (mr *MockZoneServiceIfaceMockRecorder) NewDisableHAForZoneParams(zoneid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewDisableHAForZoneParams", reflect.TypeOf((*MockZoneServiceIface)(nil).NewDisableHAForZoneParams), zoneid)
}

// NewDisableOutOfBandManagementForZoneParams mocks base method.
func (m *MockZoneServiceIface) NewDisableOutOfBandManagementForZoneParams(zoneid string) *DisableOutOfBandManagementForZoneParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewDisableOutOfBandManagementForZoneParams", zoneid)
	ret0, _ := ret[0].(*DisableOutOfBandManagementForZoneParams)
	return ret0
}

// NewDisableOutOfBandManagementForZoneParams indicates an expected call of NewDisableOutOfBandManagementForZoneParams.
func (mr *MockZoneServiceIfaceMockRecorder) NewDisableOutOfBandManagementForZoneParams(zoneid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewDisableOutOfBandManagementForZoneParams", reflect.TypeOf((*MockZoneServiceIface)(nil).NewDisableOutOfBandManagementForZoneParams), zoneid)
}

// NewEnableHAForZoneParams mocks base method.
func (m *MockZoneServiceIface) NewEnableHAForZoneParams(zoneid string) *EnableHAForZoneParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewEnableHAForZoneParams", zoneid)
	ret0, _ := ret[0].(*EnableHAForZoneParams)
	return ret0
}

// NewEnableHAForZoneParams indicates an expected call of NewEnableHAForZoneParams.
func (mr *MockZoneServiceIfaceMockRecorder) NewEnableHAForZoneParams(zoneid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewEnableHAForZoneParams", reflect.TypeOf((*MockZoneServiceIface)(nil).NewEnableHAForZoneParams), zoneid)
}

// NewEnableOutOfBandManagementForZoneParams mocks base method.
func (m *MockZoneServiceIface) NewEnableOutOfBandManagementForZoneParams(zoneid string) *EnableOutOfBandManagementForZoneParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewEnableOutOfBandManagementForZoneParams", zoneid)
	ret0, _ := ret[0].(*EnableOutOfBandManagementForZoneParams)
	return ret0
}

// NewEnableOutOfBandManagementForZoneParams indicates an expected call of NewEnableOutOfBandManagementForZoneParams.
func (mr *MockZoneServiceIfaceMockRecorder) NewEnableOutOfBandManagementForZoneParams(zoneid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewEnableOutOfBandManagementForZoneParams", reflect.TypeOf((*MockZoneServiceIface)(nil).NewEnableOutOfBandManagementForZoneParams), zoneid)
}

// NewListDedicatedZonesParams mocks base method.
func (m *MockZoneServiceIface) NewListDedicatedZonesParams() *ListDedicatedZonesParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListDedicatedZonesParams")
	ret0, _ := ret[0].(*ListDedicatedZonesParams)
	return ret0
}

// NewListDedicatedZonesParams indicates an expected call of NewListDedicatedZonesParams.
func (mr *MockZoneServiceIfaceMockRecorder) NewListDedicatedZonesParams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListDedicatedZonesParams", reflect.TypeOf((*MockZoneServiceIface)(nil).NewListDedicatedZonesParams))
}

// NewListZonesParams mocks base method.
func (m *MockZoneServiceIface) NewListZonesParams() *ListZonesParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListZonesParams")
	ret0, _ := ret[0].(*ListZonesParams)
	return ret0
}

// NewListZonesParams indicates an expected call of NewListZonesParams.
func (mr *MockZoneServiceIfaceMockRecorder) NewListZonesParams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListZonesParams", reflect.TypeOf((*MockZoneServiceIface)(nil).NewListZonesParams))
}

// NewReleaseDedicatedZoneParams mocks base method.
func (m *MockZoneServiceIface) NewReleaseDedicatedZoneParams(zoneid string) *ReleaseDedicatedZoneParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewReleaseDedicatedZoneParams", zoneid)
	ret0, _ := ret[0].(*ReleaseDedicatedZoneParams)
	return ret0
}

// NewReleaseDedicatedZoneParams indicates an expected call of NewReleaseDedicatedZoneParams.
func (mr *MockZoneServiceIfaceMockRecorder) NewReleaseDedicatedZoneParams(zoneid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewReleaseDedicatedZoneParams", reflect.TypeOf((*MockZoneServiceIface)(nil).NewReleaseDedicatedZoneParams), zoneid)
}

// NewUpdateZoneParams mocks base method.
func (m *MockZoneServiceIface) NewUpdateZoneParams(id string) *UpdateZoneParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewUpdateZoneParams", id)
	ret0, _ := ret[0].(*UpdateZoneParams)
	return ret0
}

// NewUpdateZoneParams indicates an expected call of NewUpdateZoneParams.
func (mr *MockZoneServiceIfaceMockRecorder) NewUpdateZoneParams(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewUpdateZoneParams", reflect.TypeOf((*MockZoneServiceIface)(nil).NewUpdateZoneParams), id)
}

// ReleaseDedicatedZone mocks base method.
func (m *MockZoneServiceIface) ReleaseDedicatedZone(p *ReleaseDedicatedZoneParams) (*ReleaseDedicatedZoneResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReleaseDedicatedZone", p)
	ret0, _ := ret[0].(*ReleaseDedicatedZoneResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReleaseDedicatedZone indicates an expected call of ReleaseDedicatedZone.
func (mr *MockZoneServiceIfaceMockRecorder) ReleaseDedicatedZone(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReleaseDedicatedZone", reflect.TypeOf((*MockZoneServiceIface)(nil).ReleaseDedicatedZone), p)
}

// UpdateZone mocks base method.
func (m *MockZoneServiceIface) UpdateZone(p *UpdateZoneParams) (*UpdateZoneResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateZone", p)
	ret0, _ := ret[0].(*UpdateZoneResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateZone indicates an expected call of UpdateZone.
func (mr *MockZoneServiceIfaceMockRecorder) UpdateZone(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateZone", reflect.TypeOf((*MockZoneServiceIface)(nil).UpdateZone), p)
}