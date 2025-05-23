# CF Services
[![Go Reference](https://pkg.go.dev/badge/github.com/Piszmog/cfservices.svg)](https://pkg.go.dev/github.com/Piszmog/cfservices/v2)
[![Build Status](https://github.com/Piszmog/cfservices/workflows/Go/badge.svg)](https://github.com/Piszmog/cfservices/workflows/Go/badge.svg)
[![Coverage Status](https://coveralls.io/repos/github/Piszmog/cfservices/badge.svg?branch=master)](https://coveralls.io/github/Piszmog/cfservices?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/Piszmog/cfservices)](https://goreportcard.com/report/github.com/Piszmog/cfservices/v2)
[![GitHub release](https://img.shields.io/github/release/Piszmog/cfservices.svg)](https://github.com/Piszmog/cfservices/releases/latest)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

This library is aimed at removing the boilerplate code and let developers just worry about using actually connecting to 
services they have bounded to their app.

`go get github.com/Piszmog/cfservices/v2`

## Retrieving VCAP_SERVICES
Simply use `cfservices.GetServices()`.

```go
package main
import "github.com/Piszmog/cfservices/v2"

func main() {
	services, err := cfservices.GetServices()
	if err != nil {
		// handle error
	}
	service := services["serviceA"]
	// Use information about service A to perform actions (such as creating an OAuth2 Client)
}
```

## Retrieving Credentials of a Service
Call `cfservices.GetServiceCredentials(..)` by passing the `VCAP_SERVICES` marshalled JSON and the name of the service to retrieve the 
credentials for. If `VCAP_SERVICES` is guaranteed to be an environment variable use `cfservices.GetServiceCredentialsFromEnvironment(..)` 
instead.

```go
package main
import "github.com/Piszmog/cfservices/v2"

func main() {
	var services map[string][]cfservices.Service
	// Read the services into the struct
	creds, err := cfservices.GetServiceCredentials(services, "serviceB")
	if err != nil {
		// handle error
	}
	// Use credentials...
	
	// Retrieve the JSON from the environment
	creds, err = cfservices.GetServiceCredentialsFromEnvironment("serviceB")
	if err != nil {
		// handle error
	}
	// Use credentials...
}
```
