package main

import "errors"

// TODO: Use something like redis instead for this.
var _db map[ServerIDType]*GameServer

func initCore() {
	_db = make(map[ServerIDType]*GameServer)
}

func getServersByGameId(gid GameIDType) []GameServer {
	ret := make([]GameServer, 0)
	ret = append(ret, GameServer{Name: "Test"})
	return ret
}

func getServerById(id ServerIDType) *GameServer {
	return _db[id]
}

func insertServer(server GameServer) {
	_db[server.ID] = &server
}

func updateServerById(id ServerIDType, server *GameServer) error {
	curr, ok := _db[id]
	if !ok {
		return errors.New("No server with ID type: " + string(id))
	}
	err := curr.UpdateFrom(server)
	if err != nil {
		return err
	}
	return nil
}

func pingServerAlive(id ServerIDType) {

}
