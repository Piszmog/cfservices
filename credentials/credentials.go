package credentials

// ServiceCredentials is the container for all credentials for a service type.
type ServiceCredentials struct {
	Credentials []Credentials
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
