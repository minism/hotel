package router

import (
	"github.com/gorilla/mux"
	"github.com/minism/hotel/src/master/config"
	"github.com/minism/hotel/src/master/controllers"
	"github.com/minism/hotel/src/master/session"
)

// CreateRouter will create the main mux router used by the master HTTP service.
func CreateRouter(config *config.Config, store *session.SessionStore) *mux.Router {
	var router = mux.NewRouter()
	router.Use(config.Middleware)

	// Unauthenticated routes.
	router.HandleFunc("/health", controllers.CheckHealth).Methods("GET")
	router.HandleFunc("/identify", store.HandleIdentify).Methods("POST")

	// Authenticated routes.
	authRouter := router.PathPrefix("/").Subrouter()
	authRouter.Use(store.Middleware)
	authRouter.HandleFunc("/servers", controllers.ListGameServers).Methods("GET")
	authRouter.HandleFunc("/servers/{id}", controllers.GetGameServer).Methods("GET")
	authRouter.HandleFunc("/servers", controllers.CreateGameServer).Methods("POST")
	authRouter.HandleFunc("/servers/{id}", controllers.UpdateGameServer).Methods("PUT")
	authRouter.HandleFunc("/servers/{id}", controllers.DeleteGameServer).Methods("DELETE")
	authRouter.HandleFunc("/spawn", controllers.SpawnGameServer).Methods("POST")

	return router
}
