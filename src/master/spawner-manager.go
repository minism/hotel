package master

import (
	"errors"
	"fmt"
	"log"
	"time"

	"minornine.com/hotel/src/shared"
)

func RegisterSpawner(spawner Spawner) {
	log.Printf("Registered new spawner at %v:%v", spawner.Host, spawner.Port)
}

func SpawnServerForGame(gameId shared.GameIDType) (GameServer, error) {
	var ret GameServer
	spawners := DbGetSpawnersByGameId(gameId)
	if len(spawners) < 1 {
		return ret, errors.New(fmt.Sprintf("No spawners available for game ID '%v'", gameId))
	}
	return ret, nil
}

func InitSpawnerManager(config *Config) {
	// Start a routine which updates the status of spawners.
	go func() {
		for {
			time.Sleep(config.SpawnerCheckInterval.Duration)
		}
	}()
}
