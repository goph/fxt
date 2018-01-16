// Code generated by mockery v1.0.0
package mocks

import context "golang.org/x/net/context"

import mock "github.com/stretchr/testify/mock"

// Carrier is an autogenerated mock type for the Carrier type
type Carrier struct {
	mock.Mock
}

// GetCorrelationID provides a mock function with given fields: ctx
func (_m *Carrier) GetCorrelationID(ctx context.Context) (string, bool) {
	ret := _m.Called(ctx)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context) string); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(context.Context) bool); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// SetCorrelationID provides a mock function with given fields: ctx, correlationID
func (_m *Carrier) SetCorrelationID(ctx context.Context, correlationID string) context.Context {
	ret := _m.Called(ctx, correlationID)

	var r0 context.Context
	if rf, ok := ret.Get(0).(func(context.Context, string) context.Context); ok {
		r0 = rf(ctx, correlationID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	return r0
}
