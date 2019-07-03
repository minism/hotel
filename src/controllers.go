package main

import (
	"encoding/json"
	"net/http"
	"strconv"

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
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// TODO: Full error object should not be returned in production mode
		http.Error(w, "Failed to parse request: "+err.Error(), http.StatusBadRequest)
		return
	}
	server := getServerById(ServerIDType(id))
	json.NewEncoder(w).Encode(server)
}

func handleCreateServer(w http.ResponseWriter, r *http.Request) {
	server, err := DecodeAndValidateServer(r.Body)
	if err != nil {
		// TODO: Full error object should not be returned in production mode
		http.Error(w, "Failed to parse request: "+err.Error(), http.StatusBadRequest)
		return
	}
	server, err = insertServer(server)
	if err != nil {
		http.Error(w, "Failed to create server: "+err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(server)
}

func handleUpdateServer(w http.ResponseWriter, r *http.Request) {
	server, err := DecodeAndValidateServer(r.Body)
	if err != nil {
		http.Error(w, "Failed to parse request: "+err.Error(), http.StatusBadRequest)
		return
	}
	server, err = updateServerById(server.ID, server)
	if err != nil {
		http.Error(w, "Failed to create server: "+err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(server)
}

func handleServerAlive(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// TODO: Full error object should not be returned in production mode
		http.Error(w, "Failed to parse request: "+err.Error(), http.StatusBadRequest)
		return
	}
	pingServerAlive(ServerIDType(id))
	json.NewEncoder(w).Encode(ok)
}
