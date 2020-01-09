package spawner

import (
	"minornine.com/hotel/src/shared"
)

// Config contains global configuration for the hotel spawner instance.
// The configuration is loaded from a JSON file provided at runtime.
type Config struct {
	SupportedGameID shared.GameIDType `json:"supportedGameId"`
	MaxGameServers  uint32            `json:"maxGameServers"`
	GameServerPath  string            `json:"gameServerPath"`
	GameServerFlags []string          `json:"gameServerFlags"`

	// Whether the spawner should start game servers immediately, versus
	// waiting for clients to request spawns.
	Autorun bool `json:"autorun"`

	// The internet hostname that the spawner and gameservers are running on.
	// This can be left empty, and the master server will infer host automatically.
	// However, for a single docker network stack, this is a bad idea because
	// the master server will discover the docker IP for the spawner rather than
	// the public IP. So, the best practice is to always set this.
	Host string

	// Port that the spawner is running on, not loaded from the config file.
	// At the moment this is hardcoded since we expect the spawner process
	// to be ran inside a docker container, where the port can be remapped.
	Port uint32
}

// LoadConfig loads and returns a spawner config from the given path.
func LoadConfig(configPath string) Config {
	var config Config
	shared.LoadConfigFromPath(configPath, &config)
	return config
}
