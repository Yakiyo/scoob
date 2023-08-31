package manifest

// info about a single architecture
type Arch struct {
	Bin         Bin                `json:"bin"`
	Installer   Installer          `json:"installer"`
	Uninstaller Installer          `json:"uninstaller"`
	Hash        Vectorized[string] `json:"hash"`
	Url         Vectorized[string] `json:"url"`
	Extract_dir string             `json:"extract_dir"`
	Extract_to  string             `json:"extract_to"`
	Shortcuts   [][]string         `json:"shortcuts"`
}
