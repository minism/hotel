package main

import "github.com/gorilla/mux"

func createRouter() *mux.Router {
	var router = mux.NewRouter()

	router.HandleFunc("/health", handleHealth).Methods("GET")
	router.HandleFunc("/servers/{gameId}", handleListServers).Methods("GET")
	return router
}
