package manifest

// A struct representing a scoob manifest
type Manifest struct {
	Version        string             `json:"version"`
	Url            string             `json:"url"`
	Description    string             `json:"description"`
	Extract_dir    string             `json:"extract_dir"`
	Extract_to     string             `json:"extract_to"`
	Env_add_path   string             `json:"env_add_path"`
	Innosetup      bool               `json:"innosetup"`
	Env_set        map[string]string  `json:"env_set"`
	Architecture   map[string]Arch    `json:"architecture"`
	License        License            `json:"License"`
	Bin            Bin                `json:"bin"`
	Hash           Vectorized[string] `json:"hash"`
	Notes          Vectorized[string] `json:"notes"`
	Persist        Vectorized[string] `json:"persist"`
	Post_install   Vectorized[string] `json:"post_install"`
	Post_uninstall Vectorized[string] `json:"post_uninstall"`
	Pre_install    Vectorized[string] `json:"pre_install"`
	Pre_uninstall  Vectorized[string] `json:"pre_uninstall"`
}
