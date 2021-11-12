package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/inlets/cloud-provision/provision"
)

func main() {
	var (
		accessToken  string
		userDataFile string
		userdata     string
		hostname     string
		region       string
	)

	flag.StringVar(&accessToken, "access-token", "", "Access token for provisioning a host")
	flag.StringVar(&userDataFile, "userdata-file", "", "Apply user-data from a file to configure the host")
	flag.StringVar(&hostname, "hostname", "provision-example", "Name for the host")
	flag.StringVar(&region, "region", "lon1", "Region for the host")

	flag.Parse()

	if len(accessToken) == 0 {
		fmt.Fprintf(os.Stderr, "--access-token required")
		os.Exit(1)
	}

	provisioner, err := provision.NewDigitalOceanProvisioner(accessToken)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err.Error())
		os.Exit(1)
	}

	if len(userDataFile) > 0 {
		res, err := ioutil.ReadFile(userDataFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s", err.Error())
			os.Exit(1)
		}

		userdata = string(res)
	}

	// Find examples here for other clouds -> https://github.com/inlets/inletsctl/blob/356886f41e7c48a9644a24532027d1defa1d69e8/cmd/create.go
	res, err := provisioner.Provision(provision.BasicHost{
		Name:       hostname,
		OS:         "ubuntu-18-04-x64",
		Plan:       "s-1vcpu-1gb",
		Region:     region,
		UserData:   userdata,
		Additional: map[string]string{},
	})

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err.Error())
		os.Exit(1)
	}

	fmt.Printf("Host ID: %s\n", res.ID)

	pollStatusAttempts := 250
	waitInterval := time.Second * 2
	for i := 0; i <= pollStatusAttempts; i++ {
		fmt.Printf("Polling status: %d/%d\n", i+1, pollStatusAttempts)
		res, err := provisioner.Status(res.ID)

		if err != nil {
			fmt.Fprintf(os.Stderr, "%s", err.Error())
			os.Exit(1)
		}
		if res.Status == provision.ActiveStatus {
			fmt.Printf("Your IP address is: %s\n", res.IP)
			break
		}
		time.Sleep(waitInterval)
	}
}
