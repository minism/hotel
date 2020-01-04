package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	hotel_pb "minornine.com/hotel/src/proto"
	"minornine.com/hotel/src/spawner"

	"google.golang.org/grpc"
	"minornine.com/hotel/src/shared"
)

const (
	DEFAULT_PORT           uint32 = 3002
	MASTER_CONTACT_RETRIES        = 5
)

var masterServerAddress = shared.GetEnv("HOTEL_MASTER_ADDRESS", "")
var maxServersVar = shared.GetEnv("HOTEL_SPAWNER_MAX_SERVERS", "5")
var configPath = shared.GetEnv("HOTEL_CONFIG_PATH", "./services/spawner/example.config.json")

func main() {
	shared.InitLogging()

	// Validate environment variables.
	maxServers, err := strconv.ParseUint(maxServersVar, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("Unable to parse HOTEL_SPAWNER_MAX_SERVERS: %v", maxServersVar))
	}
	if masterServerAddress == "" {
		panic("Must provide HOTEL_MASTER_ADDRESS environment variable.")
	}
	log.Printf("Spawner configured to handle %v max servers.", maxServers)

	// Initialize main components.
	config := spawner.LoadConfig(configPath)
	service := spawner.NewSpawnerService(&config)

	// Start the RPC server in a goroutine.
	port := DEFAULT_PORT
	addr := fmt.Sprintf(":%v", port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(fmt.Sprintf("Error binding TCP socket to %v", addr))
	}
	grpcServer := grpc.NewServer()
	hotel_pb.RegisterSpawnerServiceServer(grpcServer, &service)
	log.Println("Running gRPC server on", addr)
	go func() {
		log.Fatal(grpcServer.Serve(listener))
	}()

	// Register with the master server on startup.
	go func() {
		masterClient := spawner.NewMasterClient(masterServerAddress)
		for i := 0; i < MASTER_CONTACT_RETRIES; i++ {
			err := masterClient.Register(port, service.GetStatus())
			if err == nil {
				return
			}
			time.Sleep(5 * time.Second)
		}

		// Failed after allowed retries, abort.
		panic("Unable to contact master server after allowed retries, shutting down.")
	}()

	// Setup a SIGINT (CTRL+C) shutdown signal and block on it.
	c := shared.CreateSigintChannel()
	<-c
	log.Println("Shutting down.")
	os.Exit(0)
}
