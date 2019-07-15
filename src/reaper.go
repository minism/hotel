package main

import (
	"log"
	"time"
)

func InitReaper(config Config, store *SessionStore) {
	go func() {
		for {
			reapSessions(config, store)
			reapServers(config)
			time.Sleep(config.ReaperInterval.Duration)
		}
	}()
}

func reapSessions(config Config, store *SessionStore) {
	log.Println("Checking state to reap, num sessions: ", len(store.Sessions))
	for token, session := range store.Sessions {
		if time.Now().Sub(session.LastAccess) > config.SessionExpiration.Duration {
			store.DeleteSession(token)
		}
	}
}

func reapServers(config Config) {
	oldestTime := time.Now().Add(-config.ServerExpiration.Duration)
	err := DeleteServersOlderThan(oldestTime.Unix())
	if err != nil {
		log.Println(err)
	}
}
