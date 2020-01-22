package cfservices

import (
	"encoding/json"
	"fmt"
	"os"
)

// VCAPServices is the environment variable that Cloud Foundry places all configurations for bounded services.
const VCAPServices = "VCAP_SERVICES"

// ServiceCredentials is the container for all credentials for a service type.
type ServiceCredentials struct {
	Credentials []Credentials
}

// Service is contains all the information for a service bounded to an application.
type Service struct {
	Name         string      `json:"name"`
	InstanceName string      `json:"instance_name"`
	BindingName  string      `json:"binding_name"`
	Credentials  Credentials `json:"credentials"`
	Label        string      `json:"label"`
}

// Credentials is the credentials of a single bounded services.
type Credentials struct {
	Uri            string      `json:"uri"`
	JDBCUrl        string      `json:"jdbcUrl"`
	APIUri         string      `json:"http_api_uri"`
	LicenceKey     string      `json:"licenseKey"`
	ClientSecret   string      `json:"client_secret"`
	ClientId       string      `json:"client_id"`
	AccessTokenUri string      `json:"access_token_uri"`
	Hostname       string      `json:"hostname"`
	Username       string      `json:"username"`
	Password       string      `json:"password"`
	Port           json.Number `json:"port"`
	Name           string      `json:"name"`
}

// GetServices retrieves the JSON from the environment variable 'VCAP_SERVICES'.
func GetServices() string {
	return os.Getenv(VCAPServices)
}

// GetServicesAsMap retrieves the JSON from the environment variable 'VCAP_SERVICES' and converts to to a map.
func GetServicesAsMap() (map[string][]Service, error) {
	services := GetServices()
	servicesJSON := make(map[string][]Service)
	err := json.Unmarshal([]byte(services), &servicesJSON)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}
	return servicesJSON, nil
}

// GetServiceCredentials retrieves from credentials for the provided service from the 'VCAP_SERVICES' JSON.
func GetServiceCredentials(serviceName string, services string) (*ServiceCredentials, error) {
	servicesJSON := make(map[string][]Service)
	err := json.Unmarshal([]byte(services), &servicesJSON)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}
	service := servicesJSON[serviceName]
	if service == nil {
		return nil, fmt.Errorf("VCAP Service JSON does not contain %s", serviceName)
	}
	if len(service) == 0 {
		return nil, fmt.Errorf("%s has no data", serviceName)
	}
	serviceCreds := make([]Credentials, len(service))
	for index, serviceObj := range service {
		serviceCreds[index] = serviceObj.Credentials
	}
	return &ServiceCredentials{Credentials: serviceCreds}, nil
}

// GetServiceCredentialsFromEnvironment retrieves from credentials for the environment variable 'VCAP_SERVICES'.
func GetServiceCredentialsFromEnvironment(serviceName string) (*ServiceCredentials, error) {
	services, err := GetServicesAsMap()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve services from environment: %w", err)
	}
	service := services[serviceName]
	if service == nil {
		return nil, fmt.Errorf("VCAP Service JSON does not contain %s", serviceName)
	}
	if len(service) == 0 {
		return nil, fmt.Errorf("%s has no data in JSON", serviceName)
	}
	serviceCreds := make([]Credentials, len(service))
	for index, serviceObj := range service {
		serviceCreds[index] = serviceObj.Credentials
	}
	return &ServiceCredentials{Credentials: serviceCreds}, nil
}
