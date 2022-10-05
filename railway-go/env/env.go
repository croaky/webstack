// Package env provides converts environment variables into Go data.
package env

import (
	"os"
)

// String returns the value of the named environment variable.
// If name isn't in the environment or is empty, it returns fallback.
func String(name string, fallback string) string {
	value, ok := os.LookupEnv(name)
	if !ok {
		value = fallback
	}
	return value
}
