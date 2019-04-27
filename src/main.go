package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	DEFAULT_PORT = 3000
)

func main() {
	var router = mux.NewRouter()

	router.HandleFunc("/health", handleHealth).Methods("GET")
	router.HandleFunc("/servers", handleListServers).Methods("GET")

	addr := fmt.Sprintf(":%v", DEFAULT_PORT)
	fmt.Println("Running server on", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
