package main

import (
	"fmt"
	"github.com/tealok-tech/tealok/networkd"
	"github.com/tealok-tech/tealok/podman"
)

func main() {
	fmt.Println("Subnet", networkd.Subnet())
	fmt.Println("Network", podman.Networks())

}
