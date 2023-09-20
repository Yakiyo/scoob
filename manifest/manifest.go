package manifest

import (
	"os"

	json "github.com/json-iterator/go"
)

// A struct representing a scoob manifest
type Manifest struct {
	Version        string              `json:"version"`
	Homepage       string              `json:"homepage"`
	Description    string              `json:"description"`
	Extract_dir    string              `json:"extract_dir"`
	Extract_to     string              `json:"extract_to"`
	Env_add_path   string              `json:"env_add_path"`
	Innosetup      bool                `json:"innosetup"`
	Shortcuts      [][]string          `json:"shortcuts"`
	Suggest        map[string][]string `json:"suggest"`
	Env_set        map[string]string   `json:"env_set"`
	Architecture   map[string]Arch     `json:"architecture"`
	Url            Vectorized[string]  `json:"url"`
	Hash           Vectorized[string]  `json:"hash"`
	Notes          Vectorized[string]  `json:"notes"`
	Persist        Vectorized[string]  `json:"persist"`
	Post_install   Vectorized[string]  `json:"post_install"`
	Post_uninstall Vectorized[string]  `json:"post_uninstall"`
	Pre_install    Vectorized[string]  `json:"pre_install"`
	Pre_uninstall  Vectorized[string]  `json:"pre_uninstall"`
	License        License             `json:"License"`
	Bin            Bin                 `json:"bin"`
	Installer      Installer           `json:"installer"`
	Uninstaller    Installer           `json:"uninstaller"`
}

// parse a manifest file
func Parse(file string) (Manifest, error) {
	m := Manifest{}
	b, err := os.ReadFile(file)
	if err != nil {
		return m, err
	}
	err = json.Unmarshal(b, &m)
	return m, err
}
