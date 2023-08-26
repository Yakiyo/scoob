package bucket

import (
	"io/fs"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/Yakiyo/scoob/utils"
)

// data recovered from a bucket directory
type BucketDir struct {
	Name    string
	Path    string
	Source  string
	LastMod time.Time
}

// create a `BucketDir` struct from a directory info
func fromDir(f fs.DirEntry, path string) BucketDir {
	b := BucketDir{
		Name:   strings.TrimSpace(f.Name()),
		Source: path,
	}
	s, err := f.Info()
	if err != nil {
		utils.Error("Unable to read info for bucket directory", "dir", f.Name(), "err", err)
	}
	b.LastMod = s.ModTime()
	// if bucket is a git dir and git is installed in system, find source url
	if utils.PathExists(filepath.Join(path, ".git")) && utils.GitExists() {

		url, err := exec.Command("git", "-C", path, "config", "remote.origin.url").Output()
		if err != nil {
			utils.Error("Unable to get remote url from git repo", "dir", path, "err", err)
		}
		b.Source = strings.TrimSpace(string(url))
	}
	return b
}
