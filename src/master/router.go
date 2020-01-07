package master

import (
	"github.com/gorilla/mux"
)

// CreateRouter will create the main mux router used by the master HTTP service.
func CreateRouter(config *Config, store *SessionStore) *mux.Router {
	var router = mux.NewRouter()
	router.Use(config.Middleware)

	// Unauthenticated routes.
	router.HandleFunc("/health", HandleHealth).Methods("GET")
	router.HandleFunc("/identify", store.HandleIdentify).Methods("POST")

	// Authenticated routes.
	authRouter := router.PathPrefix("/").Subrouter()
	authRouter.Use(store.Middleware)
	authRouter.HandleFunc("/servers", HandleListGameServers).Methods("GET")
	authRouter.HandleFunc("/servers/{id}", HandleGetGameServer).Methods("GET")
	authRouter.HandleFunc("/servers", HandleCreateGameServer).Methods("POST")
	authRouter.HandleFunc("/servers/{id}", HandleUpdateGameServer).Methods("PUT")
	authRouter.HandleFunc("/servers/{id}", HandleDeleteGameServer).Methods("DELETE")
	authRouter.HandleFunc("/spawn", HandleSpawnGameServer).Methods("POST")

	return router
}
