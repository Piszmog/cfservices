package cfservices

import (
    "github.com/Piszmog/cf-services/credentials"
    "encoding/json"
    "github.com/pkg/errors"
    "os"
)

const VCAPServices = "VCAP_SERVICES"

func LoadFromEnvironment() string {
    return os.Getenv(VCAPServices)
}

func GetServiceCredentials(serviceName string, services string) (*credentials.Credentials, error) {
    servicesJson := make(map[string][]map[string]credentials.Credentials)
    err := json.Unmarshal([]byte(services), &servicesJson)
    if err != nil {
        return nil, errors.Wrap(err, "failed to unmarshal JSON")
    }
    service := servicesJson[serviceName]
    if service == nil {
        return nil, errors.New("VCAP Service JSON does not contain " + serviceName)
    }
    var serviceCred credentials.Credentials
    if len(service) == 0 {
        return nil, errors.Errorf("%v has no data in JSON", serviceName)
    }
    serviceCred = service[0]["credentials"]
    return &serviceCred, nil
}
