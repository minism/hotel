package main

import (
	"encoding/json"
	"net/http"
)

func handleHealth(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("OK")
}

func handleGetLobbies(w http.ResponseWriter, r *http.Request) {
	lobbies := make([]int, 0)
	json.NewEncoder(w).Encode(lobbies)
}
