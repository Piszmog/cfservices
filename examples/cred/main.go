package main

import (
	"fmt"
	"log"

	"github.com/Piszmog/cfservices/v2"
)

func main() {
	services := map[string][]cfservices.Service{
		"serviceA": {
			{
				Name: "Service A",
				Credentials: cfservices.Credentials{
					URI: "example_uri",
				},
			},
		},
	}
	cred, err := cfservices.GetServiceCredentials(services, "serviceA")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%+v\n", cred)
}
