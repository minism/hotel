package shared

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Unique ID for a game type, so hotel can multiplex across games.
type GameIDType string

// A GameDefinition defines a specific game supported by a hotel server cluster,
// identified by a game ID, along with policies on how the game is managed.
type GameDefinition struct {
	GameID     GameIDType `json:"gameId"`
	HostPolicy HostPolicy `json:"hostPolicy"`
}

// HostPolicy determines how the game is allowed to be hosted/started by clients.
type HostPolicy int

const (
	// New servers may not be hosted or spawned by regular clients. They can only be
	// hosted by priviledged clients (server owners).
	HostPolicy_DISABLED HostPolicy = iota

	// New servers can be hosted and registered by any client, but clients cannot request
	// new instances to be spawned.
	HostPolicy_CLIENTS_ONLY

	// New servers can be spwan requested, but arbitrary servers cannot be hosted/registered.
	HostPolicy_SPAWN_ONLY

	// New servers can be hosted by any client as well as requested to be spawned.
	HostPolicy_ANY
)

func (hp *HostPolicy) UnmarshalJSON(data []byte) error {
	var str string
	err := json.Unmarshal(data, &str)
	if err != nil {
		return err
	}
	if str == "DISABLED" {
		*hp = HostPolicy_DISABLED
	} else if str == "CLIENTS_ONLY" {
		*hp = HostPolicy_CLIENTS_ONLY
	} else if str == "SPAWN_ONLY" {
		*hp = HostPolicy_SPAWN_ONLY
	} else if str == "ANY" {
		*hp = HostPolicy_ANY
	} else {
		return errors.New(fmt.Sprintf("Invalid host policy value: %v", str))
	}
	return nil
}
