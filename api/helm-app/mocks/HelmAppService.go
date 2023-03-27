// Code generated by mockery v2.18.0. DO NOT EDIT.

package mocks

import (
	context "context"

	client "github.com/devtron-labs/devtron/api/helm-app"

	http "net/http"

	mock "github.com/stretchr/testify/mock"

	openapi "github.com/devtron-labs/devtron/api/helm-app/openapiClient"

	openapiClient "github.com/devtron-labs/devtron/api/openapi/openapiClient"
)

// HelmAppService is an autogenerated mock type for the HelmAppService type
type HelmAppService struct {
	mock.Mock
}

// DecodeAppId provides a mock function with given fields: appId
func (_m *HelmAppService) DecodeAppId(appId string) (*client.AppIdentifier, error) {
	ret := _m.Called(appId)

	var r0 *client.AppIdentifier
	if rf, ok := ret.Get(0).(func(string) *client.AppIdentifier); ok {
		r0 = rf(appId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.AppIdentifier)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(appId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteApplication provides a mock function with given fields: ctx, app
func (_m *HelmAppService) DeleteApplication(ctx context.Context, app *client.AppIdentifier) (*openapi.UninstallReleaseResponse, error) {
	ret := _m.Called(ctx, app)

	var r0 *openapi.UninstallReleaseResponse
	if rf, ok := ret.Get(0).(func(context.Context, *client.AppIdentifier) *openapi.UninstallReleaseResponse); ok {
		r0 = rf(ctx, app)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*openapi.UninstallReleaseResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *client.AppIdentifier) error); ok {
		r1 = rf(ctx, app)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EncodeAppId provides a mock function with given fields: appIdentifier
func (_m *HelmAppService) EncodeAppId(appIdentifier *client.AppIdentifier) string {
	ret := _m.Called(appIdentifier)

	var r0 string
	if rf, ok := ret.Get(0).(func(*client.AppIdentifier) string); ok {
		r0 = rf(appIdentifier)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetApplicationDetail provides a mock function with given fields: ctx, app
func (_m *HelmAppService) GetApplicationDetail(ctx context.Context, app *client.AppIdentifier) (*client.AppDetail, error) {
	ret := _m.Called(ctx, app)

	var r0 *client.AppDetail
	if rf, ok := ret.Get(0).(func(context.Context, *client.AppIdentifier) *client.AppDetail); ok {
		r0 = rf(ctx, app)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.AppDetail)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *client.AppIdentifier) error); ok {
		r1 = rf(ctx, app)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetApplicationDetailWithFilter provides a mock function with given fields: ctx, app, resourceTreeFilter
func (_m *HelmAppService) GetApplicationDetailWithFilter(ctx context.Context, app *client.AppIdentifier, resourceTreeFilter *client.ResourceTreeFilter) (*client.AppDetail, error) {
	ret := _m.Called(ctx, app, resourceTreeFilter)

	var r0 *client.AppDetail
	if rf, ok := ret.Get(0).(func(context.Context, *client.AppIdentifier, *client.ResourceTreeFilter) *client.AppDetail); ok {
		r0 = rf(ctx, app, resourceTreeFilter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.AppDetail)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *client.AppIdentifier, *client.ResourceTreeFilter) error); ok {
		r1 = rf(ctx, app, resourceTreeFilter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetApplicationStatus provides a mock function with given fields: ctx, app
func (_m *HelmAppService) GetApplicationStatus(ctx context.Context, app *client.AppIdentifier) (string, error) {
	ret := _m.Called(ctx, app)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, *client.AppIdentifier) string); ok {
		r0 = rf(ctx, app)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *client.AppIdentifier) error); ok {
		r1 = rf(ctx, app)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetClusterConf provides a mock function with given fields: clusterId
func (_m *HelmAppService) GetClusterConf(clusterId int) (*client.ClusterConfig, error) {
	ret := _m.Called(clusterId)

	var r0 *client.ClusterConfig
	if rf, ok := ret.Get(0).(func(int) *client.ClusterConfig); ok {
		r0 = rf(clusterId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.ClusterConfig)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(clusterId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDeploymentDetail provides a mock function with given fields: ctx, app, version
func (_m *HelmAppService) GetDeploymentDetail(ctx context.Context, app *client.AppIdentifier, version int32) (*openapi.HelmAppDeploymentManifestDetail, error) {
	ret := _m.Called(ctx, app, version)

	var r0 *openapi.HelmAppDeploymentManifestDetail
	if rf, ok := ret.Get(0).(func(context.Context, *client.AppIdentifier, int32) *openapi.HelmAppDeploymentManifestDetail); ok {
		r0 = rf(ctx, app, version)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*openapi.HelmAppDeploymentManifestDetail)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *client.AppIdentifier, int32) error); ok {
		r1 = rf(ctx, app, version)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDeploymentHistory provides a mock function with given fields: ctx, app
func (_m *HelmAppService) GetDeploymentHistory(ctx context.Context, app *client.AppIdentifier) (*client.HelmAppDeploymentHistory, error) {
	ret := _m.Called(ctx, app)

	var r0 *client.HelmAppDeploymentHistory
	if rf, ok := ret.Get(0).(func(context.Context, *client.AppIdentifier) *client.HelmAppDeploymentHistory); ok {
		r0 = rf(ctx, app)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.HelmAppDeploymentHistory)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *client.AppIdentifier) error); ok {
		r1 = rf(ctx, app)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDesiredManifest provides a mock function with given fields: ctx, app, resource
func (_m *HelmAppService) GetDesiredManifest(ctx context.Context, app *client.AppIdentifier, resource *openapi.ResourceIdentifier) (*openapi.DesiredManifestResponse, error) {
	ret := _m.Called(ctx, app, resource)

	var r0 *openapi.DesiredManifestResponse
	if rf, ok := ret.Get(0).(func(context.Context, *client.AppIdentifier, *openapi.ResourceIdentifier) *openapi.DesiredManifestResponse); ok {
		r0 = rf(ctx, app, resource)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*openapi.DesiredManifestResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *client.AppIdentifier, *openapi.ResourceIdentifier) error); ok {
		r1 = rf(ctx, app, resource)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDevtronHelmAppIdentifier provides a mock function with given fields:
func (_m *HelmAppService) GetDevtronHelmAppIdentifier() *client.AppIdentifier {
	ret := _m.Called()

	var r0 *client.AppIdentifier
	if rf, ok := ret.Get(0).(func() *client.AppIdentifier); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.AppIdentifier)
		}
	}

	return r0
}

// GetNotes provides a mock function with given fields: ctx, request
func (_m *HelmAppService) GetNotes(ctx context.Context, request *client.InstallReleaseRequest) (string, error) {
	ret := _m.Called(ctx, request)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, *client.InstallReleaseRequest) string); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *client.InstallReleaseRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetValuesYaml provides a mock function with given fields: ctx, app
func (_m *HelmAppService) GetValuesYaml(ctx context.Context, app *client.AppIdentifier) (*client.ReleaseInfo, error) {
	ret := _m.Called(ctx, app)

	var r0 *client.ReleaseInfo
	if rf, ok := ret.Get(0).(func(context.Context, *client.AppIdentifier) *client.ReleaseInfo); ok {
		r0 = rf(ctx, app)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.ReleaseInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *client.AppIdentifier) error); ok {
		r1 = rf(ctx, app)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HibernateApplication provides a mock function with given fields: ctx, app, hibernateRequest
func (_m *HelmAppService) HibernateApplication(ctx context.Context, app *client.AppIdentifier, hibernateRequest *openapi.HibernateRequest) ([]*openapi.HibernateStatus, error) {
	ret := _m.Called(ctx, app, hibernateRequest)

	var r0 []*openapi.HibernateStatus
	if rf, ok := ret.Get(0).(func(context.Context, *client.AppIdentifier, *openapi.HibernateRequest) []*openapi.HibernateStatus); ok {
		r0 = rf(ctx, app, hibernateRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*openapi.HibernateStatus)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *client.AppIdentifier, *openapi.HibernateRequest) error); ok {
		r1 = rf(ctx, app, hibernateRequest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InstallRelease provides a mock function with given fields: ctx, clusterId, installReleaseRequest
func (_m *HelmAppService) InstallRelease(ctx context.Context, clusterId int, installReleaseRequest *client.InstallReleaseRequest) (*client.InstallReleaseResponse, error) {
	ret := _m.Called(ctx, clusterId, installReleaseRequest)

	var r0 *client.InstallReleaseResponse
	if rf, ok := ret.Get(0).(func(context.Context, int, *client.InstallReleaseRequest) *client.InstallReleaseResponse); ok {
		r0 = rf(ctx, clusterId, installReleaseRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.InstallReleaseResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, *client.InstallReleaseRequest) error); ok {
		r1 = rf(ctx, clusterId, installReleaseRequest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsReleaseInstalled provides a mock function with given fields: ctx, app
func (_m *HelmAppService) IsReleaseInstalled(ctx context.Context, app *client.AppIdentifier) (bool, error) {
	ret := _m.Called(ctx, app)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, *client.AppIdentifier) bool); ok {
		r0 = rf(ctx, app)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *client.AppIdentifier) error); ok {
		r1 = rf(ctx, app)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListHelmApplications provides a mock function with given fields: ctx, clusterIds, w, token, helmAuth
func (_m *HelmAppService) ListHelmApplications(ctx context.Context, clusterIds []int, w http.ResponseWriter, token string, helmAuth func(string, string) bool) {
	_m.Called(ctx, clusterIds, w, token, helmAuth)
}

// RollbackRelease provides a mock function with given fields: ctx, app, version
func (_m *HelmAppService) RollbackRelease(ctx context.Context, app *client.AppIdentifier, version int32) (bool, error) {
	ret := _m.Called(ctx, app, version)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, *client.AppIdentifier, int32) bool); ok {
		r0 = rf(ctx, app, version)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *client.AppIdentifier, int32) error); ok {
		r1 = rf(ctx, app, version)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TemplateChart provides a mock function with given fields: ctx, templateChartRequest
func (_m *HelmAppService) TemplateChart(ctx context.Context, templateChartRequest *openapiClient.TemplateChartRequest) (*openapiClient.TemplateChartResponse, error) {
	ret := _m.Called(ctx, templateChartRequest)

	var r0 *openapiClient.TemplateChartResponse
	if rf, ok := ret.Get(0).(func(context.Context, *openapiClient.TemplateChartRequest) *openapiClient.TemplateChartResponse); ok {
		r0 = rf(ctx, templateChartRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*openapiClient.TemplateChartResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *openapiClient.TemplateChartRequest) error); ok {
		r1 = rf(ctx, templateChartRequest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UnHibernateApplication provides a mock function with given fields: ctx, app, hibernateRequest
func (_m *HelmAppService) UnHibernateApplication(ctx context.Context, app *client.AppIdentifier, hibernateRequest *openapi.HibernateRequest) ([]*openapi.HibernateStatus, error) {
	ret := _m.Called(ctx, app, hibernateRequest)

	var r0 []*openapi.HibernateStatus
	if rf, ok := ret.Get(0).(func(context.Context, *client.AppIdentifier, *openapi.HibernateRequest) []*openapi.HibernateStatus); ok {
		r0 = rf(ctx, app, hibernateRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*openapi.HibernateStatus)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *client.AppIdentifier, *openapi.HibernateRequest) error); ok {
		r1 = rf(ctx, app, hibernateRequest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateApplication provides a mock function with given fields: ctx, app, request
func (_m *HelmAppService) UpdateApplication(ctx context.Context, app *client.AppIdentifier, request *openapi.UpdateReleaseRequest) (*openapi.UpdateReleaseResponse, error) {
	ret := _m.Called(ctx, app, request)

	var r0 *openapi.UpdateReleaseResponse
	if rf, ok := ret.Get(0).(func(context.Context, *client.AppIdentifier, *openapi.UpdateReleaseRequest) *openapi.UpdateReleaseResponse); ok {
		r0 = rf(ctx, app, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*openapi.UpdateReleaseResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *client.AppIdentifier, *openapi.UpdateReleaseRequest) error); ok {
		r1 = rf(ctx, app, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateApplicationWithChartInfo provides a mock function with given fields: ctx, clusterId, updateReleaseRequest
func (_m *HelmAppService) UpdateApplicationWithChartInfo(ctx context.Context, clusterId int, updateReleaseRequest *client.InstallReleaseRequest) (*openapi.UpdateReleaseResponse, error) {
	ret := _m.Called(ctx, clusterId, updateReleaseRequest)

	var r0 *openapi.UpdateReleaseResponse
	if rf, ok := ret.Get(0).(func(context.Context, int, *client.InstallReleaseRequest) *openapi.UpdateReleaseResponse); ok {
		r0 = rf(ctx, clusterId, updateReleaseRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*openapi.UpdateReleaseResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, *client.InstallReleaseRequest) error); ok {
		r1 = rf(ctx, clusterId, updateReleaseRequest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateApplicationWithChartInfoWithExtraValues provides a mock function with given fields: ctx, appIdentifier, chartRepository, extraValues, extraValuesYamlUrl, useLatestChartVersion
func (_m *HelmAppService) UpdateApplicationWithChartInfoWithExtraValues(ctx context.Context, appIdentifier *client.AppIdentifier, chartRepository *client.ChartRepository, extraValues map[string]interface{}, extraValuesYamlUrl string, useLatestChartVersion bool) (*openapi.UpdateReleaseResponse, error) {
	ret := _m.Called(ctx, appIdentifier, chartRepository, extraValues, extraValuesYamlUrl, useLatestChartVersion)

	var r0 *openapi.UpdateReleaseResponse
	if rf, ok := ret.Get(0).(func(context.Context, *client.AppIdentifier, *client.ChartRepository, map[string]interface{}, string, bool) *openapi.UpdateReleaseResponse); ok {
		r0 = rf(ctx, appIdentifier, chartRepository, extraValues, extraValuesYamlUrl, useLatestChartVersion)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*openapi.UpdateReleaseResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *client.AppIdentifier, *client.ChartRepository, map[string]interface{}, string, bool) error); ok {
		r1 = rf(ctx, appIdentifier, chartRepository, extraValues, extraValuesYamlUrl, useLatestChartVersion)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewHelmAppService interface {
	mock.TestingT
	Cleanup(func())
}

// NewHelmAppService creates a new instance of HelmAppService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewHelmAppService(t mockConstructorTestingTNewHelmAppService) *HelmAppService {
	mock := &HelmAppService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}