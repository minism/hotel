package main

// Unique ID for a game type, so hotel can multiplex across games.
type GameId string

// An connectable game server instance.
type GameServer struct {
	Name string
}
