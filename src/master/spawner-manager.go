package master

import (
	"errors"
	"fmt"
	"log"
	"sort"
	"time"

	"minornine.com/hotel/src/shared"
)

func RegisterSpawner(spawner Spawner) {
	err := DbInsertSpawner(spawner)
	if err != nil {
		log.Printf("Error registering spawner: %v", err)
	} else {
		log.Printf("Registered new spawner at %v:%v", spawner.Host, spawner.Port)
	}
}

func SpawnServerForGame(gameId shared.GameIDType) (GameServer, error) {
	var ret GameServer
	spawners := DbGetSpawnersByGameId(gameId)
	if len(spawners) < 1 {
		return ret, errors.New(fmt.Sprintf("No spawners available for game ID '%v'", gameId))
	}

	// Implement basic load balancing by sorting spawners by capacity.
	// TODO: Should sort by done in SQL?
	sort.Slice(spawners, func(i, j int) bool {
		return (spawners[i].Capacity() < spawners[j].Capacity())
	})

	// Ensure there is at least some capacity.
	spawner := spawners[0]
	if spawner.Capacity() < 1 {
		return ret, errors.New(fmt.Sprintf("No capacity left for game ID '%v'", gameId))
	}

	// RPC to request a spawn.

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
