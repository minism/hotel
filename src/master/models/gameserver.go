package models

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/minism/hotel/src/shared"
)

// ServerIDType is a type alias for game server IDs.
type ServerIDType int

// GameServer represents a connectable game server instance.
type GameServer struct {
	ID         ServerIDType      `json:"id"`
	GameID     shared.GameIDType `json:"gameId"`
	SessionID  int               `json:"-"`
	Name       string            `json:"name"`
	Host       string            `json:"host"`
	Port       int               `json:"port"`
	NumPlayers *int              `json:"numPlayers"`
	MaxPlayers *int              `json:"maxPlayers"`
}

// DecodeAndValidateGameServer decodes a JSON version of GameServer, validates it, and returns it.
// If any steps fail, an error is returned.
func DecodeAndValidateGameServer(reader io.Reader, isUpdate bool) (GameServer, error) {
	decoder := json.NewDecoder(reader)
	var server GameServer
	err := decoder.Decode(&server)
	if err != nil {
		return server, err
	}
	err = server.Validate(isUpdate)
	if err != nil {
		return server, err
	}
	return server, nil
}

// Validate checks if any GameServer fields are invalid and returns an error if so.
func (s *GameServer) Validate(isUpdate bool) error {
	if !isUpdate && len(s.Name) < 1 {
		return errors.New("Name must be non-empty.")
	}
	if !isUpdate && s.Port < 1 || s.Port > 65535 {
		return errors.New("Port not in range.")
	}
	return nil
}

// Merge will merge all fields from the given game server into this instance.
func (s *GameServer) Merge(other GameServer) {
	// TODO: This should use reflection to walk mutable fields and apply them.
	if len(other.Name) > 0 {
		s.Name = other.Name
	}
	if len(other.Host) > 0 {
		s.Host = other.Host
	}
	if other.Port > 0 {
		s.Port = other.Port
	}
	if other.NumPlayers != nil {
		s.NumPlayers = other.NumPlayers
	}
	if other.MaxPlayers != nil {
		s.MaxPlayers = other.MaxPlayers
	}
}
