package main

import "github.com/gorilla/mux"

func createRouter(store *SessionStore) *mux.Router {
	var router = mux.NewRouter()

	// Unauthenticated routes.
	router.HandleFunc("/health", handleHealth).Methods("GET")
	router.HandleFunc("/identify", store.HandleIdentify).Methods("POST")

	// Authenticated routes.
	authRouter := router.PathPrefix("/").Subrouter()
	authRouter.Use(store.Middleware)
	authRouter.HandleFunc("/servers", handleListServers).Methods("GET")
	authRouter.HandleFunc("/servers/{id}", handleGetServer).Methods("GET")
	authRouter.HandleFunc("/servers", handleCreateServer).Methods("POST")
	authRouter.HandleFunc("/servers/{id}", handleUpdateServer).Methods("PUT")
	authRouter.HandleFunc("/servers/{id}/alive", handleServerAlive).Methods("PUT")

	return router
}
