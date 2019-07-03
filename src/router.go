package main

import "github.com/gorilla/mux"

func createRouter() *mux.Router {
	var router = mux.NewRouter()

	router.HandleFunc("/health", handleHealth).Methods("GET")
	router.HandleFunc("/servers", handleListServers).Methods("GET")
	router.HandleFunc("/servers/{id}", handleGetServer).Methods("GET")
	router.HandleFunc("/servers", handleCreateServer).Methods("POST")
	router.HandleFunc("/servers/{id}", handleUpdateServer).Methods("PUT")
	router.HandleFunc("/servers/{id}/alive", handleServerAlive).Methods("PUT")

	return router
}
