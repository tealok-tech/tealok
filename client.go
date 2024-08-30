package main

import (
	"fmt"
	"github.com/tealok-tech/tealok/server"
	"log"
	"net/rpc"
	"os"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "[::1]:1050")
	if err != nil {
		log.Fatal("dialing:", err)
		os.Exit(1)
	}

	args := &server.Args{7, 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
		os.Exit(2)
	}
	fmt.Printf("Arith: %d*%d=%d", args.A, args.B, reply)
}
