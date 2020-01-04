package spawner

import mapset "github.com/deckarep/golang-set"

// ServerController manages the lifecycle of game server processes.
type ServerController struct {
	config         *Config
	servers        []ServerProcess
	availablePorts mapset.Set
}

// RunningServer represents an actively running game server process.
type ServerProcess struct {
	Port uint32
	Pid  int
}

func NewServerController(config *Config) *ServerController {
	// Initialize a pool of available ports, staring at the spawners base port + 1.
	ports := mapset.NewSet()
	for i := uint32(0); i < config.MaxGameServers; i++ {
		ports.Add(config.Port + i + 1)
	}

	return &ServerController{
		config:         config,
		servers:        make([]ServerProcess, 0),
		availablePorts: ports,
	}
}

func (c *ServerController) NumRunningServers() int {
	return len(c.servers)
}

func (c *ServerController) Capacity() int {
	return c.availablePorts.Cardinality()
}

// Spawn a server process and return its port.
func (c *ServerController) SpawnServer() (uint32, error) {
	nextPort := c.availablePorts.Pop().(uint32)
	pid, err := LaunchGameServer(c.config, nextPort)
	if err != nil {
		c.availablePorts.Add(nextPort)
		return nextPort, err
	} else {
		c.servers = append(c.servers, ServerProcess{Port: nextPort, Pid: pid})
	}
	return nextPort, nil
}
