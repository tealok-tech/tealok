package main

import (
	"fmt"
	"github.com/tealok-tech/tealok/log"
	"github.com/tealok-tech/tealok/networkd"
	"github.com/tealok-tech/tealok/podman"
	"os"
)

func main() {
	// Set up the structured log
	log.Setup()
	// Get the correct subnet from the log
	_, subnet, err := networkd.Subnet()
	if err != nil {
		fmt.Println("Failed to get subnet", err)
		os.Exit(1)
	}
	fmt.Println("Subnet should be", subnet.IP, subnet.Mask.String())

	// Get or create a network with the correct subnet
	containerNetwork, err := podman.EnsureNetworkWithSubnet(subnet.IP)
	if err != nil {
		os.Exit(2)
	}
	fmt.Println("Tealok network name:", containerNetwork.Name)
}
