package main

// Unique ID for a game type, so hotel can multiplex across games.
type GameIDType string

// Unique ID for a game server instance.
type ServerIDType string

// An connectable game server instance.
type GameServer struct {
	GameID     GameIDType
	ServerID   ServerIDType
	Name       string
	Host       string
	Port       string
	NumPlayers int
	MaxPlayers int
}
