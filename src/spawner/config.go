package spawner

import (
	"minornine.com/hotel/src/shared"
)

const (
	ConfigContextKey = 1
)

type Config struct {
	SupportedGameID shared.GameIDType `json:"supportedGameId"`
	MaxGameServers  uint32            `json:"maxGameServers"`
	GameServerPath  string            `json:"gameServerPath"`

	// Whether the spawner should start game servers immediately, versus
	// waiting for clients to request spawns.
	Autorun bool `json:"autorun"`

	// Port that the spawner is running on, not loaded from the config file.
	Port uint32
}

func LoadConfig(configPath string) Config {
	var config Config
	shared.LoadConfigFromPath(configPath, &config)
	return config
}
