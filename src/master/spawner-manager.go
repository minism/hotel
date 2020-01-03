package master

import (
	"log"
	"time"
)

var (
	spawners []Spawner = make([]Spawner, 0)
)

func RegisterSpawner(spawner Spawner) {
	log.Printf("Registered new spawner at %v:%v", spawner.Host, spawner.Port)
}

func InitSpawnerManager(config *Config) {
	// Start a routine which updates the status of spawners.
	go func() {
		for {
			time.Sleep(config.SpawnerCheckInterval.Duration)
		}
	}()
}
