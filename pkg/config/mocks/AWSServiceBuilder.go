// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import config "github.com/Optum/dce/pkg/config"
import mock "github.com/stretchr/testify/mock"

// AWSServiceBuilder is an autogenerated mock type for the AWSServiceBuilder type
type AWSServiceBuilder struct {
	mock.Mock
}

// Build provides a mock function with given fields:
func (_m *AWSServiceBuilder) Build() (*config.DefaultConfigurationBuilder, error) {
	ret := _m.Called()

	var r0 *config.DefaultConfigurationBuilder
	if rf, ok := ret.Get(0).(func() *config.DefaultConfigurationBuilder); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*config.DefaultConfigurationBuilder)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WithCodeBuild provides a mock function with given fields:
func (_m *AWSServiceBuilder) WithCodeBuild() *config.DefaultAWSServiceBuilder {
	ret := _m.Called()

	var r0 *config.DefaultAWSServiceBuilder
	if rf, ok := ret.Get(0).(func() *config.DefaultAWSServiceBuilder); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*config.DefaultAWSServiceBuilder)
		}
	}

	return r0
}

// WithCognito provides a mock function with given fields:
func (_m *AWSServiceBuilder) WithCognito() *config.DefaultAWSServiceBuilder {
	ret := _m.Called()

	var r0 *config.DefaultAWSServiceBuilder
	if rf, ok := ret.Get(0).(func() *config.DefaultAWSServiceBuilder); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*config.DefaultAWSServiceBuilder)
		}
	}

	return r0
}

// WithDynamoDB provides a mock function with given fields:
func (_m *AWSServiceBuilder) WithDynamoDB() *config.DefaultAWSServiceBuilder {
	ret := _m.Called()

	var r0 *config.DefaultAWSServiceBuilder
	if rf, ok := ret.Get(0).(func() *config.DefaultAWSServiceBuilder); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*config.DefaultAWSServiceBuilder)
		}
	}

	return r0
}

// WithS3 provides a mock function with given fields:
func (_m *AWSServiceBuilder) WithS3() *config.DefaultAWSServiceBuilder {
	ret := _m.Called()

	var r0 *config.DefaultAWSServiceBuilder
	if rf, ok := ret.Get(0).(func() *config.DefaultAWSServiceBuilder); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*config.DefaultAWSServiceBuilder)
		}
	}

	return r0
}

// WithSNS provides a mock function with given fields:
func (_m *AWSServiceBuilder) WithSNS() *config.DefaultAWSServiceBuilder {
	ret := _m.Called()

	var r0 *config.DefaultAWSServiceBuilder
	if rf, ok := ret.Get(0).(func() *config.DefaultAWSServiceBuilder); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*config.DefaultAWSServiceBuilder)
		}
	}

	return r0
}

// WithSQS provides a mock function with given fields:
func (_m *AWSServiceBuilder) WithSQS() *config.DefaultAWSServiceBuilder {
	ret := _m.Called()

	var r0 *config.DefaultAWSServiceBuilder
	if rf, ok := ret.Get(0).(func() *config.DefaultAWSServiceBuilder); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*config.DefaultAWSServiceBuilder)
		}
	}

	return r0
}

// WithSSM provides a mock function with given fields:
func (_m *AWSServiceBuilder) WithSSM() *config.DefaultAWSServiceBuilder {
	ret := _m.Called()

	var r0 *config.DefaultAWSServiceBuilder
	if rf, ok := ret.Get(0).(func() *config.DefaultAWSServiceBuilder); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*config.DefaultAWSServiceBuilder)
		}
	}

	return r0
}

// WithSTS provides a mock function with given fields:
func (_m *AWSServiceBuilder) WithSTS() *config.DefaultAWSServiceBuilder {
	ret := _m.Called()

	var r0 *config.DefaultAWSServiceBuilder
	if rf, ok := ret.Get(0).(func() *config.DefaultAWSServiceBuilder); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*config.DefaultAWSServiceBuilder)
		}
	}

	return r0
}
