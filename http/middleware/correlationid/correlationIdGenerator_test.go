// Code generated by mockery v1.0.0. DO NOT EDIT.
package correlationid_test

import mock "github.com/stretchr/testify/mock"

// correlationIdGenerator is an autogenerated mock type for the correlationIdGenerator type
type correlationIdGenerator struct {
	mock.Mock
}

// Generate provides a mock function with given fields:
func (_m *correlationIdGenerator) Generate() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}