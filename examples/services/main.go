package main

import (
	"fmt"
	"github.com/Piszmog/cfservices"
	"os"
)

func main() {
	// Manually setting env variable to illustrate usage
	os.Setenv("VCAP_SERVICES", `{
      "serviceA": [
        {
          "name":"service_a",
          "credentials": {
            "uri": "example_uri"
          }
        }
      ]
    }`)
	defer os.Unsetenv("VCAP_SERVICES")

	services, err := cfservices.GetServices()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", services)
}
