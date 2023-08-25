package utils

import (
	"errors"
	"os"
)

// check if a file exists
func PathExists(path string) bool {
	_, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		return false
	} else {
		return true
	}
}
