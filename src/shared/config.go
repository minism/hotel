package shared

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"time"
)

// SerializableDuration is a Duration type which can be serialized to/from JSON.
type SerializableDuration struct {
	time.Duration
}

// UnmarshalJSON fills the instance from JSON data.
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

// LoadConfigFromPath loads an arbitrary object from JSON data specified by the given path.
func LoadConfigFromPath(path string, config interface{}) interface{} {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Panicln("Could not load config:", err)
	}
	err = json.Unmarshal([]byte(data), &config)
	if err != nil {
		log.Panicln("Could not load config:", err)
	}
	log.Printf("Loaded config file %v\n", path)
	return config
}
