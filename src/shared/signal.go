package shared

import (
	"os/signal"
	"os"
)

func CreateSigintChannel() chan os.Signal {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	return c
}
