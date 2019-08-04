package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

const (
	DEFAULT_PORT = 3000
)

var configPath = flag.String("config_path", "./example.config.json", "Path to configuration file.")
var dataPath = flag.String("data_path", ".", "Path to directory for storing data.")

func main() {
	flag.Parse()

	// Global configuration.
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Initialize main components.
	store := NewSessionStore()
	config := LoadConfig(*configPath)
	InitDb(*dataPath)
	StartReaper(config, store)

	addr := fmt.Sprintf(":%v", DEFAULT_PORT)
	mainRouter := handlers.LoggingHandler(os.Stdout, CreateRouter(store))

	// TODO: Run in goroutine with signal handling to not block
	// https://github.com/gorilla/mux
	log.Println("Running server on", addr)
	log.Fatal(http.ListenAndServe(addr, mainRouter))
}
