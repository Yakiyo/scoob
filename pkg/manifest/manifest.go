package manifest

// A struct representing a scoob manifest
type Manifest struct {
	Version        string              `json:"version"`
	Url            string              `json:"url"`
	Description    string              `json:"description"`
	License        License             `json:"License"`
	Architecture   map[string]ArchInfo `json:"architecture"`
	Env_add_path   Vectorized[string]  `json:"env_add_path"`
	Bin            Bin                 `json:"bin"`
	Hash           Vectorized[string]  `json:"hash"`
	Extract_dir    string              `json:"extract_dir"`
	Extract_to     string              `json:"extract_to"`
	Innosetup      bool                `json:"innosetup"`
	Notes          Vectorized[string]  `json:"notes"`
	Persist        Vectorized[string]  `json:"persist"`
	Post_install   Vectorized[string]  `json:"post_install"`
	Post_uninstall Vectorized[string]  `json:"post_uninstall"`
	Pre_install    Vectorized[string]  `json:"pre_install"`
	Pre_uninstall  Vectorized[string]  `json:"pre_uninstall"`
}

// info specific to an architecture
type ArchInfo struct {
	Url         string             `json:"url"`
	Hash        Vectorized[string] `json:"hash"`
	Extract_dir string             `json:"extract_dir"`
	Extract_to  string             `json:"extract_to"`
}
