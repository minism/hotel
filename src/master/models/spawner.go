package models

import (
	"fmt"

	"github.com/minism/hotel/src/shared"
)

// Spawner represents a running hotel spawner instance which we can RPC to.
type Spawner struct {
	ID             int
	Host           string
	Port           uint32
	GameID         shared.GameIDType
	NumGameServers uint32
	MaxGameServers uint32
}

// Capacity returns the number of available game servers this spawner could spawn.
func (s *Spawner) Capacity() uint32 {
	return s.MaxGameServers - s.NumGameServers
}

// Address returns a fully qualified host:port string for the spawner.
func (s *Spawner) Address() string {
	return fmt.Sprintf("%v:%v", s.Host, s.Port)
}
