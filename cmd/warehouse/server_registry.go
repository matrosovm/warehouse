package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"

	"github.com/matrosovm/warehouse/internal/app/warehouse"
	"github.com/matrosovm/warehouse/internal/pkg/database"
)

func runRPC(store database.Store) {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer listener.Close()

	server := rpc.NewServer()
	if err = server.Register(warehouse.NewService(store)); err != nil {
		log.Fatalf("failed to register service: %v", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("failed to accept: %v", err)
		}

		go server.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
