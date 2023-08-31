package utils

import (
	"errors"
	"runtime"
	"strings"

	"github.com/spf13/viper"
)

// get default architecture
func GetDefaultArch() (string, error) {
	arch := runtime.GOARCH
	if c := viper.GetString("arch"); c != "" {
		arch = c
	}
	return FormatArch(arch)
}

// format architecture string to app friendly one
func FormatArch(arch string) (string, error) {
	switch strings.ToLower(arch) {
	case "64bit", "64", "x64", "amd64", "x86_64", "x86-64":
		return "64bit", nil
	case "32bit", "32", "x86", "i386", "386", "i686":
		return "32bit", nil
	case "arm64", "arm", "aarch64":
		return "arm64", nil
	default:
		return "", errors.New("invalid architecture " + arch)
	}
}
