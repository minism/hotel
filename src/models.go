package main

// Unique string for a game type, so hotel can multiplex across games.
type GameIDType string

// An connectable game server instance.
type GameServer struct {
	GameID     GameIDType
	Name       string
	Host       string
	Port       string
	NumPlayers int
	MaxPlayers int
}
