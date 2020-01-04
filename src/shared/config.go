package shared

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"time"
)

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
