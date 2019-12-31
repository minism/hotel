package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"minornine.com/hotel/src/master"
	"minornine.com/hotel/src/shared"
)

const (
	DEFAULT_PORT = 3000
)

var dataPath = shared.GetEnv("HOTEL_DATA_PATH", ".")
var configPath = shared.GetEnv("HOTEL_CONFIG_PATH", "./default.config.json")

func main() {
	flag.Parse()
	shared.InitLogging()

	// Initialize main components.
	store := master.NewSessionStore()
	config := master.LoadConfig(configPath)
	master.InitDb(dataPath)
	master.StartReaper(config, store)

	// Run the HTTP server in a goroutine.
	addr := fmt.Sprintf(":%v", DEFAULT_PORT)
	mainRouter := handlers.LoggingHandler(os.Stdout, master.CreateRouter(store))
	log.Println("Running server on", addr)
	go func() {
		log.Fatal(http.ListenAndServe(addr, mainRouter))
	}()

	// Setup a SIGINT (CTRL+C) shutdown signal and block on it.
	c := shared.CreateSigintChannel()
	<-c
	log.Println("Shutting down.")
	os.Exit(0)
}
