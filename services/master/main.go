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

	addr := fmt.Sprintf(":%v", DEFAULT_PORT)
	mainRouter := handlers.LoggingHandler(os.Stdout, master.CreateRouter(store))

	// TODO: Run in goroutine with signal handling to not block
	// https://github.com/gorilla/mux
	log.Println("Running server on", addr)
	log.Fatal(http.ListenAndServe(addr, mainRouter))
}
