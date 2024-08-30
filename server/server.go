package server

import (
	"errors"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type AddArgs struct {
	Name string
}

type Container int

func (t *Container) Add(args *AddArgs, reply *int) error {
	if args.Name == "" {
		return errors.New("Empty name")
	}
	log.Println("Adding", args.Name)
	return nil
}

func Run() {
	arith := new(Container)
	rpc.Register(arith)
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", ":1050")
	if err != nil {
		log.Fatal("listen error:", err)
	}
	http.Serve(l, nil)
}
