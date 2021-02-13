package main

import (
	"fmt"
	"github.com/Piszmog/cfservices"
	"log"
)

func main() {
	services := map[string][]cfservices.Service{
		"serviceA": {
			{
				Name: "Service A",
				Credentials: cfservices.Credentials{
					Uri: "example_uri",
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
