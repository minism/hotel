package master

import (
	"fmt"
	"log"
	"sort"
	"time"

	"minornine.com/hotel/src/master/db"
	"minornine.com/hotel/src/master/models"
	"minornine.com/hotel/src/shared"
)

// RegisterSpawner adds the given spawner to the database.
func RegisterSpawner(spawner models.Spawner) {
	err := db.DbInsertSpawner(spawner)
	if err != nil {
		log.Printf("Error registering spawner: %v", err)
	} else {
		log.Printf("Registered new spawner at %v:%v", spawner.Host, spawner.Port)
	}
}

// GetAvailableSpawnerForGame returns the best available spawner for the given game ID.
// Spawners are attempted to be load balanced by capacity, so this function should return
// a spawner with the most capacity.
func GetAvailableSpawnerForGame(gameId shared.GameIDType) (models.Spawner, error) {
	var ret models.Spawner
	spawners := db.DbGetSpawnersByGameId(gameId)
	if len(spawners) < 1 {
		return ret, fmt.Errorf("No spawners available for game ID '%v'", gameId)
	}

	// Implement basic load balancing by sorting spawners by capacity.
	// TODO: Should sort by done in SQL?
	sort.Slice(spawners, func(i, j int) bool {
		return (spawners[i].Capacity() > spawners[j].Capacity())
	})

	// Ensure there is at least some capacity.
	ret = spawners[0]
	if ret.Capacity() < 1 {
		return ret, fmt.Errorf("No capacity left for game ID '%v'", gameId)
	}

	return ret, nil
}

// SpawnServerForGame asks the given spawner to spawn a game server for the given game ID.
// If successful, a GameServer representing the newly running server is returned (or the
// expected game server that is being started).
func SpawnServerForGame(spawner models.Spawner, gameId shared.GameIDType) (models.GameServer, error) {
	var ret models.GameServer

	// RPC to request a game server spawn.
	response, err := SendSpawnServerRequest(&spawner)
	if err != nil {
		log.Printf("Error making spawn RPC: %v", err)
		return ret, err
	}

	// Update the spawner status in the DB immediately.
	db.DbUpdateSpawnerFromStatus(spawner.ID, response.Status)

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

// InitSpawnerManager kicks off a goroutine which cleans up dead spawners.
// TODO: Should be in reaper?
func InitSpawnerManager(config *Config) {
	// Start a routine which updates the status of spawners.
	go func() {
		// TODO: Make this a count query instead.
		var spawners = db.DbGetSpawners()
		log.Printf("Discovered %v existing spawners in database.", len(spawners))
		for {
			for _, spawner := range db.DbGetSpawners() {
				status, err := SendCheckStatusRequest(&spawner)
				if err != nil {
					log.Printf("Error checking status of spawner at %v, removing from pool.", spawner.Address())
					db.DbDeleteSpawnerById(spawner.ID)
				} else {
					db.DbUpdateSpawnerFromStatus(spawner.ID, status)
				}
			}
			time.Sleep(config.SpawnerCheckInterval.Duration)
		}
	}()
}
