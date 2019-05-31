// Code generated by mockery v1.0.0. DO NOT EDIT.

package core

import mock "github.com/stretchr/testify/mock"

// MockScEngineManager is an autogenerated mock type for the ScEngineManager type
type MockScEngineManager struct {
	mock.Mock
}

// CreateEngine provides a mock function with given fields:
func (_m *MockScEngineManager) CreateEngine() ScEngine {
	ret := _m.Called()

	var r0 ScEngine
	if rf, ok := ret.Get(0).(func() ScEngine); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ScEngine)
		}
	}

	return r0
}

// RunScheduledEvents provides a mock function with given fields: contractUtxo, scStorage
func (_m *MockScEngineManager) RunScheduledEvents(contractUtxo []*UTXO, scStorage *ScState, blkHeight uint64, seed int64) {
	_m.Called(contractUtxo, scStorage, blkHeight, seed)
}
