package session

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/context"
	"minornine.com/hotel/src/master/models"
	"minornine.com/hotel/src/shared"
)

const (
	// SessionContextKey is the key where session is stored in the HTTP request context.
	SessionContextKey = 0
)

// Session represents a single user session (currently kept in memory).
type Session struct {
	// The unique ID for the session.
	ID int

	// The token for the session.
	Token string

	// Servers owned by this session.
	Servers map[models.ServerIDType]bool

	// The last access time for the session.
	LastAccess time.Time
}

// OwnsServerId returns whether the given server ID is owned by this session.
func (session *Session) OwnsServerId(id models.ServerIDType) bool {
	if ok, exists := session.Servers[id]; exists {
		return ok
	}
	return false
}

var nextSessionId int = 1

// SessionStore is a singleton which stores all session data.
// TODO: This should be externalized to filesystem or database at some point once
// we need to run multiple instances.
type SessionStore struct {
	// Mapping from session token to session data.
	Sessions map[string]Session
}

// NewSessionStore creates and initializes a session store.
func NewSessionStore() *SessionStore {
	store := &SessionStore{}
	store.Sessions = make(map[string]Session)
	return store
}

// Middleware creates a middleware function which injects correct session into the HTTP context.
// If no session could be loaded for the request, HTTP 403 is returned instead.
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

// HandleIdentify generates a session ID for future communication.
// The session ID is meant to not be guessable, and we rely on HTTPS
// (with certificate pinning) for reasonably secure communication and protection.
// The session ID is primarily used to manage resources created by a particular
// client. For example, a game server will register itself and generate a session,
// and only that server can update its status.
// TODO: For session expiry, this should return a token alias instead of echo.
func (store *SessionStore) HandleIdentify(w http.ResponseWriter, r *http.Request) {
	var sessionToken string
	if session, exists := store.loadSessionForRequest(r); exists {
		sessionToken = session.Token
	} else {
		var err error
		sessionToken, err = shared.GenerateRandomB32String(32)
		if err != nil {
			log.Println("Error generating session ID: ", err)
			http.Error(w, "Failed to identify.", http.StatusBadRequest)
			return
		}
		store.CreateSession(sessionToken)
	}

	var response models.IdentifyResponse
	response.Token = sessionToken
	json.NewEncoder(w).Encode(response)
}

// CreateSession creates a session for the given token and stores it.
func (store *SessionStore) CreateSession(sessionToken string) {
	session := Session{
		ID:      nextSessionId,
		Token:   sessionToken,
		Servers: make(map[models.ServerIDType]bool),
	}
	nextSessionId++
	store.Sessions[sessionToken] = session
}

// DeleteSession delets the session associated with the given token.
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
