package cfservices

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"reflect"
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
	URI            string                 `json:"uri"`
	JDBCUrl        string                 `json:"jdbcUrl"`
	APIUri         string                 `json:"http_api_uri"`
	LicenceKey     string                 `json:"licenseKey"`
	ClientSecret   string                 `json:"client_secret"`
	ClientID       string                 `json:"client_id"`
	AccessTokenURI string                 `json:"access_token_uri"`
	Hostname       string                 `json:"hostname"`
	Username       string                 `json:"username"`
	Password       string                 `json:"password"`
	Name           string                 `json:"name"`
	Additional     map[string]interface{} `json:"-"`
	Port           json.Number            `json:"port"`
}

type _cred Credentials

func (c *Credentials) UnmarshalJSON(b []byte) error {
	var cred _cred
	if err := json.Unmarshal(b, &cred); err != nil {
		return err
	}
	*c = Credentials(cred)
	var additional map[string]interface{}
	if err := json.Unmarshal(b, &additional); err != nil {
		return err
	}
	credStruct := reflect.TypeOf(Credentials{})
	for i := 0; i < credStruct.NumField(); i++ {
		field := credStruct.Field(i)
		if field.Name != "Additional" {
			delete(additional, field.Tag.Get("json"))
		}
	}
	if len(additional) > 0 {
		c.Additional = additional
	}
	return nil
}

// GetServiceCredentialsFromEnvironment retrieves from credentials for the environment variable 'VCAP_SERVICES'.
func GetServiceCredentialsFromEnvironment(serviceName string) (*ServiceCredentials, error) {
	services, err := GetServices()
	if err != nil {
		return nil, err
	}
	return GetServiceCredentials(services, serviceName)
}

// GetServices retrieves the JSON from the environment variable 'VCAP_SERVICES' and converts to to a map.
func GetServices() (map[string][]Service, error) {
	services := getServicesFromEnvironment()
	servicesJSON := make(map[string][]Service)
	err := json.Unmarshal([]byte(services), &servicesJSON)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}
	return servicesJSON, nil
}

func getServicesFromEnvironment() string {
	return os.Getenv(VCAPServices)
}

// ErrMissingService is the error when the service does not exist in provided slice of services.
var ErrMissingService = errors.New("service does not exist")

// GetServiceCredentials retrieves the credentials for the specified service from the provided services.
func GetServiceCredentials(services map[string][]Service, serviceName string) (*ServiceCredentials, error) {
	service := services[serviceName]
	if service == nil {
		return nil, ErrMissingService
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
