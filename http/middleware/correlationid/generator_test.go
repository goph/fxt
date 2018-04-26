package correlationid_test

import "github.com/stretchr/testify/mock"

type MockGenerator struct {
	mock.Mock
}

// Generate provides a mock function with given fields:
func (_m *MockGenerator) Generate() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}
