package master

import (
	"strings"
	"log"
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
	SendTestSpawnerRequest()
	json.NewEncoder(w).Encode(ok)
}

func HandleListGameServers(w http.ResponseWriter, r *http.Request) {
	gid := GameIDType(r.URL.Query().Get("gameId"))
	servers := DbGetGameServersByGameId(gid)
	var response ListServersResponse
	response.Servers = servers
	json.NewEncoder(w).Encode(response)
}

func HandleGetGameServer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// TODO: Full error object should not be returned in production mode
		http.Error(w, "Failed to parse request: "+err.Error(), http.StatusBadRequest)
		return
	}
	server, exists := DbGetGameServerById(ServerIDType(id))
	if !exists {
		http.Error(w, "Server by that ID not found.", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(server)
}

func HandleCreateGameServer(w http.ResponseWriter, r *http.Request) {
	session := context.Get(r, SessionContextKey).(Session)
	server, err := DecodeAndValidateGameServer(r.Body, false)
	fillImplicitGameServerFields(&server, r, session)
	if err != nil {
		// TODO: Full error object should not be returned in production mode
		http.Error(w, "Failed to parse request: "+err.Error(), http.StatusBadRequest)
		return
	}
	server, err = DbInsertGameServer(server)
	if err != nil {
		http.Error(w, "Failed to create server: "+err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(server)
}

func HandleUpdateGameServer(w http.ResponseWriter, r *http.Request) {
	session := context.Get(r, SessionContextKey).(Session)
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// TODO: Full error object should not be returned in production mode
		http.Error(w, "Failed to parse request URL: "+err.Error(), http.StatusBadRequest)
		return
	}
	serverId := ServerIDType(id)
	existingServer, exists := DbGetGameServerById(serverId)
	if !exists {
		http.Error(w, "Server by that ID not found.", http.StatusNotFound)
		return
	}
	if existingServer.SessionID != session.ID {
		http.Error(w, "Not authorized to modify that server.", http.StatusForbidden)
		return
	}
	newServer, err := DecodeAndValidateGameServer(r.Body, true)
	if err != nil {
		http.Error(w, "Failed to parse request body: "+err.Error(), http.StatusBadRequest)
		return
	}
	existingServer.Merge(newServer)
	newServer, err = DbUpdateGameServerById(serverId, existingServer)
	if err != nil {
		http.Error(w, "Failed to update server: "+err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(existingServer)
}

func fillImplicitGameServerFields(server *GameServer, r *http.Request, session Session) {
	server.SessionID = session.ID
	if len(server.Host) < 1 {
		clientAddr := r.Header.Get("X-Forwarded-For")
		if len(clientAddr) < 1 {
			clientAddr = r.RemoteAddr
		}
		clientAddr = strings.Split(clientAddr, ":")[0]
		// If the host was not provided, assume the host is the request origin.
		log.Printf("Host not specified, inferring from request origin: %v", r.RemoteAddr)
		server.Host = r.RemoteAddr
	}
}
