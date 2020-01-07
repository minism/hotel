package master

import (
	"log"
	"time"

	"minornine.com/hotel/src/master/db"
)

// StartReaper will kick off a goroutine for all reaper tasks.
func StartReaper(config *Config, store *SessionStore) {
	go func() {
		for {
			reapSessions(config, store)
			reapServers(config)
			time.Sleep(config.ReaperInterval.Duration)
		}
	}()
}

func reapSessions(config *Config, store *SessionStore) {
	for token, session := range store.Sessions {
		if time.Now().Sub(session.LastAccess) > config.SessionExpiration.Duration {
			store.DeleteSession(token)
		}
	}
}

func reapServers(config *Config) {
	oldestTime := time.Now().Add(-config.ServerExpiration.Duration)
	err := db.DeleteServersOlderThan(oldestTime.Unix())
	if err != nil {
		log.Println(err)
	}
}
