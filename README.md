# CF Services
[![Build Status](https://travis-ci.org/Piszmog/cfservices.svg?branch=develop)](https://travis-ci.org/Piszmog/cfservices)
[![Coverage Status](https://coveralls.io/repos/github/Piszmog/cfservices/badge.svg?branch=develop)](https://coveralls.io/github/Piszmog/cfservices?branch=develop)
[![Go Report Card](https://goreportcard.com/badge/github.com/Piszmog/cfservices)](https://goreportcard.com/report/github.com/Piszmog/cfservices)
[![GitHub release](https://img.shields.io/github/release/Piszmog/cfservices.svg)](https://github.com/Piszmog/cfservices/releases/latest)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

When I started learning Go and deploying it to a cloud environment, a majority of the tutorials and examples I came across 
use Docker/containers. My professional background involves deploying applications to Cloud Foundry where containers are not 
a worry. As I started deploying Go apps to Cloud Foundry, I found myself rewriting the same code to pull down the `VCAP_SERVICES` 
environment variable and parsing out the credentials to connect to services.

This library is aimed at removing the boilerplate code and let developers just worry about using actually connecting to 
services they have bounded to their app.

## Loading from Environment
Simply use the returned string from `cfservices.LoadFromEnvironment()`.

## Retrieving Credentials of a Service
Call `cfservices.GetServiceCredentials(..)` by passing the `VCAP_SERVICES` JSON and the name of the service to retrieve the 
credentials for.

## Dependencies
* [errors](https://github.com/pkg/errors)