package podman

import (
	"context"
	"fmt"
	"os"

	"github.com/containers/podman/v5/pkg/bindings"
	"github.com/containers/podman/v5/pkg/bindings/images"
)

func Networks() string {
	fmt.Println("Welcome to the Podman Go bindings tutorial")

	// Connect to Podman socket
	conn, err := bindings.NewConnection(context.Background(), "unix:///run/podman/podman.sock")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// List images
	imageSummary, err := images.List(conn, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var names []string
	for _, i := range imageSummary {
		names = append(names, i.RepoTags...)
	}
	fmt.Println("Listing images...")
	fmt.Println(names)
	return "yeah"
}
