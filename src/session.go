package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/context"
)

const (
	SessionContextKey = 0
)

var nextSessionId int = 1

// TODO: This should be externalized to filesystem or database at some point once
// we need to run multiple instances.
type SessionStore struct {
	// Mapping from session token to session data.
	Sessions map[string]Session
}

type Session struct {
	// The unique ID for the session.
	ID int

	// The token for the session.
	Token string

	// Servers owned by this session.
	Servers []ServerIDType

	// The last access time for the session.
	LastAccess time.Time
}

func NewSessionStore() *SessionStore {
	store := &SessionStore{}
	store.Sessions = make(map[string]Session)
	return store
}

func (store *SessionStore) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if session, exists := store.loadSessionForRequest(r); exists {
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
	var sessionToken string
	if session, exists := store.loadSessionForRequest(r); exists {
		sessionToken = session.Token
	} else {
		var err error
		sessionToken, err = GenerateRandomB32String(32)
		if err != nil {
			log.Println("Error generating session ID: ", err)
			http.Error(w, "Failed to identify.", http.StatusBadRequest)
			return
		}
		store.CreateSession(sessionToken)
	}

	var response IdentifyResponse
	response.Token = sessionToken
	json.NewEncoder(w).Encode(response)
}

func (store *SessionStore) CreateSession(sessionToken string) {
	session := Session{
		ID:    nextSessionId,
		Token: sessionToken,
	}
	nextSessionId++
	store.Sessions[sessionToken] = session
}

func (store *SessionStore) DeleteSession(sessionToken string) {
	delete(store.Sessions, sessionToken)
}

func (store *SessionStore) loadSessionForRequest(r *http.Request) (Session, bool) {
	var ret Session
	token := r.Header.Get("X-Session-Token")
	if session, exists := store.Sessions[token]; exists {
		session.LastAccess = time.Now()
		store.Sessions[token] = session
		return session, true
	}
	return ret, false
}
