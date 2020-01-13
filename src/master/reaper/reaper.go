package reaper

import (
	"log"
	"time"

	"github.com/minism/hotel/src/master/config"
	"github.com/minism/hotel/src/master/db"
	"github.com/minism/hotel/src/master/rpc"
	"github.com/minism/hotel/src/master/session"
)

// StartReaper will kick off a goroutine for all reaper tasks.
func StartReaper(config *config.Config, store *session.SessionStore) {
	go func() {
		for {
			reapSessions(config, store)
			reapServers(config)
			reapSpawners(config)
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

func reapSpawners(config *config.Config) {
	for _, spawner := range db.GetSpawners() {
		status, err := rpc.SendCheckStatusRequest(spawner)
		if err != nil {
			log.Printf("Error checking status of spawner at %v, removing from pool.", spawner.Address())
			db.DeleteSpawnerById(spawner.ID)
		} else {
			db.UpdateSpawnerFromStatus(spawner.ID, status)
		}
	}
}
