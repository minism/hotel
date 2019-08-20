package master

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

func HandleHealth(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(ok)
}

func HandleListServers(w http.ResponseWriter, r *http.Request) {
	gid := GameIDType(r.URL.Query().Get("gameId"))
	servers := GetServersByGameId(gid)
	var response ListServersResponse
	response.Servers = servers
	json.NewEncoder(w).Encode(response)
}

func HandleGetServer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// TODO: Full error object should not be returned in production mode
		http.Error(w, "Failed to parse request: "+err.Error(), http.StatusBadRequest)
		return
	}
	server, exists := GetServerById(ServerIDType(id))
	if !exists {
		http.Error(w, "Server by that ID not found.", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(server)
}

func HandleCreateServer(w http.ResponseWriter, r *http.Request) {
	session := context.Get(r, SessionContextKey).(Session)
	server, err := DecodeAndValidateServer(r.Body)
	server.SessionID = session.ID
	if err != nil {
		// TODO: Full error object should not be returned in production mode
		http.Error(w, "Failed to parse request: "+err.Error(), http.StatusBadRequest)
		return
	}
	server, err = InsertServer(server)
	if err != nil {
		http.Error(w, "Failed to create server: "+err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(server)
}

func HandleUpdateServer(w http.ResponseWriter, r *http.Request) {
	session := context.Get(r, SessionContextKey).(Session)
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// TODO: Full error object should not be returned in production mode
		http.Error(w, "Failed to parse request URL: "+err.Error(), http.StatusBadRequest)
		return
	}
	serverId := ServerIDType(id)
	server, exists := GetServerById(serverId)
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
	newServer, err = UpdateServerById(serverId, newServer)
	if err != nil {
		http.Error(w, "Failed to update server: "+err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(newServer)
}

func HandleServerAlive(w http.ResponseWriter, r *http.Request) {
	session := context.Get(r, SessionContextKey).(Session)
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// TODO: Full error object should not be returned in production mode
		http.Error(w, "Failed to parse request URL: "+err.Error(), http.StatusBadRequest)
		return
	}
	serverId := ServerIDType(id)
	server, exists := GetServerById(serverId)
	if !exists {
		http.Error(w, "Server by that ID not found.", http.StatusNotFound)
		return
	}
	if server.SessionID != session.ID {
		http.Error(w, "Not authorized to modify that server.", http.StatusForbidden)
		return
	}
	UpdateServerAlive(serverId)
	json.NewEncoder(w).Encode(ok)
}