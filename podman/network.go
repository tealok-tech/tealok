package podman

import (
	"context"
	"fmt"
	"os"

	"github.com/containers/podman/v5/pkg/bindings"
	"github.com/containers/podman/v5/pkg/bindings/network"
)

func Networks() string {
	fmt.Println("Welcome to the Podman Go bindings tutorial")

	// Connect to Podman socket
	conn, err := bindings.NewConnection(context.Background(), "unix:///run/podman/podman.sock")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// List networks
	networks, err := network.List(conn, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, network := range networks {
		fmt.Println(network.Name)
	}
	return "yeah"
}
