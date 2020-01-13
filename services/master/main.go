package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/minism/hotel/src/master/config"
	"github.com/minism/hotel/src/master/db"
	"github.com/minism/hotel/src/master/reaper"
	"github.com/minism/hotel/src/master/router"
	"github.com/minism/hotel/src/master/rpc"
	"github.com/minism/hotel/src/master/session"
	hotel_pb "github.com/minism/hotel/src/proto"
	"github.com/minism/hotel/src/shared"
)

const (
	DEFAULT_HTTP_PORT = 3000
	DEFAULT_RPC_PORT  = 3001
)

var dataPath = shared.GetEnv("HOTEL_DATA_PATH", ".")
var configPath = shared.GetEnv("HOTEL_CONFIG_PATH", "./services/master/example.config.json")

func main() {
	flag.Parse()
	shared.InitLogging()

	// Initialize main components.
	store := session.NewSessionStore()
	config := config.LoadConfig(configPath)
	db.InitDb(dataPath)
	reaper.StartReaper(config, store)

	// Start the HTTP server in a goroutine.
	httpAddr := fmt.Sprintf(":%v", DEFAULT_HTTP_PORT)
	mainRouter := handlers.LoggingHandler(os.Stdout, router.CreateRouter(config, store))
	log.Println("Running HTTP server on", httpAddr)
	go func() {
		log.Fatal(http.ListenAndServe(httpAddr, mainRouter))
	}()

	// Start the RPC server in a goroutine.
	rpcAddr := fmt.Sprintf(":%v", DEFAULT_RPC_PORT)
	tcpListener, err := net.Listen("tcp", rpcAddr)
	if err != nil {
		panic(fmt.Sprintf("Error binding TCP socket to %v", rpcAddr))
	}
	grpcServer := grpc.NewServer()
	hotel_pb.RegisterMasterServiceServer(grpcServer, &rpc.MasterService{})
	log.Println("Running RPC server on", rpcAddr)
	go func() {
		log.Fatal(grpcServer.Serve(tcpListener))
	}()

	// Setup a SIGINT (CTRL+C) shutdown signal and block on it.
	shared.WaitForSigIntAndQuit()
}
