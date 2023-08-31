package manifest

// info about a single architecture
type Arch struct {
	Url         string             `json:"url"`
	Hash        Vectorized[string] `json:"hash"`
	Extract_dir string             `json:"extract_dir"`
	Extract_to  string             `json:"extract_to"`
}
