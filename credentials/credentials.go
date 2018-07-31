package credentials

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
