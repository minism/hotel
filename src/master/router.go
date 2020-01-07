package master

import (
	"github.com/gorilla/mux"
	"minornine.com/hotel/src/master/config"
	"minornine.com/hotel/src/master/controllers"
	"minornine.com/hotel/src/master/session"
)

// CreateRouter will create the main mux router used by the master HTTP service.
func CreateRouter(config *config.Config, store *session.SessionStore) *mux.Router {
	var router = mux.NewRouter()
	router.Use(config.Middleware)

	// Unauthenticated routes.
	router.HandleFunc("/health", controllers.HandleHealth).Methods("GET")
	router.HandleFunc("/identify", store.HandleIdentify).Methods("POST")

	// Authenticated routes.
	authRouter := router.PathPrefix("/").Subrouter()
	authRouter.Use(store.Middleware)
	authRouter.HandleFunc("/servers", controllers.HandleListGameServers).Methods("GET")
	authRouter.HandleFunc("/servers/{id}", controllers.HandleGetGameServer).Methods("GET")
	authRouter.HandleFunc("/servers", controllers.HandleCreateGameServer).Methods("POST")
	authRouter.HandleFunc("/servers/{id}", controllers.HandleUpdateGameServer).Methods("PUT")
	authRouter.HandleFunc("/servers/{id}", controllers.HandleDeleteGameServer).Methods("DELETE")
	authRouter.HandleFunc("/spawn", controllers.HandleSpawnGameServer).Methods("POST")

	return router
}
