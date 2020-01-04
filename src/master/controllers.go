package master

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"minornine.com/hotel/src/shared"
)

const (
	ok = "OK"
)

func HandleHealth(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(ok)
}

func HandleListGameServers(w http.ResponseWriter, r *http.Request) {
	gid := shared.GameIDType(r.URL.Query().Get("gameId"))
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
	config := context.Get(r, ConfigContextKey).(*Config)
	session := context.Get(r, SessionContextKey).(Session)
	server, err := DecodeAndValidateGameServer(r.Body, false)
	fillImplicitGameServerFields(&server, r, session)
	if err != nil {
		// TODO: Full error object should not be returned in production mode
		http.Error(w, "Failed to parse request: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Check if game definition allows server creation.
	def, ok := config.GetGameDefinition(server.GameID)
	if !ok {
		http.Error(w, fmt.Sprintf("No definition for game ID '%v', and this master server doesn't allow undefined games.", server.GameID), http.StatusBadRequest)
		return
	}
	if def.HostPolicy != shared.HostPolicy_ANY && def.HostPolicy != shared.HostPolicy_CLIENTS_ONLY {
		http.Error(w, fmt.Sprintf("Game '%v' doesn't allow client server creation.", server.GameID), http.StatusBadRequest)
		return
	}

	// Db call.
	server, err = DbInsertGameServer(server)
	if err != nil {
		http.Error(w, "Failed to create server: "+err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(server)
}

func HandleSpawnGameServer(w http.ResponseWriter, r *http.Request) {
	config := context.Get(r, ConfigContextKey).(*Config)
	// session := context.Get(r, SessionContextKey).(Session)
	gid := shared.GameIDType(r.URL.Query().Get("gameId"))

	// Check if game definition allows spawning.
	def, ok := config.GetGameDefinition(gid)
	if !ok {
		http.Error(w, fmt.Sprintf("No definition for game ID '%v', and this master server doesn't allow undefined games.", gid), http.StatusBadRequest)
		return
	}
	if def.HostPolicy != shared.HostPolicy_ANY && def.HostPolicy != shared.HostPolicy_SPAWN_ONLY {
		http.Error(w, fmt.Sprintf("Game '%v' doesn't allow spawning.", gid), http.StatusBadRequest)
		return
	}

	server, err := SpawnServerForGame(gid)
	if err != nil {
		http.Error(w, "Failed to spawn server (internal error).", http.StatusInternalServerError)
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
		http.Error(w, "Failed to update server: "+err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(existingServer)
}

func HandleDeleteGameServer(w http.ResponseWriter, r *http.Request) {
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

	err = DbDeleteGameServerById(serverId)
	if err != nil {
		http.Error(w, "Failed to delete server: "+err.Error(), http.StatusInternalServerError)
	}
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
