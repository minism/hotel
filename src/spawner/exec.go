package spawner

import (
	"fmt"
	"log"
	"os/exec"
)

func LaunchGameServer(config *Config, port uint32) (int, error) {
	cmd := exec.Command(config.GameServerPath, gameServerFlags(config, port)...)
	log.Printf("Launching game server with command: %v", cmd)
	err := cmd.Start()
	if err != nil {
		return 0, err
	}
	return cmd.Process.Pid, nil
}

func gameServerFlags(config *Config, port uint32) []string {
	ret := []string{
		fmt.Sprintf("--port=%v", port),
	}
	return ret
}
