package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Piszmog/cfservices/v2"
)

func main() {
	// Manually setting env variable to illustrate usage
	err := os.Setenv("VCAP_SERVICES", `{
      "serviceA": [
        {
          "name":"service_a",
          "credentials": {
            "uri": "example_uri"
          }
        }
      ]
    }`)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = os.Unsetenv("VCAP_SERVICES"); err != nil {
			log.Fatal(err)
		}
	}()

	cred, err := cfservices.GetServiceCredentialsFromEnvironment("serviceA")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", cred)
}
