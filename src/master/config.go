package master

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/context"
	"minornine.com/hotel/src/shared"
)

const (
	ConfigContextKey = 1
)

type Config struct {
	ReaperInterval       SerializableDuration    `json:"reaperInterval"`
	SessionExpiration    SerializableDuration    `json:"sessionExpiration"`
	ServerExpiration     SerializableDuration    `json:"serverExpiration"`
	SpawnerCheckInterval SerializableDuration    `json:"spawnerCheckInterval"`
	GameDefs             []shared.GameDefinition `json:"gameDefs"`
	AllowUndefinedGames  bool                    `json:"allowUndefinedGames"`

	gameDefsById map[shared.GameIDType]shared.GameDefinition
}

type SerializableDuration struct {
	time.Duration
}

func (sd *SerializableDuration) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	duration, err := time.ParseDuration(str)
	if err != nil {
		return err
	}
	sd.Duration = duration
	return nil
}

func (config *Config) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		context.Set(r, ConfigContextKey, config)
		next.ServeHTTP(w, r)
	})
}

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

func (config *Config) defaultGameDefinition() shared.GameDefinition {
	return shared.GameDefinition{
		GameID:     "_DEFAULT",
		HostPolicy: shared.HostPolicy_ANY,
	}
}

func LoadConfig(configPath string) Config {
	c := loadFromPath(configPath)

	// Fill in denormalized/private fields.
	c.gameDefsById = make(map[shared.GameIDType]shared.GameDefinition)
	for _, def := range c.GameDefs {
		log.Printf("Supported game: %v", def.GameID)
		c.gameDefsById[def.GameID] = def
	}
	return c
}

func loadFromPath(path string) Config {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Panicln("Could not load config:", err)
	}
	var config Config
	err = json.Unmarshal([]byte(data), &config)
	if err != nil {
		log.Panicln("Could not load config:", err)
	}
	log.Printf("Loaded config file %v\n", path)
	return config
}
