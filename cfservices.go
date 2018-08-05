package cfservices

import (
	"encoding/json"
	"github.com/Piszmog/cfservices/credentials"
	"github.com/pkg/errors"
	"os"
)

const VCAPServices = "VCAP_SERVICES"

// Retrieves the JSON from the environment variables 'VCAP_SERVICES'
func LoadFromEnvironment() string {
	return os.Getenv(VCAPServices)
}

// Retrieves from credentials for the provided service from the 'VCAP_SERVICES' JSON
func GetServiceCredentials(serviceName string, services string) (*credentials.ServiceCredentials, error) {
	servicesJson := make(map[string][]map[string]credentials.Credentials)
	err := json.Unmarshal([]byte(services), &servicesJson)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal JSON")
	}
	service := servicesJson[serviceName]
	if service == nil {
		return nil, errors.New("VCAP Service JSON does not contain " + serviceName)
	}
	if len(service) == 0 {
		return nil, errors.Errorf("%v has no data in JSON", serviceName)
	}
	var serviceCreds []credentials.Credentials
	for _, serviceObj := range service {
		serviceCreds = append(serviceCreds, serviceObj["credentials"])
	}
	return &credentials.ServiceCredentials{Credentials: serviceCreds}, nil
}
