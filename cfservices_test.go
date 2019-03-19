package cfservices

import (
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
	credentials, err := GetServiceCredentials("serviceA", services)
	if err != nil || credentials == nil {
		t.Errorf("failed to get credentials. %v", err)
	}
	if credentials.Credentials[0].Uri != "example_uri" {
		t.Errorf("retrieved uri does not match %v", "example_uri")
	}
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
	serviceCred, err := GetServiceCredentials("serviceA", services)
	if err != nil || serviceCred == nil {
		t.Errorf("failed to get credentials. %v", err)
	}
	creds := serviceCred.Credentials
	if len(creds) != 2 {
		t.Errorf("service does not contain both credentials")
	}
	for _, cred := range creds {
		if cred.Uri != "example_uri" {
			t.Errorf("retrieved uri does not match %v", "example_uri")
		}
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
	_, err := GetServiceCredentials("serviceB", services)
	if err == nil {
		t.Errorf("retrieved creditenals from non-existent service %v", err)
	}
}

func TestInvalidJson(t *testing.T) {
	const services = `{
      "serviceA": [
        {
          "credentials": {
            "uri": "example_uri"
          },
        }
      ]
    }`
	_, err := GetServiceCredentials("serviceB", services)
	if err == nil {
		t.Errorf("retrieved creditenals from non-existent service %v", err)
	}
}

func TestGetCredentialsFromEmptyService(t *testing.T) {
	const services = `{
      "serviceA": [
      ]
    }`
	_, err := GetServiceCredentials("serviceA", services)
	if err == nil {
		t.Errorf("retrieved creditenals from non-existent service %v", err)
	}
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
            "jdbcUrl": "jdbc:mysql:/url"
          }
        }
      ]
    }`
	serviceCreds, err := GetServiceCredentials("serviceA", services)
	if err != nil || serviceCreds == nil {
		t.Errorf("failed to get credentials. %v", err)
	}
	for _, credentials := range serviceCreds.Credentials {
		if credentials.Uri != "example_uri" {
			t.Errorf("retrieved uri does not match %v", "example_uri")
		}
		if credentials.APIUri != "example_httpAPIUri" {
			t.Errorf("retrieved http api uri does not match %v", "example_httpAPIUri")
		}
		if credentials.LicenceKey != "example_licenseKey" {
			t.Errorf("retrieved license ket does not match %v", "example_licenseKey")
		}
		if credentials.ClientSecret != "example_clientSecret" {
			t.Errorf("retrieved client secret does not match %v", "example_clientSecret")
		}
		if credentials.ClientId != "example_clientId" {
			t.Errorf("retrieved client id does not match %v", "example_clientId")
		}
		if credentials.AccessTokenUri != "example_accessTokenUri" {
			t.Errorf("retrieved access token uri does not match %v", "example_accessTokenUri")
		}
		if credentials.Hostname != "example_hostname" {
			t.Errorf("retrieved hostname does not match %v", "example_hostname")
		}
		if credentials.Username != "example_username" {
			t.Errorf("retrieved username does not match %v", "example_username")
		}
		if credentials.Password != "example_password" {
			t.Errorf("retrieved password does not match %v", "example_password")
		}
		if credentials.Port.String() != "1234" {
			t.Errorf("retrieved port does not match %v", "example_port")
		}
		if credentials.JDBCUrl != "jdbc:mysql:/url" {
			t.Errorf("retrieved port does not match %v", "example_port")
		}
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
	serviceCreds, err := GetServiceCredentials("serviceA", services)
	if err != nil || serviceCreds == nil {
		t.Errorf("failed to get credentials. %v", err)
	}
	for _, credentials := range serviceCreds.Credentials {
		port, _ := credentials.Port.Int64()
		if port != 1234 {
			t.Errorf("retrieved port does not match %v", "example_port")
		}
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
	os.Setenv(VCAPServices, services)
	defer os.Unsetenv(VCAPServices)
	credentials, err := GetServiceCredentialsFromEnvironment("serviceA")
	if err != nil || credentials == nil {
		t.Errorf("failed to get credentials. %v", err)
	}
	if credentials.Credentials[0].Uri != "example_uri" {
		t.Errorf("retrieved uri does not match %v", "example_uri")
	}
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
	os.Setenv(VCAPServices, services)
	defer os.Unsetenv(VCAPServices)
	serviceCred, err := GetServiceCredentialsFromEnvironment("serviceA")
	if err != nil || serviceCred == nil {
		t.Errorf("failed to get credentials. %v", err)
	}
	creds := serviceCred.Credentials
	if len(creds) != 2 {
		t.Errorf("service does not contain both credentials")
	}
	for _, cred := range creds {
		if cred.Uri != "example_uri" {
			t.Errorf("retrieved uri does not match %v", "example_uri")
		}
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
	os.Setenv(VCAPServices, services)
	defer os.Unsetenv(VCAPServices)
	_, err := GetServiceCredentialsFromEnvironment("serviceB")
	if err == nil {
		t.Errorf("retrieved creditenals from non-existent service %v", err)
	}
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
	os.Setenv(VCAPServices, services)
	defer os.Unsetenv(VCAPServices)
	_, err := GetServiceCredentialsFromEnvironment("serviceB")
	if err == nil {
		t.Errorf("retrieved creditenals from non-existent service %v", err)
	}
}

func TestGetCredentialsFromEmptyServiceInEnv(t *testing.T) {
	const services = `{
      "serviceA": [
      ]
    }`
	os.Setenv(VCAPServices, services)
	defer os.Unsetenv(VCAPServices)
	_, err := GetServiceCredentialsFromEnvironment("serviceA")
	if err == nil {
		t.Errorf("retrieved creditenals from non-existent service %v", err)
	}
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
            "port": "example_port"
          }
        }
      ]
    }`
	os.Setenv(VCAPServices, services)
	defer os.Unsetenv(VCAPServices)
	serviceCreds, err := GetServiceCredentialsFromEnvironment("serviceA")
	if err != nil || serviceCreds == nil {
		t.Errorf("failed to get credentials. %v", err)
	}
	for _, credentials := range serviceCreds.Credentials {
		if credentials.Uri != "example_uri" {
			t.Errorf("retrieved uri does not match %v", "example_uri")
		}
		if credentials.APIUri != "example_httpAPIUri" {
			t.Errorf("retrieved http api uri does not match %v", "example_httpAPIUri")
		}
		if credentials.LicenceKey != "example_licenseKey" {
			t.Errorf("retrieved license ket does not match %v", "example_licenseKey")
		}
		if credentials.ClientSecret != "example_clientSecret" {
			t.Errorf("retrieved client secret does not match %v", "example_clientSecret")
		}
		if credentials.ClientId != "example_clientId" {
			t.Errorf("retrieved client id does not match %v", "example_clientId")
		}
		if credentials.AccessTokenUri != "example_accessTokenUri" {
			t.Errorf("retrieved access token uri does not match %v", "example_accessTokenUri")
		}
		if credentials.Hostname != "example_hostname" {
			t.Errorf("retrieved hostname does not match %v", "example_hostname")
		}
		if credentials.Username != "example_username" {
			t.Errorf("retrieved username does not match %v", "example_username")
		}
		if credentials.Password != "example_password" {
			t.Errorf("retrieved password does not match %v", "example_password")
		}
		if credentials.Port != "example_port" {
			t.Errorf("retrieved port does not match %v", "example_port")
		}
	}
}

