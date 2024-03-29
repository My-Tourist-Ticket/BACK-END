// Code generated by mockery v2.39.2. DO NOT EDIT.

package mocks

import (
	dashboard "my-tourist-ticket/features/dashboard"

	mock "github.com/stretchr/testify/mock"
)

// DashboardData is an autogenerated mock type for the DashboardDataInterface type
type DashboardData struct {
	mock.Mock
}

// GetRecentTransaction provides a mock function with given fields:
func (_m *DashboardData) GetRecentTransaction() ([]dashboard.Booking, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetRecentTransaction")
	}

	var r0 []dashboard.Booking
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]dashboard.Booking, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []dashboard.Booking); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dashboard.Booking)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTopTour provides a mock function with given fields:
func (_m *DashboardData) GetTopTour() ([]dashboard.Tour, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetTopTour")
	}

	var r0 []dashboard.Tour
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]dashboard.Tour, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []dashboard.Tour); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dashboard.Tour)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTotalCustomer provides a mock function with given fields:
func (_m *DashboardData) GetTotalCustomer() (int, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetTotalCustomer")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func() (int, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTotalPengelola provides a mock function with given fields:
func (_m *DashboardData) GetTotalPengelola() (int, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetTotalPengelola")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func() (int, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTotalTour provides a mock function with given fields:
func (_m *DashboardData) GetTotalTour() (int, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetTotalTour")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func() (int, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTotalTransaction provides a mock function with given fields:
func (_m *DashboardData) GetTotalTransaction() (int, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetTotalTransaction")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func() (int, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserRoleById provides a mock function with given fields: userId
func (_m *DashboardData) GetUserRoleById(userId int) (string, error) {
	ret := _m.Called(userId)

	if len(ret) == 0 {
		panic("no return value specified for GetUserRoleById")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (string, error)); ok {
		return rf(userId)
	}
	if rf, ok := ret.Get(0).(func(int) string); ok {
		r0 = rf(userId)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewDashboardData creates a new instance of DashboardData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDashboardData(t interface {
	mock.TestingT
	Cleanup(func())
}) *DashboardData {
	mock := &DashboardData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
