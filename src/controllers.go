package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	ok = "OK"
)

func handleHealth(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(ok)
}

func handleListServers(w http.ResponseWriter, r *http.Request) {
	gid := GameIDType(r.URL.Query().Get("gameId"))
	servers := getServersByGameId(gid)
	json.NewEncoder(w).Encode(servers)
}

func handleGetServer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := ServerIDType(vars["id"])
	server := getServerById(id)
	json.NewEncoder(w).Encode(server)
}

func handleCreateServer(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(ok)
}

func handleUpdateServer(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(ok)
}

func handleServerAlive(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := ServerIDType(vars["id"])
	pingServerAlive(id)
	json.NewEncoder(w).Encode(ok)
}
