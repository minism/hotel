package reaper

import (
	"log"
	"time"

	"minornine.com/hotel/src/master/config"
	"minornine.com/hotel/src/master/db"
	"minornine.com/hotel/src/master/session"
)

// StartReaper will kick off a goroutine for all reaper tasks.
func StartReaper(config *config.Config, store *session.SessionStore) {
	go func() {
		for {
			reapSessions(config, store)
			reapServers(config)
			time.Sleep(config.ReaperInterval.Duration)
		}
	}()
}

func reapSessions(config *config.Config, store *session.SessionStore) {
	for token, session := range store.Sessions {
		if time.Now().Sub(session.LastAccess) > config.SessionExpiration.Duration {
			store.DeleteSession(token)
		}
	}
}

func reapServers(config *config.Config) {
	oldestTime := time.Now().Add(-config.ServerExpiration.Duration)
	err := db.DeleteServersOlderThan(oldestTime.Unix())
	if err != nil {
		log.Println(err)
	}
}
