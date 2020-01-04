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
	GameServerPath  string            `json:"gameSeverPath"`
}

func LoadConfig(configPath string) Config {
	var config Config
	shared.LoadConfigFromPath(configPath, &config)
	return config
}
