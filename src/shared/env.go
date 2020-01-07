package shared

import "os"

// GetEnv returns an OS environment variable with a safety fallback.
func GetEnv(key string, def string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return def
}
