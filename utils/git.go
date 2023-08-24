package utils

import (
	"errors"
	"os/exec"
)

// Wether git is in path or not
func GitExists() bool {
	return exec.Command("git").Path != ""
}

// run git, suppress should be true when we dont want output to show
func RunGit(suppress bool, args ...string) (int, error) {
	c := exec.Command("git", args...)
	if suppress {
		c.Stdout = nil
		c.Stderr = nil
		c.Stdin = nil
	}
	if err := c.Run(); err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			return exiterr.ProcessState.ExitCode(), errors.New("failed to run git successfully, got non-zero exit")
		}
		return 1, err
	}
	return 0, nil
}
