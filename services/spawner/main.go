package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	hotel_pb "minornine.com/hotel/src/proto"

	"google.golang.org/grpc"
	"minornine.com/hotel/src/shared"
	"minornine.com/hotel/src/spawner"
)

const (
	DEFAULT_PORT = 3002
)

var masterServerAddress = shared.GetEnv("HOTEL_MASTER_ADDRESS", "")
var maxServersVar = shared.GetEnv("HOTEL_SPAWNER_MAX_SERVERS", "5")

func main() {
	shared.InitLogging()

	// Parse environment variables
	maxServers, err := strconv.ParseUint(maxServersVar, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("Unable to parse HOTEL_SPAWNER_MAX_SERVERS: %v", maxServersVar))
	}
	if masterServerAddress == "" {
		panic("Must provide HOTEL_MASTER_ADDRESS environment variable.")
	}
	log.Printf("Spawner configured to handle %v max servers.", maxServers)

	// Start the RPC server in a goroutine.
	addr := fmt.Sprintf(":%v", DEFAULT_PORT)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(fmt.Sprintf("Error binding TCP socket to %v", addr))
	}
	grpcServer := grpc.NewServer()
	hotel_pb.RegisterSpawnerServiceServer(grpcServer, &spawner.SpawnerService{})
	log.Println("Running gRPC server on", addr)
	go func() {
		log.Fatal(grpcServer.Serve(listener))
	}()

	// Test request loop.
	go func() {
		for {
			cl := spawner.NewMasterClient(masterServerAddress)
			cl.Test()
			time.Sleep(5 * time.Second)
		}
	}()

	// Setup a SIGINT (CTRL+C) shutdown signal and block on it.
	c := shared.CreateSigintChannel()
	<-c
	log.Println("Shutting down.")
	os.Exit(0)
}
