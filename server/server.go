package server

import (
	"database/sql"
	"errors"
	"github.com/tealok-tech/tealok/database"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type AddArgs struct {
	Name string
}

type Server struct {
	DB *sql.DB
}

func (s *Server) Add(args *AddArgs, reply *int) error {
	if args.Name == "" {
		return errors.New("Empty name")
	}
	log.Println("Adding", args.Name)

	err := database.AddContainer(s.DB, args.Name)
	if err != nil {
		log.Fatal("Unable to add container: " + err.Error())
	}
	return nil
}

func Run(db *sql.DB) {
	server := new(Server)
	server.DB = db

	rpc.Register(server)
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", ":1050")
	if err != nil {
		log.Fatal("listen error:", err)
	}
	http.Serve(l, nil)
}
