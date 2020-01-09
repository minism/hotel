package spawner

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
)

// LaunchGameServer execs the game server subprocess on the given port.
// The Process pointer representing the new process is returned.
func LaunchGameServer(config *Config, port uint32) (*os.Process, error) {
	cmd := exec.Command(config.GameServerPath, gameServerFlags(config, port)...)
	log.Printf("Launching game server with command: %v", cmd)

	// Setup stdout/stderr logging.
	// TODO: Configure logging to file instead.
	pipe, _ := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	scanner := bufio.NewScanner(pipe)
	go func() {
		for scanner.Scan() {
			log.Printf("[game-server %v] %v", port, scanner.Text())
		}
	}()

	// Start the process async.
	if err := cmd.Start(); err != nil {
		return nil, err
	}
	return cmd.Process, nil
}

func gameServerFlags(config *Config, port uint32) []string {
	flags := []string{
		fmt.Sprintf("--port=%v", port),
	}
	if config.Host != "" {
		flags = append(flags, fmt.Sprintf("--host=%v", config.Host))
	}
	return append(config.GameServerFlags, flags...)
}
