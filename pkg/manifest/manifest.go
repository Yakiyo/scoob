package manifest

// A struct representing a scoob manifest
type Manifest struct {
	Version        string              `json:"version"`
	Url            string              `json:"url"`
	Description    string              `json:"description"`
	License        License             `json:"License"`
	Architecture   map[string]ArchInfo `json:"architecture"`
	Env_add_path   TOrSliceT[string]   `json:"env_add_path"`
	Bin            Bin                 `json:"bin"`
	Hash           TOrSliceT[string]   `json:"hash"`
	Extract_dir    string              `json:"extract_dir"`
	Extract_to     string              `json:"extract_to"`
	Innosetup      bool                `json:"innosetup"`
	Notes          TOrSliceT[string]   `json:"notes"`
	Persist        TOrSliceT[string]   `json:"persist"`
	Post_install   TOrSliceT[string]   `json:"post_install"`
	Post_uninstall TOrSliceT[string]   `json:"post_uninstall"`
	Pre_install    TOrSliceT[string]   `json:"pre_install"`
	Pre_uninstall  TOrSliceT[string]   `json:"pre_uninstall"`
}

// info specific to an architecture
type ArchInfo struct {
	Url         string            `json:"url"`
	Hash        TOrSliceT[string] `json:"hash"`
	Extract_dir string            `json:"extract_dir"`
	Extract_to  string            `json:"extract_to"`
}
