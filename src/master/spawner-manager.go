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

func GetAvailableSpawnerForGame(gameId shared.GameIDType) (Spawner, error) {
	var ret Spawner
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
	ret = spawners[0]
	if ret.Capacity() < 1 {
		return ret, errors.New(fmt.Sprintf("No capacity left for game ID '%v'", gameId))
	}

	return ret, nil
}

func SpawnServerForGame(spawner Spawner, gameId shared.GameIDType) (GameServer, error) {
	var ret GameServer

	// RPC to request a game server spawn.
	response, err := SendSpawnServerRequest(&spawner)
	if err != nil {
		log.Printf("Error making spawn RPC: %v", err)
		return ret, err
	}

	// We return a partially filled GameServer struct, which at a minimum will have host:port
	// for the client to connect to, because the spawner will know about this.
	// We don't have the full struct including ID because it wont be generated until the
	// game server itself starts up and communicates with the master server.
	ret.Host = response.Host
	if ret.Host == "" {
		ret.Host = spawner.Host
	}
	ret.Port = int(response.Port)
	ret.GameID = spawner.GameID
	return ret, nil
}

func InitSpawnerManager(config *Config) {
	// Start a routine which updates the status of spawners.
	go func() {
		// TODO: Make this a count query instead.
		var spawners = DbGetSpawners()
		log.Printf("Discovered %v existing spawners in database.", len(spawners))
		for {
			for _, spawner := range DbGetSpawners() {
				status, err := SendCheckStatusRequest(&spawner)
				if err != nil {
					log.Printf("Error checking status of spawner at %v, removing from pool.", spawner.Address())
					DbDeleteSpawnerById(spawner.ID)
				} else {
					DbUpdateSpawnerFromStatus(spawner.ID, status)
				}
			}
			time.Sleep(config.SpawnerCheckInterval.Duration)
		}
	}()
}
