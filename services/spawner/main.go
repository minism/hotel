package main

import (
	"fmt"
	"log"
	"net"
	"time"

	hotel_pb "github.com/minism/hotel/src/proto"
	"github.com/minism/hotel/src/spawner"

	"google.golang.org/grpc"
	"github.com/minism/hotel/src/shared"
)

const (
	MASTER_CONTACT_RETRIES        = 5
)

var masterServerAddress = shared.GetEnv("HOTEL_MASTER_ADDRESS", "")
var dataPath = shared.GetEnv("HOTEL_DATA_PATH", ".")
var configPath = shared.GetEnv("HOTEL_CONFIG_PATH", "./services/spawner/example.config.json")

func main() {
	shared.InitLogging()

	// Validate environment variables.
	if masterServerAddress == "" {
		panic("Must provide HOTEL_MASTER_ADDRESS environment variable.")
	}

	// Initialize main components.
	config := spawner.LoadConfig(configPath)
	controller := spawner.NewServerController(&config)
	service := spawner.NewSpawnerService(&config, controller)

	// Start the RPC server in a goroutine.
	addr := fmt.Sprintf(":%v", config.Port)
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
			err := masterClient.Register(config.Port, service.GetStatus())
			if err == nil {
				return
			}
			time.Sleep(5 * time.Second)
		}

		// Failed after allowed retries, abort.
		panic("Unable to contact master server after allowed retries, shutting down.")
	}()

	// Setup a SIGINT (CTRL+C) shutdown signal and block on it.
	shared.WaitForSigIntAndQuit()
}
