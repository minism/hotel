package shared

import "os"

func GetEnv(key string, def string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return def
}
