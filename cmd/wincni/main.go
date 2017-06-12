package main

import (
	"net"
	"os"

	"github.com/containernetworking/cni/pkg/types/current"
)

func main() {
	_, ipnet, _ := net.ParseCIDR("1.2.3.4/24")

	ipconfig := &current.IPConfig{
		Version: "4",
		Address: *ipnet,
	}
	result := current.Result{
		CNIVersion: "0.3.1",
		IPs:        []*current.IPConfig{ipconfig},
	}

	result.Print()
	os.Exit(0)
}
