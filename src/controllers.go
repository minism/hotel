package main

import (
	"encoding/json"
	"net/http"
)

func handleHealth(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("OK")
}

func handleListServers(w http.ResponseWriter, r *http.Request) {
	servers := make([]GameServer, 0)
	servers = append(servers, GameServer{Name: "Test"})
	json.NewEncoder(w).Encode(servers)
}
