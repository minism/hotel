package shared

import (
	"encoding/json"
	"fmt"
)

// GameIDType is a unique ID for a game type, so hotel can multiplex across games.
type GameIDType string

// GameDefinition defines a specific game supported by a hotel server cluster,
// identified by a game ID, along with policies on how the game is managed.
type GameDefinition struct {
	GameID     GameIDType `json:"gameId"`
	HostPolicy HostPolicy `json:"hostPolicy"`
}

// HostPolicy determines how the game is allowed to be hosted/started by clients.
type HostPolicy int

const (
	// HostPolicy_DISABLED means new servers may not be hosted or spawned by regular clients.
	// They can only be hosted by privileged clients (server owners).
	HostPolicy_DISABLED HostPolicy = iota

	// HostPolicy_CLIENTS_ONLY means new servers can be hosted and registered by any client,
	// but clients cannot request new instances to be spawned.
	HostPolicy_CLIENTS_ONLY

	// HostPolicy_SPAWN_ONLY means new servers can be spawn requested, but arbitrary servers
	// cannot be hosted/registered.
	HostPolicy_SPAWN_ONLY

	// HostPolicy_ANY means new servers can be hosted by any client as well as requested to be spawned.
	HostPolicy_ANY
)

// UnmarshalJSON fills the instance from JSON data.
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
		return fmt.Errorf("Invalid host policy value: %v", str)
	}
	return nil
}
