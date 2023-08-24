package utils

import "testing"

// this test should generally pass since most people and CI's do have git in em
func TestGitExists(t *testing.T) {
	if !GitExists() {
		t.Fatalf("Couldnt find git")
	}
}

func RunGitTest(t *testing.T) {
	RunGit(true, "-v")
}
