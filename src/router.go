package main

import "github.com/gorilla/mux"

func createRouter() *mux.Router {
	var router = mux.NewRouter()
	ss := SessionStore{}
	ss.Initialize()

	// Unauthenticated routes.
	router.HandleFunc("/health", handleHealth).Methods("GET")

	// Authenticated routes.
	authRouter := router.PathPrefix("/").Subrouter()
	authRouter.Use(ss.Middleware)
	authRouter.HandleFunc("/servers", handleListServers).Methods("GET")
	authRouter.HandleFunc("/servers/{id}", handleGetServer).Methods("GET")
	authRouter.HandleFunc("/servers", handleCreateServer).Methods("POST")
	authRouter.HandleFunc("/servers/{id}", handleUpdateServer).Methods("PUT")
	authRouter.HandleFunc("/servers/{id}/alive", handleServerAlive).Methods("PUT")

	return router
}
