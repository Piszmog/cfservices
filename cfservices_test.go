package cfservices

import (
    "testing"
    "os"
)

func TestGetUriCredentials(t *testing.T) {
    const services = `{
      "serviceA": [
        {
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
    if credentials.Uri != "example_uri" {
        t.Errorf("retrived uri does not match %v", "example_uri")
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
        t.Errorf("retrived creditenals from non-existant service %v", err)
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
        t.Errorf("retrived creditenals from non-existant service %v", err)
    }
}

func TestGetCredentialsFromEmptyService(t *testing.T) {
    const services = `{
      "serviceA": [
      ]
    }`
    _, err := GetServiceCredentials("serviceA", services)
    if err == nil {
        t.Errorf("retrived creditenals from non-existant service %v", err)
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
            "port": "example_port"
          }
        }
      ]
    }`
    credentials, err := GetServiceCredentials("serviceA", services)
    if err != nil || credentials == nil {
        t.Errorf("failed to get credentials. %v", err)
    }
    if credentials.Uri != "example_uri" {
        t.Errorf("retrived uri does not match %v", "example_uri")
    }
    if credentials.APIUri != "example_httpAPIUri" {
        t.Errorf("retrived http api uri does not match %v", "example_httpAPIUri")
    }
    if credentials.LicenceKey != "example_licenseKey" {
        t.Errorf("retrived license ket does not match %v", "example_licenseKey")
    }
    if credentials.ClientSecret != "example_clientSecret" {
        t.Errorf("retrived client secret does not match %v", "example_clientSecret")
    }
    if credentials.ClientId != "example_clientId" {
        t.Errorf("retrived client id does not match %v", "example_clientId")
    }
    if credentials.AccessTokenUri != "example_accessTokenUri" {
        t.Errorf("retrived access token uri does not match %v", "example_accessTokenUri")
    }
    if credentials.Hostname != "example_hostname" {
        t.Errorf("retrived hostname does not match %v", "example_hostname")
    }
    if credentials.Username != "example_username" {
        t.Errorf("retrived username does not match %v", "example_username")
    }
    if credentials.Password != "example_password" {
        t.Errorf("retrived password does not match %v", "example_password")
    }
    if credentials.Port != "example_port" {
        t.Errorf("retrived port does not match %v", "example_port")
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
    vcapServices := LoadFromEnvironment()
    if len(vcapServices) == 0 {
        t.Errorf("failed to load services from environment")
    }
}

func TestLoadFromEnvironmentWhenNotSet(t *testing.T) {
    vcapServices := LoadFromEnvironment()
    if len(vcapServices) != 0 {
        t.Errorf("failed to load services from environment")
    }
}
