// Code generated by mockery v2.10.6. DO NOT EDIT.

package mock

import (
	exd "github.com/odpf/optimus/ext/exd"
	mock "github.com/stretchr/testify/mock"
)

// Installer is an autogenerated mock type for the Installer type
type Installer struct {
	mock.Mock
}

// Install provides a mock function with given fields: asset, metadata
func (_m *Installer) Install(asset []byte, metadata *exd.Metadata) error {
	ret := _m.Called(asset, metadata)

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte, *exd.Metadata) error); ok {
		r0 = rf(asset, metadata)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Prepare provides a mock function with given fields: metadata
func (_m *Installer) Prepare(metadata *exd.Metadata) error {
	ret := _m.Called(metadata)

	var r0 error
	if rf, ok := ret.Get(0).(func(*exd.Metadata) error); ok {
		r0 = rf(metadata)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}