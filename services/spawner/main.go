package main

import (
	"fmt"
	"log"
	"net"

	"minornine.com/hotel/src/proto"

	"google.golang.org/grpc"
	"minornine.com/hotel/src/shared"
	"minornine.com/hotel/src/spawner"
)

const (
	DEFAULT_PORT = 3001
)

func main() {
	shared.InitLogging()

	// Start a TCP server and connect gRPC to it.
	addr := fmt.Sprintf(":%v", DEFAULT_PORT)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(fmt.Sprintf("Error listening to %v", addr))
	}

	// Install all RPC handlers.
	grpcServer := grpc.NewServer()
	proto.RegisterSpawnerServiceServer(grpcServer, &spawner.TestService{})

	log.Println("Running gRPC server on", addr)
	grpcServer.Serve(listener)
}
