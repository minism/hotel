package shared

import (
	"log"
	"os"
	"os/signal"
)

// WaitForSigIntAndQuit blocks until a SIGINT is received and then exits the program.
func WaitForSigIntAndQuit() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	log.Println("Shutting down.")
	os.Exit(0)
}
