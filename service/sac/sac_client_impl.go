package sac

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"bitbucket.org/accezz-io/sac-operator/service/sac/dto"
	"github.com/google/uuid"
	"golang.org/x/oauth2/clientcredentials"
	"gopkg.in/resty.v1"
)

var ErrorPermissionDenied = fmt.Errorf("permission denied")
var ErrorNotFound = fmt.Errorf("not found")
var ErrConflict = fmt.Errorf("already exist")

func IsConflict(err error) bool {
	return err == ErrConflict
}

type SecureAccessCloudClientImpl struct {
	Setting *SecureAccessCloudSettings
	Client  *resty.Client
}

func NewSecureAccessCloudClientImpl(setting *SecureAccessCloudSettings) SecureAccessCloudClient {
	return &SecureAccessCloudClientImpl{Client: nil, Setting: setting}
}

// ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Application API
// ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (s *SecureAccessCloudClientImpl) CreateApplication(applicationDTO *dto.ApplicationDTO) (*dto.ApplicationDTO, error) {
	// TODO: implement
	return nil, nil
}

func (s *SecureAccessCloudClientImpl) UpdateApplication(applicationDTO *dto.ApplicationDTO) (*dto.ApplicationDTO, error) {
	// TODO: implement
	return nil, nil
}

func (s *SecureAccessCloudClientImpl) FindApplicationByName(name string) (*dto.ApplicationDTO, error) {
	endpoint := s.Setting.BuildAPIPrefixURL() + "/v2/applications" + "?filter=" + url.QueryEscape(name)

	var applications dto.ApplicationPageDTO
	err := s.performGetRequest(endpoint, &applications)

	if err != nil {
		return &dto.ApplicationDTO{}, err
	}

	if applications.NumberOfElements == 0 {
		return &dto.ApplicationDTO{}, ErrorNotFound
	}

	// Return the first policy if more than one found
	return &applications.Content[0], nil
}

func (s *SecureAccessCloudClientImpl) DeleteApplication(id uuid.UUID) error {
	// TODO: implement
	return nil
}

// ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Policy API
// ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (s *SecureAccessCloudClientImpl) FindPolicyByName(name string) (dto.PolicyDTO, error) {
	endpoint := s.Setting.BuildAPIPrefixURL() + "/v2/policies" + "?filter=" + url.QueryEscape(name)

	var policies dto.PoliciesPageDTO
	err := s.performGetRequest(endpoint, &policies)

	if err != nil {
		return dto.PolicyDTO{}, err
	}

	if policies.NumberOfElements == 0 {
		return dto.PolicyDTO{}, ErrorNotFound
	}

	// Return the first policy if more than one found
	return policies.Content[0], nil
}

func (s *SecureAccessCloudClientImpl) FindPoliciesByNames(name []string) ([]dto.PolicyDTO, error) {
	// TODO: implement
	return nil, nil
}

func (s *SecureAccessCloudClientImpl) UpdatePolicies(applicationId uuid.UUID, policies []uuid.UUID) error {
	// TODO: implement
	return nil
}

// ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Site API
// ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (s *SecureAccessCloudClientImpl) CreateSite(siteDTO *dto.SiteDTO) (*dto.SiteDTO, error) {
	endpoint := s.Setting.BuildAPIPrefixURL() + "/v2/sites"

	site := &dto.SiteDTO{}
	err := s.performPostRequest(endpoint, siteDTO, site)
	if err != nil {
		return &dto.SiteDTO{}, err
	}

	return site, nil
}
func (s *SecureAccessCloudClientImpl) DeleteSite(id string) error {
	endpoint := s.Setting.BuildAPIPrefixURL() + "/v2/sites/" + id

	return s.performDeleteRequest(endpoint)

}
func (s *SecureAccessCloudClientImpl) FindSiteByName(name string) (*dto.SiteDTO, error) {
	endpoint := s.Setting.BuildAPIPrefixURL() + "/v2/sites" + "?filter=" + url.QueryEscape(name)

	var pageDTO dto.SitePageDTO

	err := s.performGetRequest(endpoint, &pageDTO)

	if err != nil {
		return &dto.SiteDTO{}, err
	}

	if pageDTO.NumberOfElements == 0 {
		return &dto.SiteDTO{}, ErrorNotFound
	}

	// Return the first policy if more than one found
	return &pageDTO.Content[0], nil
}
func (s *SecureAccessCloudClientImpl) BindApplicationToSite(applicationId uuid.UUID, siteId string) error {
	// TODO: implement
	return nil
}

