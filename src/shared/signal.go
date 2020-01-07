package shared

import (
	"os"
	"os/signal"
)

// CreateSigintChannel returns a channel which listens to SIGINT.
func CreateSigintChannel() chan os.Signal {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	return c
}
