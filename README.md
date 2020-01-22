# CF Services
[![Build Status](https://travis-ci.org/Piszmog/cfservices.svg?branch=develop)](https://travis-ci.org/Piszmog/cfservices)
[![Coverage Status](https://coveralls.io/repos/github/Piszmog/cfservices/badge.svg?branch=develop)](https://coveralls.io/github/Piszmog/cfservices?branch=develop)
[![Go Report Card](https://goreportcard.com/badge/github.com/Piszmog/cfservices)](https://goreportcard.com/report/github.com/Piszmog/cfservices)
[![GitHub release](https://img.shields.io/github/release/Piszmog/cfservices.svg)](https://github.com/Piszmog/cfservices/releases/latest)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

When I started learning Go and deploying it to a cloud environment, a majority of the tutorials and examples I came across 
use Docker/containers. My professional background involves deploying applications to Cloud Foundry where containers are not 
a worry. As I started deploying Go apps to Cloud Foundry, I found myself rewriting the same code to pull down the `VCAP_SERVICES` 
environment variable and parsing out the credentials to connect to services.

This library is aimed at removing the boilerplate code and let developers just worry about using actually connecting to 
services they have bounded to their app.

## Retrieving Raw VCAP_SERVICES
Simply use the returned string from `cfservices.GetServices()`.

```go
package main
import "github.com/Piszmog/cfservices"

func main() {
	services := cfservices.GetServices()
	// Parse JSON
}
```

## Retrieving Parsed VCAP_SERVICES
To have the environment variable parsed, use `cfservices.GetServicesAsMap()`.

```go
package main
import "github.com/Piszmog/cfservices"

func main() {
	services, err := cfservices.GetServicesAsMap()
	if err != nil {
		// handle error
	}
	service := services["serviceA"]
	// Use information about service A to perform actions (such as creating an OAuth2 Client)
}
```

## Retrieving Credentials of a Service
Call `cfservices.GetServiceCredentials(..)` by passing the `VCAP_SERVICES` JSON and the name of the service to retrieve the 
credentials for. If `VCAP_SERVICES` is guaranteed to be an environment variable use `cfservices.GetServiceCredentialsFromEnvironment(..)` 
instead.

```go
package main
import "github.com/Piszmog/cfservices"

func main() {
	creds, err := cfservices.GetServiceCredentials("RAW_JSON", "serviceB")
	if err != nil {
		// handle error
	}
	// Use credentials
}
```

```go
package main
import "github.com/Piszmog/cfservices"

func main() {
	creds, err := cfservices.GetServiceCredentialsFromEnvironment("serviceB")
	if err != nil {
		// handle error
	}
	// Use credentials
}
```