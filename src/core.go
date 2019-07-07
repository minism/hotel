package main

// TODO: Use something like redis instead for this.
var _db map[ServerIDType]GameServer
var _nextID int = 0

func initCore() {
	_db = make(map[ServerIDType]GameServer)
}

func nextID() ServerIDType {
	_nextID++
	return ServerIDType(_nextID)
}
