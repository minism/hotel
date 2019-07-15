package main

import (
	"log"
	"time"
)

var reaperInterval time.Duration = time.Second * 5
var sessionAccessExpiry time.Duration = time.Minute * 15
var gameServerAccessExpiry time.Duration = time.Second * 30

func InitReaper(store *SessionStore) {
	go func() {
		for {
			reapSessions(store)
			reapServers()
			time.Sleep(reaperInterval)
		}
	}()
}

func reapSessions(store *SessionStore) {
	log.Println("Checking state to reap, num sessions: ", len(store.Sessions))
	for token, session := range store.Sessions {
		if time.Now().Sub(session.LastAccess) > sessionAccessExpiry {
			store.DeleteSession(token)
		}
	}
}

func reapServers() {
	oldestTime := time.Now().Add(-gameServerAccessExpiry)
	err := DeleteServersOlderThan(oldestTime.Unix())
	if err != nil {
		log.Println(err)
	}
}