func TestLoadFromEnvironment(t *testing.T) {
	os.Setenv(VCAPServices, `{
      "serviceA": [
        {
          "credentials": {
            "uri": "example_uri"
          }
        }
      ]
    }`)
	defer os.Unsetenv(VCAPServices)
	vcapServices := GetServices()
	if len(vcapServices) == 0 {
		t.Errorf("failed to load services from environment")
	}
}

func TestLoadFromEnvironmentWhenNotSet(t *testing.T) {
	vcapServices := GetServices()
	if len(vcapServices) != 0 {
		t.Errorf("failed to load services from environment")
	}
}

func TestLoadFromEnvironmentAsMap(t *testing.T) {
	os.Setenv(VCAPServices, `{
      "serviceA": [
        {
          "credentials": {
            "uri": "example_uri"
          }
        }
      ]
    }`)
	defer os.Unsetenv(VCAPServices)
	vcapServices, err := GetServicesAsMap()
	if len(vcapServices) == 0 {
		t.Errorf("failed to load services from environment")
	}
	if err != nil {
		t.Errorf("failed to load services from environment with error %v", err)
	}
}

func TestLoadFromEnvironmentAsMapWhenNotSet(t *testing.T) {
	vcapServices, err := GetServicesAsMap()
	if len(vcapServices) != 0 {
		t.Errorf("failed to load services from environment")
	}
	if err == nil {
		t.Errorf("failed to load services from environment with error %v", err)
	}
}
