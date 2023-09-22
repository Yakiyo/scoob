package app

import (
	"strings"

	"github.com/Yakiyo/scoob/manifest"
)

func get_arch(m manifest.Manifest, arch string) (manifest.Arch, error) {
	archs := m.Architecture
	r, ok := archs[arch]
	if !ok {
		r = manifest.Arch{
			Bin:         m.Bin,
			Installer:   m.Installer,
			Uninstaller: m.Uninstaller,
			Hash:        m.Hash,
			Url:         m.Url,
			Extract_dir: m.Extract_dir,
			Extract_to:  m.Extract_to,
			Shortcuts:   m.Shortcuts,
		}
	}
	return r, nil
}

// get file name from a url, and return the (url, fname)
func url_fname(url string) (string, string) {
	split := strings.Split(url, "#/")
	if len(split) == 2 {
		return split[0], split[1]
	}
	url = split[0]
	split = strings.Split(url, "/")
	fname := split[len(split)-1]
	return url, fname
}
