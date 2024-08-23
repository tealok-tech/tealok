package main

import (
	"fmt"
	"github.com/tealok-tech/tealok/networkd"
)

func main() {
	fmt.Println("Hey")
	fmt.Println("Subnet", networkd.Subnet())
}
