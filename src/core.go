package main

import "errors"

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

func getServersByGameId(gid GameIDType) []GameServer {
	ret := make([]GameServer, 0)
	ret = append(ret, GameServer{Name: "Test"})
	return ret
}

func getServerById(id ServerIDType) GameServer {
	return _db[id]
}

func insertServer(server GameServer) (GameServer, error) {
	server.ID = nextID()
	_db[server.ID] = server
	return server, nil
}

func updateServerById(id ServerIDType, server GameServer) (GameServer, error) {
	curr, ok := _db[id]
	if !ok {
		return curr, errors.New("No server with ID type: " + string(id))
	}
	err := curr.UpdateFrom(&server)
	if err != nil {
		return curr, err
	}
	return curr, nil
}

func pingServerAlive(id ServerIDType) {

}
