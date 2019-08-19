package master

type IdentifyResponse struct {
	Token string `json:"token"`
}

type ListServersResponse struct {
	Servers []GameServer `json:"servers"`
}
