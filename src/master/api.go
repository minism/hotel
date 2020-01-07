package master

// IdentifyResponse is returned from the identify API call.
type IdentifyResponse struct {
	Token string `json:"token"`
}

// ListServersResponse is returned from the list servers API call.
type ListServersResponse struct {
	Servers []GameServer `json:"servers"`
}
