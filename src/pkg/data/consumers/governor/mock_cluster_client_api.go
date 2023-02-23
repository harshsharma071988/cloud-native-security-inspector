// Code generated by mockery v2.20.2. DO NOT EDIT.

package consumers

import (
	context "context"
	http "net/http"

	mock "github.com/stretchr/testify/mock"

	openapi "github.com/vmware-tanzu/cloud-native-security-inspector/src/pkg/data/consumers/governor/go-client"
)

// ClustersApi is an autogenerated mock type for the ClustersApi type
type ClustersApi struct {
	mock.Mock
}

// FetchAgentConfig provides a mock function with given fields: ctx, clusterId
func (_m *ClustersApi) FetchAgentConfig(ctx context.Context, clusterId string) openapi.ApiFetchAgentConfigRequest {
	ret := _m.Called(ctx, clusterId)

	var r0 openapi.ApiFetchAgentConfigRequest
	if rf, ok := ret.Get(0).(func(context.Context, string) openapi.ApiFetchAgentConfigRequest); ok {
		r0 = rf(ctx, clusterId)
	} else {
		r0 = ret.Get(0).(openapi.ApiFetchAgentConfigRequest)
	}

	return r0
}

// FetchAgentConfigExecute provides a mock function with given fields: r
func (_m *ClustersApi) FetchAgentConfigExecute(r openapi.ApiFetchAgentConfigRequest) (string, *http.Response, error) {
	ret := _m.Called(r)

	var r0 string
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(openapi.ApiFetchAgentConfigRequest) (string, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(openapi.ApiFetchAgentConfigRequest) string); ok {
		r0 = rf(r)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(openapi.ApiFetchAgentConfigRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(openapi.ApiFetchAgentConfigRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetCluster provides a mock function with given fields: ctx, clusterId
func (_m *ClustersApi) GetCluster(ctx context.Context, clusterId string) openapi.ApiGetClusterRequest {
	ret := _m.Called(ctx, clusterId)

	var r0 openapi.ApiGetClusterRequest
	if rf, ok := ret.Get(0).(func(context.Context, string) openapi.ApiGetClusterRequest); ok {
		r0 = rf(ctx, clusterId)
	} else {
		r0 = ret.Get(0).(openapi.ApiGetClusterRequest)
	}

	return r0
}

// GetClusterExecute provides a mock function with given fields: r
func (_m *ClustersApi) GetClusterExecute(r openapi.ApiGetClusterRequest) (*openapi.KubernetesClusterDetailedResponse, *http.Response, error) {
	ret := _m.Called(r)

	var r0 *openapi.KubernetesClusterDetailedResponse
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(openapi.ApiGetClusterRequest) (*openapi.KubernetesClusterDetailedResponse, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(openapi.ApiGetClusterRequest) *openapi.KubernetesClusterDetailedResponse); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*openapi.KubernetesClusterDetailedResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(openapi.ApiGetClusterRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(openapi.ApiGetClusterRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetClusters provides a mock function with given fields: ctx
func (_m *ClustersApi) GetClusters(ctx context.Context) openapi.ApiGetClustersRequest {
	ret := _m.Called(ctx)

	var r0 openapi.ApiGetClustersRequest
	if rf, ok := ret.Get(0).(func(context.Context) openapi.ApiGetClustersRequest); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(openapi.ApiGetClustersRequest)
	}

	return r0
}

// GetClustersExecute provides a mock function with given fields: r
func (_m *ClustersApi) GetClustersExecute(r openapi.ApiGetClustersRequest) ([]openapi.KubernetesClusterResponse, *http.Response, error) {
	ret := _m.Called(r)

	var r0 []openapi.KubernetesClusterResponse
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(openapi.ApiGetClustersRequest) ([]openapi.KubernetesClusterResponse, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(openapi.ApiGetClustersRequest) []openapi.KubernetesClusterResponse); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]openapi.KubernetesClusterResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(openapi.ApiGetClustersRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(openapi.ApiGetClustersRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// RegisterCluster provides a mock function with given fields: ctx
func (_m *ClustersApi) RegisterCluster(ctx context.Context) openapi.ApiRegisterClusterRequest {
	ret := _m.Called(ctx)

	var r0 openapi.ApiRegisterClusterRequest
	if rf, ok := ret.Get(0).(func(context.Context) openapi.ApiRegisterClusterRequest); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(openapi.ApiRegisterClusterRequest)
	}

	return r0
}

// RegisterClusterExecute provides a mock function with given fields: r
func (_m *ClustersApi) RegisterClusterExecute(r openapi.ApiRegisterClusterRequest) (*http.Response, error) {
	ret := _m.Called(r)

	var r0 *http.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(openapi.ApiRegisterClusterRequest) (*http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(openapi.ApiRegisterClusterRequest) *http.Response); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(openapi.ApiRegisterClusterRequest) error); ok {
		r1 = rf(r)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UnregisterCluster provides a mock function with given fields: ctx, clusterId
func (_m *ClustersApi) UnregisterCluster(ctx context.Context, clusterId string) openapi.ApiUnregisterClusterRequest {
	ret := _m.Called(ctx, clusterId)

	var r0 openapi.ApiUnregisterClusterRequest
	if rf, ok := ret.Get(0).(func(context.Context, string) openapi.ApiUnregisterClusterRequest); ok {
		r0 = rf(ctx, clusterId)
	} else {
		r0 = ret.Get(0).(openapi.ApiUnregisterClusterRequest)
	}

	return r0
}

// UnregisterClusterExecute provides a mock function with given fields: r
func (_m *ClustersApi) UnregisterClusterExecute(r openapi.ApiUnregisterClusterRequest) (*http.Response, error) {
	ret := _m.Called(r)

	var r0 *http.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(openapi.ApiUnregisterClusterRequest) (*http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(openapi.ApiUnregisterClusterRequest) *http.Response); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(openapi.ApiUnregisterClusterRequest) error); ok {
		r1 = rf(r)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTelemetry provides a mock function with given fields: ctx, clusterId
func (_m *ClustersApi) UpdateTelemetry(ctx context.Context, clusterId string) openapi.ApiUpdateTelemetryRequest {
	ret := _m.Called(ctx, clusterId)

	var r0 openapi.ApiUpdateTelemetryRequest
	if rf, ok := ret.Get(0).(func(context.Context, string) openapi.ApiUpdateTelemetryRequest); ok {
		r0 = rf(ctx, clusterId)
	} else {
		r0 = ret.Get(0).(openapi.ApiUpdateTelemetryRequest)
	}

	return r0
}

// UpdateTelemetryExecute provides a mock function with given fields: r
func (_m *ClustersApi) UpdateTelemetryExecute(r openapi.ApiUpdateTelemetryRequest) (*http.Response, error) {
	ret := _m.Called(r)

	var r0 *http.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(openapi.ApiUpdateTelemetryRequest) (*http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(openapi.ApiUpdateTelemetryRequest) *http.Response); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(openapi.ApiUpdateTelemetryRequest) error); ok {
		r1 = rf(r)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewClustersApi interface {
	mock.TestingT
	Cleanup(func())
}

// NewClustersApi creates a new instance of ClustersApi. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewClustersApi(t mockConstructorTestingTNewClustersApi) *ClustersApi {
	mock := &ClustersApi{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
