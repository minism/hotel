package main

import (
	"log"
	"time"
)

var reaperInterval time.Duration = time.Second * 5
var sessionAccessExpiry time.Duration = time.Minute * 15
var gameServerAccessExpiry time.Duration = time.Second * 30

func initReaper(store *SessionStore) {
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
	for key, session := range store.Sessions {
		if time.Now().Sub(session.LastAccess) > sessionAccessExpiry {
			delete(store.Sessions, key)
		}
	}
}

func reapServers() {
	oldestTime := time.Now().Add(-gameServerAccessExpiry)
	err := deleteServersOlderThan(oldestTime.Unix())
	if err != nil {
		log.Println(err)
	}
}
