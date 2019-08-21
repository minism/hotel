package master

import (
	"encoding/json"
	"errors"
	"io"
)

// Unique ID for a game type, so hotel can multiplex across games.
type GameIDType string

// Unique ID for a game server instance.
type ServerIDType int

// A connectable game server instance.
type GameServer struct {
	ID         ServerIDType `json:"id"`
	GameID     GameIDType   `json:"gameId"`
	SessionID  int          `json:"-"`
	Name       string       `json:"name"`
	Host       string       `json:"host"`
	Port       int          `json:"port"`
	NumPlayers int          `json:"numPlayers"`
	MaxPlayers int          `json:"maxPlayers"`
}

// Decode a JSON version of GameServer, validate it, and return it.
// If any steps fail, returns an error.
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

func (s *GameServer) Validate(isUpdate bool) error {
	if !isUpdate && len(s.Name) < 1 {
		return errors.New("Name must be non-empty.")
	}
	return nil
}

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
	if other.NumPlayers > 0 {
		s.NumPlayers = other.NumPlayers
	}
	if other.MaxPlayers > 0 {
		s.MaxPlayers = other.MaxPlayers
	}
}
