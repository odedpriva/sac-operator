// Code generated by mockery 2.9.0. DO NOT EDIT.

package sac

import (
	dto "bitbucket.org/accezz-io/sac-operator/service/sac/dto"
	mock "github.com/stretchr/testify/mock"

	model "bitbucket.org/accezz-io/sac-operator/model"

	uuid "github.com/google/uuid"
)

// MockSecureAccessCloudClient is an autogenerated mock type for the SecureAccessCloudClient type
type MockSecureAccessCloudClient struct {
	mock.Mock
}

// BindApplicationToSite provides a mock function with given fields: applicationId, siteId
func (_m *MockSecureAccessCloudClient) BindApplicationToSite(applicationId uuid.UUID, siteId string) error {
	ret := _m.Called(applicationId, siteId)

	var r0 error
	if rf, ok := ret.Get(0).(func(uuid.UUID, string) error); ok {
		r0 = rf(applicationId, siteId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateApplication provides a mock function with given fields: applicationDTO
func (_m *MockSecureAccessCloudClient) CreateApplication(applicationDTO *dto.ApplicationDTO) (*dto.ApplicationDTO, error) {
	ret := _m.Called(applicationDTO)

	var r0 *dto.ApplicationDTO
	if rf, ok := ret.Get(0).(func(*dto.ApplicationDTO) *dto.ApplicationDTO); ok {
		r0 = rf(applicationDTO)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.ApplicationDTO)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*dto.ApplicationDTO) error); ok {
		r1 = rf(applicationDTO)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateConnector provides a mock function with given fields: siteDTO, connectorName
func (_m *MockSecureAccessCloudClient) CreateConnector(siteDTO *dto.SiteDTO, connectorName string) (*dto.ConnectorObjects, error) {
	ret := _m.Called(siteDTO, connectorName)

	var r0 *dto.ConnectorObjects
	if rf, ok := ret.Get(0).(func(*dto.SiteDTO, string) *dto.ConnectorObjects); ok {
		r0 = rf(siteDTO, connectorName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.ConnectorObjects)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*dto.SiteDTO, string) error); ok {
		r1 = rf(siteDTO, connectorName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateSite provides a mock function with given fields: siteDTO
func (_m *MockSecureAccessCloudClient) CreateSite(siteDTO *dto.SiteDTO) (*dto.SiteDTO, error) {
	ret := _m.Called(siteDTO)

	var r0 *dto.SiteDTO
	if rf, ok := ret.Get(0).(func(*dto.SiteDTO) *dto.SiteDTO); ok {
		r0 = rf(siteDTO)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.SiteDTO)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*dto.SiteDTO) error); ok {
		r1 = rf(siteDTO)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteApplication provides a mock function with given fields: id
func (_m *MockSecureAccessCloudClient) DeleteApplication(id uuid.UUID) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteConnector provides a mock function with given fields: connectorID
func (_m *MockSecureAccessCloudClient) DeleteConnector(connectorID string) error {
	ret := _m.Called(connectorID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(connectorID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteSite provides a mock function with given fields: id
func (_m *MockSecureAccessCloudClient) DeleteSite(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindApplicationByID provides a mock function with given fields: id
func (_m *MockSecureAccessCloudClient) FindApplicationByID(id uuid.UUID) (*dto.ApplicationDTO, error) {
	ret := _m.Called(id)

	var r0 *dto.ApplicationDTO
	if rf, ok := ret.Get(0).(func(uuid.UUID) *dto.ApplicationDTO); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.ApplicationDTO)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindApplicationByName provides a mock function with given fields: name
func (_m *MockSecureAccessCloudClient) FindApplicationByName(name string) (*dto.ApplicationDTO, error) {
	ret := _m.Called(name)

	var r0 *dto.ApplicationDTO
	if rf, ok := ret.Get(0).(func(string) *dto.ApplicationDTO); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.ApplicationDTO)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindPoliciesByNames provides a mock function with given fields: name
func (_m *MockSecureAccessCloudClient) FindPoliciesByNames(name []string) ([]dto.PolicyDTO, error) {
	ret := _m.Called(name)

	var r0 []dto.PolicyDTO
	if rf, ok := ret.Get(0).(func([]string) []dto.PolicyDTO); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dto.PolicyDTO)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindPolicyByName provides a mock function with given fields: name
func (_m *MockSecureAccessCloudClient) FindPolicyByName(name string) (dto.PolicyDTO, error) {
	ret := _m.Called(name)

	var r0 dto.PolicyDTO
	if rf, ok := ret.Get(0).(func(string) dto.PolicyDTO); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(dto.PolicyDTO)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindSiteByName provides a mock function with given fields: name
func (_m *MockSecureAccessCloudClient) FindSiteByName(name string) (*dto.SiteDTO, error) {
	ret := _m.Called(name)

	var r0 *dto.SiteDTO
	if rf, ok := ret.Get(0).(func(string) *dto.SiteDTO); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.SiteDTO)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListConnectorsBySite provides a mock function with given fields: siteName
func (_m *MockSecureAccessCloudClient) ListConnectorsBySite(siteName string) ([]string, error) {
	ret := _m.Called(siteName)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(siteName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(siteName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateApplication provides a mock function with given fields: applicationDTO
func (_m *MockSecureAccessCloudClient) UpdateApplication(applicationDTO *dto.ApplicationDTO) (*dto.ApplicationDTO, error) {
	ret := _m.Called(applicationDTO)

	var r0 *dto.ApplicationDTO
	if rf, ok := ret.Get(0).(func(*dto.ApplicationDTO) *dto.ApplicationDTO); ok {
		r0 = rf(applicationDTO)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.ApplicationDTO)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*dto.ApplicationDTO) error); ok {
		r1 = rf(applicationDTO)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePolicies provides a mock function with given fields: applicationId, applicationType, policies
func (_m *MockSecureAccessCloudClient) UpdatePolicies(applicationId uuid.UUID, applicationType model.ApplicationType, policies []uuid.UUID) error {
	ret := _m.Called(applicationId, applicationType, policies)

	var r0 error
	if rf, ok := ret.Get(0).(func(uuid.UUID, model.ApplicationType, []uuid.UUID) error); ok {
		r0 = rf(applicationId, applicationType, policies)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
