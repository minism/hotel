package main

// Unique ID for a game type, so hotel can multiplex across games.
type GameIDType string

// Unique ID for a game server instance.
type ServerIDType string

// An connectable game server instance.
type GameServer struct {
	ID         ServerIDType `json:"id"`
	GameID     GameIDType   `json:"gameId"`
	Name       string       `json:"name"`
	Host       string       `json:"host"`
	Port       string       `json:"port"`
	NumPlayers int          `json:"numPlayers"`
	MaxPlayers int          `json:"maxPlayers"`
}

func (s *GameServer) UpdateFrom(other *GameServer) error {
	return nil
}
