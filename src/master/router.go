package master

import "github.com/gorilla/mux"

func CreateRouter(store *SessionStore) *mux.Router {
	var router = mux.NewRouter()

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

	return router
}
