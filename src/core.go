package main

func getServersByGameId(gid GameIDType) []GameServer {
	ret := make([]GameServer, 0)
	ret = append(ret, GameServer{Name: "Test"})
	return ret
}

func getServerById(id ServerIDType) GameServer {
	return GameServer{Name: "Test"}
}

func createServer(server GameServer) {

}

func updateServerById(id ServerIDType, server GameServer) {

}
