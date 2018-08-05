package envy

import (
	"errors"
	"os"
)

// Get returns a string from the requested environment var or panics.
func Get(name string) (string, error) {
	v := os.Getenv(name)
	if v == "" {
		return "", errors.New("missing required environment variable " + name)
	}
	return v, nil
}
