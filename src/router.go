package main

import "github.com/gorilla/mux"

func CreateRouter(store *SessionStore) *mux.Router {
	var router = mux.NewRouter()

	// Unauthenticated routes.
	router.HandleFunc("/health", HandleHealth).Methods("GET")
	router.HandleFunc("/identify", store.HandleIdentify).Methods("POST")

	// Authenticated routes.
	authRouter := router.PathPrefix("/").Subrouter()
	authRouter.Use(store.Middleware)
	authRouter.HandleFunc("/servers", HandleListServers).Methods("GET")
	authRouter.HandleFunc("/servers/{id}", HandleGetServer).Methods("GET")
	authRouter.HandleFunc("/servers", HandleCreateServer).Methods("POST")
	authRouter.HandleFunc("/servers/{id}", HandleUpdateServer).Methods("PUT")
	authRouter.HandleFunc("/servers/{id}/alive", HandleServerAlive).Methods("PUT")

	return router
}
