package cfservices

import (
	"encoding/json"
	"github.com/pkg/errors"
	"os"
)

// VCAPServices is the environment variable that Cloud Foundry places all configurations for bounded services.
const VCAPServices = "VCAP_SERVICES"

// ServiceCredentials is the container for all credentials for a service type.
type ServiceCredentials struct {
	Credentials []Credentials
}

type service struct {
	Name         string      `json:"name"`
	InstanceName string      `json:"instance_name"`
	BindingName  string      `json:"binding_name"`
	Credentials  Credentials `json:"credentials"`
	Label        string      `json:"label"`
}

// Credentials is the credentials of a single bounded services.
type Credentials struct {
	Uri            string `json:"uri"`
	APIUri         string `json:"http_api_uri"`
	LicenceKey     string `json:"licenseKey"`
	ClientSecret   string `json:"client_secret"`
	ClientId       string `json:"client_id"`
	AccessTokenUri string `json:"access_token_uri"`
	Hostname       string `json:"hostname"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	Port           string `json:"port"`
}

// LoadFromEnvironment retrieves the JSON from the environment variables 'VCAP_SERVICES'.
func LoadFromEnvironment() string {
	return os.Getenv(VCAPServices)
}

// GetServiceCredentials Retrieves from credentials for the provided service from the 'VCAP_SERVICES' JSON.
func GetServiceCredentials(serviceName string, services string) (*ServiceCredentials, error) {
	servicesJSON := make(map[string][]service)
	err := json.Unmarshal([]byte(services), &servicesJSON)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal JSON")
	}
	service := servicesJSON[serviceName]
	if service == nil {
		return nil, errors.New("VCAP Service JSON does not contain " + serviceName)
	}
	if len(service) == 0 {
		return nil, errors.Errorf("%v has no data in JSON", serviceName)
	}
	var serviceCreds []Credentials
	for _, serviceObj := range service {
		serviceCreds = append(serviceCreds, serviceObj.Credentials)
	}
	return &ServiceCredentials{Credentials: serviceCreds}, nil
}
