package main

func getServersByGameId(gid GameIDType) []GameServer {
	ret := make([]GameServer, 0)
	ret = append(ret, GameServer{Name: "Test"})
	return ret
}
