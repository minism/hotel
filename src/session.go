package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/context"
)

const (
	SessionContextKey = 0
)

// TODO: This should be externalized to filesystem or database at some point once
// we need to run multiple instances.
type SessionStore struct {
	// Mapping from session token to session data.
	Sessions map[string]Session
}

type Session struct {
	// The unique ID for the session.
	ID string

	// Servers owned by this session.
	Servers []ServerIDType
}

func (store *SessionStore) Initialize() {
	store.Sessions = make(map[string]Session)
}

func (store *SessionStore) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if session, exists := store.getSession(r); exists {
			context.Set(r, SessionContextKey, session)
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}

// Generate a session ID for future communication. The session ID is meant to not be
// guessable, and we rely on HTTPS (with certificate pinning) for reasonably secure
// communication and protection.
// The session ID is primarily used to manage resources created by a particular
// client. For example, a game server will register itself and generate a session,
// and only that server can update its status.
func (store *SessionStore) HandleIdentify(w http.ResponseWriter, r *http.Request) {
	var sessionId string
	if session, exists := store.getSession(r); exists {
		sessionId = session.ID
	} else {
		var err error
		sessionId, err = GenerateRandomB64String(32)
		if err != nil {
			log.Println("Error generating session ID: ", err)
			http.Error(w, "Failed to identify.", http.StatusBadRequest)
			return
		}
		store.createSession(sessionId)
	}
	json.NewEncoder(w).Encode(sessionId)
}

func (store *SessionStore) getSession(r *http.Request) (Session, bool) {
	var ret Session
	token := r.Header.Get("X-Session-Token")
	if session, exists := store.Sessions[token]; exists {
		return session, true
	}
	return ret, false
}

func (store *SessionStore) createSession(sessionId string) {
	session := Session{
		ID: sessionId,
	}
	store.Sessions[sessionId] = session
}
