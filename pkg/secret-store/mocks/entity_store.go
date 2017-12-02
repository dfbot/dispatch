// Code generated by mockery v1.0.0
package mocks

import entitystore "github.com/vmware/dispatch/pkg/entity-store"
import mock "github.com/stretchr/testify/mock"

// EntityStore is an autogenerated mock type for the EntityStore type
type EntityStore struct {
	mock.Mock
}

// Add provides a mock function with given fields: entity
func (_m *EntityStore) Add(entity entitystore.Entity) (string, error) {
	ret := _m.Called(entity)

	var r0 string
	if rf, ok := ret.Get(0).(func(entitystore.Entity) string); ok {
		r0 = rf(entity)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(entitystore.Entity) error); ok {
		r1 = rf(entity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: organizationID, id, entity
func (_m *EntityStore) Delete(organizationID string, id string, entity entitystore.Entity) error {
	ret := _m.Called(organizationID, id, entity)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, entitystore.Entity) error); ok {
		r0 = rf(organizationID, id, entity)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: organizationID, key, entity
func (_m *EntityStore) Get(organizationID string, key string, entity entitystore.Entity) error {
	ret := _m.Called(organizationID, key, entity)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, entitystore.Entity) error); ok {
		r0 = rf(organizationID, key, entity)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// List provides a mock function with given fields: organizationID, filter, entities
func (_m *EntityStore) List(organizationID string, filter entitystore.Filter, entities interface{}) error {
	ret := _m.Called(organizationID, filter, entities)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, entitystore.Filter, interface{}) error); ok {
		r0 = rf(organizationID, filter, entities)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: lastRevision, entity
func (_m *EntityStore) Update(lastRevision uint64, entity entitystore.Entity) (int64, error) {
	ret := _m.Called(lastRevision, entity)

	var r0 int64
	if rf, ok := ret.Get(0).(func(uint64, entitystore.Entity) int64); ok {
		r0 = rf(lastRevision, entity)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64, entitystore.Entity) error); ok {
		r1 = rf(lastRevision, entity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateWithError provides a mock function with given fields: e, err
func (_m *EntityStore) UpdateWithError(e entitystore.Entity, err error) {
	_m.Called(e, err)
}
