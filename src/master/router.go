package master

import (
	"github.com/gorilla/mux"
	"minornine.com/hotel/src/master/models"
)

func CreateRouter(config *models.Config, store *SessionStore) *mux.Router {
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

	return router
}
