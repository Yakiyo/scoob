package utils

import (
	"os"

	"github.com/charmbracelet/log"
)

// use log.error and exit
func Error(msg interface{}, args ...interface{}) {
	log.Error(msg, args...)
	os.Exit(1)
}
