package podman

import (
	"context"
	"fmt"
	"net"

	"github.com/tealok-tech/tealok/log"

	"github.com/containers/common/libnetwork/types"
	"github.com/containers/podman/v5/pkg/bindings"
	"github.com/containers/podman/v5/pkg/bindings/network"
)

// findNetworkWithSubnet finds a network with the provided subnet, or nil if it can't be found.
func findNetworkWithSubnet(networks []types.Network, subnet net.IPNet) *types.Network {
	for _, network := range networks {
		for _, sub := range network.Subnets {
			fmt.Println("Checking network", network.Name, "at subnet", sub.Subnet.IP, sub.Subnet.Mask)
			if sub.Subnet.IP.Equal(subnet.IP) {
				// At this point we could also check that the subnet mask is the same.
				// we don't actually want that since the container network will be a sub-subnet
				return &network
			}
		}
	}
	return nil
}

// CreateNetworkWithSubnet creates a network with the given subnet
func CreateNetworkWithSubnet(conn context.Context, subnet net.IPNet) (*types.Network, error) {
	podmanSubnet := types.Subnet{}
	podmanSubnet.Subnet = types.IPNet{}
	podmanSubnet.Subnet.IP = subnet.IP
	podmanSubnet.Subnet.Mask = subnet.Mask

	subnets := [1]types.Subnet{podmanSubnet}
	n := new(types.Network)
	n.Name = "tealok-net"
	n.Driver = "bridge"
	n.Subnets = subnets[:]
	n.IPv6Enabled = true
	n.Internal = false
	n.DNSEnabled = false
	newNetwork, err := network.Create(conn, n)
	if err != nil {
		return nil, err
	}
	log.NetworkCreated(newNetwork.Name)
	return &newNetwork, nil
}

// EnsureNetworkWithSubnet returns the name of a network with the provided subnet CIDR.
// If a network does not exist, one is created.
func EnsureNetworkWithSubnet(ip net.IP) (*types.Network, error) {
	subnet := net.IPNet{ip, net.CIDRMask(64, 128)}

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

	net := findNetworkWithSubnet(networks, subnet)
	if net == nil {
		net, err = CreateNetworkWithSubnet(conn, subnet)
		if err != nil {
			return nil, err
		}
	}
	return net, nil
}
