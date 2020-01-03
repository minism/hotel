package master

import (
	"time"
)

var (
	spawners []Spawner = make([]Spawner, 0)
)

type Spawner struct {
	Host string
	Port uint32
}

func RegisterSpawner(spawner Spawner) {

}

func InitSpawnerManager(config *Config) {
	// Start a routine which updates the status of spawners.
	go func() {
		for {
			time.Sleep(config.SpawnerCheckInterval.Duration)
		}
	}()
}
