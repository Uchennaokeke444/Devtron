// Code generated by mockery v2.31.4. DO NOT EDIT.

package mocks

import (
	appStoreBean "github.com/devtron-labs/devtron/pkg/appStore/bean"
	chart "k8s.io/helm/pkg/proto/hapi/chart"

	context "context"

	mock "github.com/stretchr/testify/mock"

	util "github.com/devtron-labs/devtron/internal/util"
)

// ChartTemplateService is an autogenerated mock type for the ChartTemplateService type
type ChartTemplateService struct {
	mock.Mock
}

// BuildChart provides a mock function with given fields: ctx, chartMetaData, referenceTemplatePath
func (_m *ChartTemplateService) BuildChart(ctx context.Context, chartMetaData *chart.Metadata, referenceTemplatePath string) (string, error) {
	ret := _m.Called(ctx, chartMetaData, referenceTemplatePath)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *chart.Metadata, string) (string, error)); ok {
		return rf(ctx, chartMetaData, referenceTemplatePath)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *chart.Metadata, string) string); ok {
		r0 = rf(ctx, chartMetaData, referenceTemplatePath)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *chart.Metadata, string) error); ok {
		r1 = rf(ctx, chartMetaData, referenceTemplatePath)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BuildChartProxyForHelmApps provides a mock function with given fields: chartCreateRequest
func (_m *ChartTemplateService) BuildChartProxyForHelmApps(chartCreateRequest *util.ChartCreateRequest) (*util.ChartCreateResponse, error) {
	ret := _m.Called(chartCreateRequest)

	var r0 *util.ChartCreateResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(*util.ChartCreateRequest) (*util.ChartCreateResponse, error)); ok {
		return rf(chartCreateRequest)
	}
	if rf, ok := ret.Get(0).(func(*util.ChartCreateRequest) *util.ChartCreateResponse); ok {
		r0 = rf(chartCreateRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*util.ChartCreateResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(*util.ChartCreateRequest) error); ok {
		r1 = rf(chartCreateRequest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CleanDir provides a mock function with given fields: dir
func (_m *ChartTemplateService) CleanDir(dir string) {
	_m.Called(dir)
}

// CreateAndPushToGitChartProxy provides a mock function with given fields: appStoreName, tmpChartLocation, envName, installAppVersionRequest
func (_m *ChartTemplateService) CreateAndPushToGitChartProxy(appStoreName string, tmpChartLocation string, envName string, installAppVersionRequest *appStoreBean.InstallAppVersionDTO) (*util.ChartGitAttribute, error) {
	ret := _m.Called(appStoreName, tmpChartLocation, envName, installAppVersionRequest)

	var r0 *util.ChartGitAttribute
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, string, *appStoreBean.InstallAppVersionDTO) (*util.ChartGitAttribute, error)); ok {
		return rf(appStoreName, tmpChartLocation, envName, installAppVersionRequest)
	}
	if rf, ok := ret.Get(0).(func(string, string, string, *appStoreBean.InstallAppVersionDTO) *util.ChartGitAttribute); ok {
		r0 = rf(appStoreName, tmpChartLocation, envName, installAppVersionRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*util.ChartGitAttribute)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, string, *appStoreBean.InstallAppVersionDTO) error); ok {
		r1 = rf(appStoreName, tmpChartLocation, envName, installAppVersionRequest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateChartProxy provides a mock function with given fields: chartMetaData, refChartLocation, envName, installAppVersionRequest
func (_m *ChartTemplateService) CreateChartProxy(chartMetaData *chart.Metadata, refChartLocation string, envName string, installAppVersionRequest *appStoreBean.InstallAppVersionDTO) (string, *util.ChartGitAttribute, error) {
	ret := _m.Called(chartMetaData, refChartLocation, envName, installAppVersionRequest)

	var r0 string
	var r1 *util.ChartGitAttribute
	var r2 error
	if rf, ok := ret.Get(0).(func(*chart.Metadata, string, string, *appStoreBean.InstallAppVersionDTO) (string, *util.ChartGitAttribute, error)); ok {
		return rf(chartMetaData, refChartLocation, envName, installAppVersionRequest)
	}
	if rf, ok := ret.Get(0).(func(*chart.Metadata, string, string, *appStoreBean.InstallAppVersionDTO) string); ok {
		r0 = rf(chartMetaData, refChartLocation, envName, installAppVersionRequest)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(*chart.Metadata, string, string, *appStoreBean.InstallAppVersionDTO) *util.ChartGitAttribute); ok {
		r1 = rf(chartMetaData, refChartLocation, envName, installAppVersionRequest)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*util.ChartGitAttribute)
		}
	}

	if rf, ok := ret.Get(2).(func(*chart.Metadata, string, string, *appStoreBean.InstallAppVersionDTO) error); ok {
		r2 = rf(chartMetaData, refChartLocation, envName, installAppVersionRequest)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// CreateGitRepositoryForApp provides a mock function with given fields: gitOpsRepoName, baseTemplateName, version, userId
func (_m *ChartTemplateService) CreateGitRepositoryForApp(gitOpsRepoName string, baseTemplateName string, version string, userId int32) (*util.ChartGitAttribute, error) {
	ret := _m.Called(gitOpsRepoName, baseTemplateName, version, userId)

	var r0 *util.ChartGitAttribute
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, string, int32) (*util.ChartGitAttribute, error)); ok {
		return rf(gitOpsRepoName, baseTemplateName, version, userId)
	}
	if rf, ok := ret.Get(0).(func(string, string, string, int32) *util.ChartGitAttribute); ok {
		r0 = rf(gitOpsRepoName, baseTemplateName, version, userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*util.ChartGitAttribute)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, string, int32) error); ok {
		r1 = rf(gitOpsRepoName, baseTemplateName, version, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateReadmeInGitRepo provides a mock function with given fields: gitOpsRepoName, userId
func (_m *ChartTemplateService) CreateReadmeInGitRepo(gitOpsRepoName string, userId int32) error {
	ret := _m.Called(gitOpsRepoName, userId)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, int32) error); ok {
		r0 = rf(gitOpsRepoName, userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FetchValuesFromReferenceChart provides a mock function with given fields: chartMetaData, refChartLocation, templateName, userId, pipelineStrategyPath
func (_m *ChartTemplateService) FetchValuesFromReferenceChart(chartMetaData *chart.Metadata, refChartLocation string, templateName string, userId int32, pipelineStrategyPath string) (*util.ChartValues, *util.ChartGitAttribute, error) {
	ret := _m.Called(chartMetaData, refChartLocation, templateName, userId, pipelineStrategyPath)

	var r0 *util.ChartValues
	var r1 *util.ChartGitAttribute
	var r2 error
	if rf, ok := ret.Get(0).(func(*chart.Metadata, string, string, int32, string) (*util.ChartValues, *util.ChartGitAttribute, error)); ok {
		return rf(chartMetaData, refChartLocation, templateName, userId, pipelineStrategyPath)
	}
	if rf, ok := ret.Get(0).(func(*chart.Metadata, string, string, int32, string) *util.ChartValues); ok {
		r0 = rf(chartMetaData, refChartLocation, templateName, userId, pipelineStrategyPath)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*util.ChartValues)
		}
	}

	if rf, ok := ret.Get(1).(func(*chart.Metadata, string, string, int32, string) *util.ChartGitAttribute); ok {
		r1 = rf(chartMetaData, refChartLocation, templateName, userId, pipelineStrategyPath)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*util.ChartGitAttribute)
		}
	}

	if rf, ok := ret.Get(2).(func(*chart.Metadata, string, string, int32, string) error); ok {
		r2 = rf(chartMetaData, refChartLocation, templateName, userId, pipelineStrategyPath)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetByteArrayRefChart provides a mock function with given fields: chartMetaData, referenceTemplatePath
func (_m *ChartTemplateService) GetByteArrayRefChart(chartMetaData *chart.Metadata, referenceTemplatePath string) ([]byte, error) {
	ret := _m.Called(chartMetaData, referenceTemplatePath)

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(*chart.Metadata, string) ([]byte, error)); ok {
		return rf(chartMetaData, referenceTemplatePath)
	}
	if rf, ok := ret.Get(0).(func(*chart.Metadata, string) []byte); ok {
		r0 = rf(chartMetaData, referenceTemplatePath)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(*chart.Metadata, string) error); ok {
		r1 = rf(chartMetaData, referenceTemplatePath)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetChartVersion provides a mock function with given fields: location
func (_m *ChartTemplateService) GetChartVersion(location string) (string, error) {
	ret := _m.Called(location)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(location)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(location)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(location)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDir provides a mock function with given fields:
func (_m *ChartTemplateService) GetDir() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetGitOpsRepoName provides a mock function with given fields: appName
func (_m *ChartTemplateService) GetGitOpsRepoName(appName string) string {
	ret := _m.Called(appName)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(appName)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetGitOpsRepoNameFromUrl provides a mock function with given fields: gitRepoUrl
func (_m *ChartTemplateService) GetGitOpsRepoNameFromUrl(gitRepoUrl string) string {
	ret := _m.Called(gitRepoUrl)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(gitRepoUrl)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetUserEmailIdAndNameForGitOpsCommit provides a mock function with given fields: userId
func (_m *ChartTemplateService) GetUserEmailIdAndNameForGitOpsCommit(userId int32) (string, string) {
	ret := _m.Called(userId)

	var r0 string
	var r1 string
	if rf, ok := ret.Get(0).(func(int32) (string, string)); ok {
		return rf(userId)
	}
	if rf, ok := ret.Get(0).(func(int32) string); ok {
		r0 = rf(userId)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(int32) string); ok {
		r1 = rf(userId)
	} else {
		r1 = ret.Get(1).(string)
	}

	return r0, r1
}

// GitPull provides a mock function with given fields: clonedDir, repoUrl, appStoreName
func (_m *ChartTemplateService) GitPull(clonedDir string, repoUrl string, appStoreName string) error {
	ret := _m.Called(clonedDir, repoUrl, appStoreName)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(clonedDir, repoUrl, appStoreName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LoadChartInBytes provides a mock function with given fields: ChartPath, deleteChart, chartName, chartVersion
func (_m *ChartTemplateService) LoadChartInBytes(ChartPath string, deleteChart bool, chartName string, chartVersion string) ([]byte, string, error) {
	ret := _m.Called(ChartPath, deleteChart, chartName, chartVersion)

	var r0 []byte
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(string, bool, string, string) ([]byte, string, error)); ok {
		return rf(ChartPath, deleteChart, chartName, chartVersion)
	}
	if rf, ok := ret.Get(0).(func(string, bool, string, string) []byte); ok {
		r0 = rf(ChartPath, deleteChart, chartName, chartVersion)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(string, bool, string, string) string); ok {
		r1 = rf(ChartPath, deleteChart, chartName, chartVersion)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(string, bool, string, string) error); ok {
		r2 = rf(ChartPath, deleteChart, chartName, chartVersion)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// PushChartToGitRepo provides a mock function with given fields: gitOpsRepoName, referenceTemplate, version, tempReferenceTemplateDir, repoUrl, userId
func (_m *ChartTemplateService) PushChartToGitRepo(gitOpsRepoName string, referenceTemplate string, version string, tempReferenceTemplateDir string, repoUrl string, userId int32) error {
	ret := _m.Called(gitOpsRepoName, referenceTemplate, version, tempReferenceTemplateDir, repoUrl, userId)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, string, string, int32) error); ok {
		r0 = rf(gitOpsRepoName, referenceTemplate, version, tempReferenceTemplateDir, repoUrl, userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateGitRepoUrlInCharts provides a mock function with given fields: appId, chartGitAttribute, userId
func (_m *ChartTemplateService) UpdateGitRepoUrlInCharts(appId int, chartGitAttribute *util.ChartGitAttribute, userId int32) error {
	ret := _m.Called(appId, chartGitAttribute, userId)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, *util.ChartGitAttribute, int32) error); ok {
		r0 = rf(appId, chartGitAttribute, userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewChartTemplateService creates a new instance of ChartTemplateService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewChartTemplateService(t interface {
	mock.TestingT
	Cleanup(func())
}) *ChartTemplateService {
	mock := &ChartTemplateService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
