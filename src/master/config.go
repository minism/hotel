package master

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"time"

	"minornine.com/hotel/src/shared"
)

type Config struct {
	ReaperInterval      SerializableDuration    `json:"reaperInterval"`
	SessionExpiration   SerializableDuration    `json:"sessionExpiration"`
	ServerExpiration    SerializableDuration    `json:"serverExpiration"`
	AllowUndefinedGames bool                    `json:"allowUndefinedGames"`
	GameDefs            []shared.GameDefinition `json:"gameDefs"`
}

type SerializableDuration struct {
	time.Duration
}

func LoadConfig(configPath string) Config {
	c := loadFromPath(configPath)
	for _, def := range c.GameDefs {
		log.Printf("Supported game: %v", def.GameID)
	}
	return c
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
