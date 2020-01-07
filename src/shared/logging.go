package shared

import "log"

// InitLogging configures global logging parameters.
func InitLogging() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
