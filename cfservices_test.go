package cfservices_test

import (
	"encoding/json"
	"github.com/Piszmog/cfservices"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetUriCredentials(t *testing.T) {
	const services = `{
      "serviceA": [
        {
          "name":"service_a",
          "credentials": {
            "uri": "example_uri"
          }
        }
      ]
    }`
	servicesJSON := make(map[string][]cfservices.Service)
	err := json.Unmarshal([]byte(services), &servicesJSON)
	assert.NoError(t, err, "failed to unmarshal JSON")
	credentials, err := cfservices.GetServiceCredentials(servicesJSON, "serviceA")
	assert.NoError(t, err, "failed to get credentials")
	assert.Equal(t, "example_uri", credentials.Credentials[0].Uri)
}

func TestGetUriCredentialsFromMultipleServices(t *testing.T) {
	const services = `{
      "serviceA": [
        {
          "credentials": {
            "uri": "example_uri"
          }
        },
        {
          "credentials": {
            "uri": "example_uri"
          }
        }
      ]
    }`
	servicesJSON := make(map[string][]cfservices.Service)
	err := json.Unmarshal([]byte(services), &servicesJSON)
	assert.NoError(t, err, "failed to unmarshal JSON")
	serviceCred, err := cfservices.GetServiceCredentials(servicesJSON, "serviceA")
	assert.NoError(t, err, "failed to get credentials")
	creds := serviceCred.Credentials
	assert.Equal(t, 2, len(creds))
	for _, cred := range creds {
		assert.Equal(t, "example_uri", cred.Uri)
	}
}

func TestGetCredentialsFromNonexistentService(t *testing.T) {
	const services = `{
      "serviceA": [
        {
          "credentials": {
            "uri": "example_uri"
          }
        }
      ]
    }`
	servicesJSON := make(map[string][]cfservices.Service)
	err := json.Unmarshal([]byte(services), &servicesJSON)
	assert.NoError(t, err, "failed to unmarshal JSON")
	_, err = cfservices.GetServiceCredentials(servicesJSON, "serviceB")
	assert.Error(t, err, "retrieved credentials from non-existent service")
}

func TestGetCredentialsFromEmptyService(t *testing.T) {
	const services = `{
      "serviceA": [
      ]
    }`
	servicesJSON := make(map[string][]cfservices.Service)
	err := json.Unmarshal([]byte(services), &servicesJSON)
	assert.NoError(t, err, "failed to unmarshal JSON")
	_, err = cfservices.GetServiceCredentials(servicesJSON, "serviceA")
	assert.Error(t, err, "retrieved credentials from non-existent service")
}

func TestGetFullCredentials(t *testing.T) {
	const services = `{
      "serviceA": [
        {
          "credentials": {
            "uri": "example_uri",
            "http_api_uri": "example_httpAPIUri",
            "licenseKey": "example_licenseKey",
            "client_secret": "example_clientSecret",
            "client_id": "example_clientId",
            "access_token_uri": "example_accessTokenUri",
            "hostname": "example_hostname",
            "username": "example_username",
            "password": "example_password",
            "port": "1234",
            "jdbcUrl": "jdbc:mysql:/url",
            "name": "someName"
          }
        }
      ]
    }`
	servicesJSON := make(map[string][]cfservices.Service)
	err := json.Unmarshal([]byte(services), &servicesJSON)
	assert.NoError(t, err, "failed to unmarshal JSON")
	serviceCreds, err := cfservices.GetServiceCredentials(servicesJSON, "serviceA")
	assert.NoError(t, err, "failed to get credentials")
	expectedCredentials := cfservices.Credentials{
		Uri:            "example_uri",
		JDBCUrl:        "jdbc:mysql:/url",
		APIUri:         "example_httpAPIUri",
		LicenceKey:     "example_licenseKey",
		ClientSecret:   "example_clientSecret",
		ClientId:       "example_clientId",
		AccessTokenUri: "example_accessTokenUri",
		Hostname:       "example_hostname",
		Username:       "example_username",
		Password:       "example_password",
		Port:           "1234",
		Name:           "someName",
	}
	for _, credentials := range serviceCreds.Credentials {
		assert.Equal(t, expectedCredentials, credentials)
	}
}

func TestGetPortAsNumber(t *testing.T) {
	const services = `{
      "serviceA": [
        {
          "credentials": {
            "port": 1234
          }
        }
      ]
    }`
	servicesJSON := make(map[string][]cfservices.Service)
	err := json.Unmarshal([]byte(services), &servicesJSON)
	assert.NoError(t, err, "failed to unmarshal JSON")
	serviceCreds, err := cfservices.GetServiceCredentials(servicesJSON, "serviceA")
	assert.NoError(t, err, "failed to get credentials")
	for _, credentials := range serviceCreds.Credentials {
		port, err := credentials.Port.Int64()
		assert.NoError(t, err, "failed to convert port integer")
		assert.Equal(t, int64(1234), port)
	}
}

func TestGetUriCredentialsFromEnv(t *testing.T) {
	const services = `{
      "serviceA": [
        {
          "name":"service_a",
          "credentials": {
            "uri": "example_uri"
          }
        }
      ]
    }`
	err := os.Setenv(cfservices.VCAPServices, services)
	assert.NoError(t, err, "failed to set environment variable")
	defer os.Unsetenv(cfservices.VCAPServices)
	credentials, err := cfservices.GetServiceCredentialsFromEnvironment("serviceA")
	assert.NoError(t, err, "failed to get credentials")
	assert.Equal(t, "example_uri", credentials.Credentials[0].Uri)
}

func TestGetUriCredentialsFromMultipleServicesInEnv(t *testing.T) {
	const services = `{
      "serviceA": [
        {
          "credentials": {
            "uri": "example_uri"
          }
        },
        {
          "credentials": {
            "uri": "example_uri"
          }
        }
      ]
    }`
	err := os.Setenv(cfservices.VCAPServices, services)
	assert.NoError(t, err, "failed to set environment variable")
	defer os.Unsetenv(cfservices.VCAPServices)
	serviceCred, err := cfservices.GetServiceCredentialsFromEnvironment("serviceA")
	assert.NoError(t, err, "failed to get credentials")
	creds := serviceCred.Credentials
	assert.Equal(t, 2, len(creds))
	for _, cred := range creds {
		assert.Equal(t, "example_uri", cred.Uri)
	}
}

func TestGetCredentialsFromNonexistentServiceInEnv(t *testing.T) {
	const services = `{
      "serviceA": [
        {
          "credentials": {
            "uri": "example_uri"
          }
        }
      ]
    }`
	err := os.Setenv(cfservices.VCAPServices, services)
	assert.NoError(t, err, "failed to set environment variable")
	defer os.Unsetenv(cfservices.VCAPServices)
	_, err = cfservices.GetServiceCredentialsFromEnvironment("serviceB")
	assert.Error(t, err, "retrieved credentials from non-existent service")
}

func TestInvalidJsonInEnv(t *testing.T) {
	const services = `{
      "serviceA": [
        {
          "credentials": {
            "uri": "example_uri"
          },
        }
      ]
    }`
	err := os.Setenv(cfservices.VCAPServices, services)
	assert.NoError(t, err, "failed to set environment variable")
	defer os.Unsetenv(cfservices.VCAPServices)
	_, err = cfservices.GetServiceCredentialsFromEnvironment("serviceB")
	assert.Error(t, err, "marshalled invalid JSON")
}

func TestGetCredentialsFromEmptyServiceInEnv(t *testing.T) {
	const services = `{
      "serviceA": [
      ]
    }`
	err := os.Setenv(cfservices.VCAPServices, services)
	assert.NoError(t, err, "failed to set environment variable")
	defer os.Unsetenv(cfservices.VCAPServices)
	_, err = cfservices.GetServiceCredentialsFromEnvironment("serviceA")
	assert.Error(t, err, "retrieved credentials from non-existent service")
}

func TestGetFullCredentialsInEnv(t *testing.T) {
	const services = `{
      "serviceA": [
        {
          "credentials": {
            "uri": "example_uri",
            "http_api_uri": "example_httpAPIUri",
            "licenseKey": "example_licenseKey",
            "client_secret": "example_clientSecret",
            "client_id": "example_clientId",
            "access_token_uri": "example_accessTokenUri",
            "hostname": "example_hostname",
            "username": "example_username",
            "password": "example_password",
            "port": "8080"
          }
        }
      ]
    }`
	err := os.Setenv(cfservices.VCAPServices, services)
	assert.NoError(t, err, "failed to set environment variable")
	defer os.Unsetenv(cfservices.VCAPServices)
	serviceCreds, err := cfservices.GetServiceCredentialsFromEnvironment("serviceA")
	assert.NoError(t, err, "failed to get credentials")
	expectedCredentials := cfservices.Credentials{
		Uri:            "example_uri",
		APIUri:         "example_httpAPIUri",
		LicenceKey:     "example_licenseKey",
		ClientSecret:   "example_clientSecret",
		ClientId:       "example_clientId",
		AccessTokenUri: "example_accessTokenUri",
		Hostname:       "example_hostname",
		Username:       "example_username",
		Password:       "example_password",
		Port:           "8080",
	}
	for _, credentials := range serviceCreds.Credentials {
		assert.Equal(t, expectedCredentials, credentials)
	}
}

func TestGetServices(t *testing.T) {
	err := os.Setenv(cfservices.VCAPServices, `{
      "serviceA": [
        {
          "credentials": {
            "uri": "example_uri"
          }
        }
      ]
    }`)
	assert.NoError(t, err, "failed to set environment variable")
	defer os.Unsetenv(cfservices.VCAPServices)
	vcapServices, err := cfservices.GetServices()
	assert.NotEqual(t, 0, len(vcapServices), "failed to load services from environment")
	assert.NoError(t, err, "failed to load services from environment with error")
}

func TestGetServicesWhenNotSet(t *testing.T) {
	vcapServices, err := cfservices.GetServices()
	assert.Equal(t, 0, len(vcapServices), "loaded services from environment")
	assert.Error(t, err, "loaded services from environment with error")
}
