package main

import (
	"flag"
	"fmt"
	"github.com/tealok-tech/tealok/server"
	"log"
	"net/rpc"
	"os"
)

func main() {
	// Parse the sub-command we want
	if len(os.Args) < 2 {
		fmt.Println("You must specify a command such as 'add'")
		os.Exit(1)
	}
	command := os.Args[1]
	switch command {
	case "add":
		add()
	case "help":
		fmt.Println("help.")
	default:
		fmt.Println("Unrecognized command", command)
		os.Exit(2)
	}
}

func add() {
	name := flag.String("name", "", "The name of the container to add")
	flag.CommandLine.Parse(os.Args[2:])
	if *name == "" {
		fmt.Println("You must specify a name such as '-name foo'")
		os.Exit(3)
	}
	fmt.Println("Adding ", *name)

	client, err := rpc.DialHTTP("tcp", "[::1]:1050")
	if err != nil {
		log.Fatal("dialing:", err)
		os.Exit(1)
	}

	args := &server.AddArgs{*name}
	var reply int
	err = client.Call("Server.Add", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
		os.Exit(2)
	}
	fmt.Printf("Container: %s=%d", args.Name, reply)
}
