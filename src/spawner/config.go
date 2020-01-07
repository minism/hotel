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

	// Port that the spawner is running on, not loaded from the config file.
	Port uint32
}

// LoadConfig loads and returns a spawner config from the given path.
func LoadConfig(configPath string) Config {
	var config Config
	shared.LoadConfigFromPath(configPath, &config)
	return config
}
