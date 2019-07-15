package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/context"
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
	server, exists := getServerById(ServerIDType(id))
	if !exists {
		http.Error(w, "Server by that ID not found.", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(server)
}

func handleCreateServer(w http.ResponseWriter, r *http.Request) {
	session := context.Get(r, SessionContextKey).(Session)
	server, err := DecodeAndValidateServer(r.Body)
	server.SessionID = session.ID
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
	session := context.Get(r, SessionContextKey).(Session)
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// TODO: Full error object should not be returned in production mode
		http.Error(w, "Failed to parse request URL: "+err.Error(), http.StatusBadRequest)
		return
	}
	serverId := ServerIDType(id)
	server, exists := getServerById(serverId)
	if !exists {
		http.Error(w, "Server by that ID not found.", http.StatusNotFound)
		return
	}
	if server.SessionID != session.ID {
		http.Error(w, "Not authorized to modify that server.", http.StatusForbidden)
		return
	}
	newServer, err := DecodeAndValidateServer(r.Body)
	if err != nil {
		http.Error(w, "Failed to parse request body: "+err.Error(), http.StatusBadRequest)
		return
	}
	newServer, err = updateServerById(serverId, newServer)
	if err != nil {
		http.Error(w, "Failed to update server: "+err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(newServer)
}

func handleServerAlive(w http.ResponseWriter, r *http.Request) {
	session := context.Get(r, SessionContextKey).(Session)
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// TODO: Full error object should not be returned in production mode
		http.Error(w, "Failed to parse request URL: "+err.Error(), http.StatusBadRequest)
		return
	}
	serverId := ServerIDType(id)
	server, exists := getServerById(serverId)
	if !exists {
		http.Error(w, "Server by that ID not found.", http.StatusNotFound)
		return
	}
	if server.SessionID != session.ID {
		http.Error(w, "Not authorized to modify that server.", http.StatusForbidden)
		return
	}
	updateServerAlive(serverId)
	json.NewEncoder(w).Encode(ok)
}