func (s *SecureAccessCloudClientImpl) CreateConnector(siteDTO *dto.SiteDTO, connectorName string) (*dto.ConnectorObjects, error) {
	endpoint := s.Setting.BuildAPIPrefixURL() + "/v2/connectors?bind_to_site_id=" + siteDTO.ID

	connector := &dto.ConnectorObjects{
		Name:           connectorName,
		DeploymentType: "linux",
	}

	err := s.performPostRequest(endpoint, connector, connector)
	if err != nil {
		return &dto.ConnectorObjects{}, err
	}

	return connector, nil
}

func (s *SecureAccessCloudClientImpl) ListConnectorsBySite(siteName string) ([]string, error) {
	site, err := s.FindSiteByName(siteName)
	if err != nil {
		return nil, err
	}
	return site.Connectors, nil
}

func (s *SecureAccessCloudClientImpl) DeleteConnector(id string) error {
	endpoint := s.Setting.BuildAPIPrefixURL() + "/v2/connectors/" + id

	return s.performDeleteRequest(endpoint)
}

// ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// ConnectorObjects API
// ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Private Functions
// ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (s *SecureAccessCloudClientImpl) getClient() *resty.Client {
	if s.Client != nil {
		return s.Client
	}

	cfg := clientcredentials.Config{
		ClientID:     s.Setting.ClientID,
		ClientSecret: s.Setting.ClientSecret,
		TokenURL:     s.Setting.BuildOAuthTokenURL(),
		Scopes:       []string{},
	}

	oauthClient := cfg.Client(context.Background())
	client := resty.New().SetRetryCount(0).SetTimeout(1 * time.Minute)
	// http://godoc.org/golang.org/x/oauth2 implements `httpRoundTripper` interface
	// Set the oauthClient transport
	client.SetTransport(oauthClient.Transport)

	s.Client = client
	return s.Client
}

func (s *SecureAccessCloudClientImpl) performGetRequest(endpoint string, obj interface{}) error {
	// 1. Get Authorization Token
	client := s.getClient()

	// 2. Perform the GET request
	response, err := client.NewRequest().Get(endpoint)
	if err != nil {
		return err
	}

	if response.StatusCode() == http.StatusNotFound {
		return ErrorNotFound
	}

	if response.StatusCode() != http.StatusOK {
		return fmt.Errorf("failed with status-code: %d and body: %s", response.StatusCode(), response.String())
	}

	// 3. Convert to Commit model
	err = json.Unmarshal(response.Body(), &obj)
	if err != nil {
		return err
	}

	return nil
}

func (s *SecureAccessCloudClientImpl) performPostRequest(endpoint string, body, obj interface{}) error {
	// 1. Get Authorization Token
	client := s.getClient()

	// 2. Perform the POST request
	response, err := client.NewRequest().SetBody(body).Post(endpoint)
	if err != nil {
		return err
	}

	if response.StatusCode() == http.StatusConflict {
		return ErrConflict
	}

	if response.StatusCode() != http.StatusCreated {
		return fmt.Errorf("failed with status-code: %d and body: %s", response.StatusCode(), response.String())
	}

	// 3. Convert to Commit model
	err = json.Unmarshal(response.Body(), obj)
	if err != nil {
		return err
	}

	return nil
}

func (s *SecureAccessCloudClientImpl) performDeleteRequest(endpoint string) error {
	// 1. Get Authorization Token
	client := s.getClient()

	// 2. Perform the DELETE request
	response, err := client.NewRequest().Delete(endpoint)
	if err != nil {
		return err
	}

	if response.StatusCode() != http.StatusNoContent {
		return fmt.Errorf("failed with status-code: %d and body: %s", response.StatusCode(), response.String())
	}

	return nil
}
