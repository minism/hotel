package master

import (
	"log"
	"time"

	"minornine.com/hotel/src/master/models"
)

func StartReaper(config *models.Config, store *SessionStore) {
	go func() {
		for {
			reapSessions(config, store)
			reapServers(config)
			time.Sleep(config.ReaperInterval.Duration)
		}
	}()
}

func reapSessions(config *models.Config, store *SessionStore) {
	for token, session := range store.Sessions {
		if time.Now().Sub(session.LastAccess) > config.SessionExpiration.Duration {
			store.DeleteSession(token)
		}
	}
}

func reapServers(config *models.Config) {
	oldestTime := time.Now().Add(-config.ServerExpiration.Duration)
	err := DeleteServersOlderThan(oldestTime.Unix())
	if err != nil {
		log.Println(err)
	}
}
