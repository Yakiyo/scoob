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

// create dir if it does not exist
func EnsureDir(path string) error {
	if PathExists(path) {
		return nil
	}
	return os.MkdirAll(path, os.ModePerm)
}

// create file if it does not exist
func EnsureFile(path string) error {
	if PathExists(path) {
		return nil
	}
	return os.WriteFile(path, []byte{}, os.ModePerm)
}
