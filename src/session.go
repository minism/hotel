package main

import (
	"log"
	"net/http"
)

// TODO: This should be externalized to filesystem or database at some point once
// we need to run multiple instances.
type SessionStore struct {
	// Mapping from session token to session data.
	sessions map[string]Session
}

type Session struct {
	// Servers owned by this session.
	servers []ServerIDType
}

func (store *SessionStore) Initialize() {}

func (store *SessionStore) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-Session-Token")
		if session, exists := store.sessions[token]; exists {
			log.Println(session.servers)
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}
