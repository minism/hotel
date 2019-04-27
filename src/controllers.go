package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func handleHealth(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("OK")
}

func handleListServers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gid := GameIDType(vars["gameId"])
	servers := getServersByGameId(gid)
	json.NewEncoder(w).Encode(servers)
}
