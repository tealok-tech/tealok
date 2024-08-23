package podman

import (
	"context"
	"errors"
	"fmt"
	"net"

	"github.com/containers/common/libnetwork/types"
	"github.com/containers/podman/v5/pkg/bindings"
	"github.com/containers/podman/v5/pkg/bindings/network"
)

// findNetworkWithSubnet finds a network with the provided subnet, or nil if it can't be found.
func findNetworkWithSubnet(networks []types.Network, subnet *net.IPNet) (*types.Network, error) {
	for _, network := range networks {
		for _, sub := range network.Subnets {
			fmt.Println("Network", network.Name, "at subnet", sub.Subnet.IP, sub.Subnet.Mask)
			if sub.Subnet.IP.Equal(subnet.IP) {
				// At this point we could also check that the subnet mask is the same.
				// we don't actually want that since the container network will be a sub-subnet
				return &network, nil
			}
		}
	}
	return nil, errors.New("No matching subnet")
}

func CreateNetworkWithSubnet(conn context.Context, subnet *net.IPNet) (*types.Network, error) {
	return nil, errors.New("Not implemented")
}

// EnsureNetworkWithSubnet returns the name of a network with the provided subnet CIDR.
// If a network does not exist, one is created.
func EnsureNetworkWithSubnet(subnet *net.IPNet) (*types.Network, error) {
	// Connect to Podman socket
	conn, err := bindings.NewConnection(context.Background(), "unix:///run/podman/podman.sock")
	if err != nil {
		return nil, err
	}

	// List networks
	networks, err := network.List(conn, nil)
	if err != nil {
		return nil, err
	}

	net, err := findNetworkWithSubnet(networks, subnet)
	if err != nil {
		return nil, err
	}
	if net == nil {
		net, err = CreateNetworkWithSubnet(conn, subnet)
		if err != nil {
			return nil, err
		}
	}
	return net, nil
}
