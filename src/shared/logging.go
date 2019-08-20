package shared

import "log"

/// Configure global logging parameters.
func InitLogging() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
