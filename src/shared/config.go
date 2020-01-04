package shared

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

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
