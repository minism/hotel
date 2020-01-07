package config

import (
	"log"
	"net/http"

	"github.com/gorilla/context"
	"minornine.com/hotel/src/shared"
)

const (
	// ConfigContextKey is the key where config is stored in the HTTP request context.
	ConfigContextKey = 1
)

// Config contains global configuration for the hotel master instance.
// The configuration is loaded from a JSON file provided at runtime.
type Config struct {
	ReaperInterval       shared.SerializableDuration `json:"reaperInterval"`
	SessionExpiration    shared.SerializableDuration `json:"sessionExpiration"`
	ServerExpiration     shared.SerializableDuration `json:"serverExpiration"`
	SpawnerCheckInterval shared.SerializableDuration `json:"spawnerCheckInterval"`
	GameDefs             []shared.GameDefinition     `json:"gameDefs"`
	AllowUndefinedGames  bool                        `json:"allowUndefinedGames"`

	gameDefsById map[shared.GameIDType]shared.GameDefinition
}

// Middleware creates an HTTP middleware which injects config into the context.
func (config *Config) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		context.Set(r, ConfigContextKey, config)
		next.ServeHTTP(w, r)
	})
}

// GetGameDefinition returns the GameDefinition struct for the given game ID,
// possibly falling back to a default definition.
// If the second return argument is false, no definition was found (the game
// ID is unsupported.)
func (config *Config) GetGameDefinition(gid shared.GameIDType) (shared.GameDefinition, bool) {
	def, ok := config.gameDefsById[gid]
	if !ok {
		if config.AllowUndefinedGames {
			return config.defaultGameDefinition(), true
		}
		ret := shared.GameDefinition{}
		return ret, false
	}
	return def, true
}

// LoadConfig takes a path and returns a master config instance.
func LoadConfig(configPath string) Config {
	var config Config
	shared.LoadConfigFromPath(configPath, &config)

	// Fill in denormalized/private fields.
	config.gameDefsById = make(map[shared.GameIDType]shared.GameDefinition)
	for _, def := range config.GameDefs {
		log.Printf("Supported game: %v", def.GameID)
		config.gameDefsById[def.GameID] = def
	}
	return config
}

func (config *Config) defaultGameDefinition() shared.GameDefinition {
	return shared.GameDefinition{
		GameID:     "_DEFAULT",
		HostPolicy: shared.HostPolicy_ANY,
	}
}
