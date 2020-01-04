package spawner

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func LaunchGameServer(config *Config, port uint32) (*os.Process, error) {
	cmd := exec.Command(config.GameServerPath, gameServerFlags(config, port)...)
	log.Printf("Launching game server with command: %v", cmd)
	err := cmd.Start()
	if err != nil {
		return nil, err
	}
	return cmd.Process, nil
}

func gameServerFlags(config *Config, port uint32) []string {
	ret := []string{
		fmt.Sprintf("--port=%v", port),
	}
	return ret
}
