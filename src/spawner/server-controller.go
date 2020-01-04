package spawner

import (
	"log"

	mapset "github.com/deckarep/golang-set"
)

// ServerController manages the lifecycle of game server processes.
type ServerController struct {
	config         *Config
	availablePorts mapset.Set
}

func NewServerController(config *Config) *ServerController {
	// Initialize a pool of available ports, staring at the spawners base port + 1.
	ports := mapset.NewSet()
	for i := uint32(0); i < config.MaxGameServers; i++ {
		ports.Add(config.Port + i + 1)
	}

	controller := &ServerController{
		config:         config,
		availablePorts: ports,
	}

	// If we're configured to autorun, do that now.
	if config.Autorun {
		for i := uint32(0); i < config.MaxGameServers; i++ {
			controller.SpawnServer()
		}
	}

	return controller
}

func (c *ServerController) Capacity() int {
	return c.availablePorts.Cardinality()
}

func (c *ServerController) NumRunningServers() int {
	return int(c.config.MaxGameServers) - c.Capacity()
}

// Spawn a server process and return its port.
func (c *ServerController) SpawnServer() (uint32, error) {
	port := c.availablePorts.Pop().(uint32)
	process, err := LaunchGameServer(c.config, port)
	if err != nil {
		c.availablePorts.Add(port)
		return port, err
	} else {
		go func() {
			// Make the port available when the server has terminated.
			state, err := process.Wait()
			if err != nil {
				log.Printf("Game server terminated with error: %v", err)
			}
			log.Printf("Game server pid %v terminated, returning port %v to pool.", process.Pid, port)
			log.Printf(state.String())
			c.availablePorts.Add(port)
		}()
	}
	return port, nil
}
