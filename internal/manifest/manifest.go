package manifest

// A struct representing a scoob manifest
type Manifest struct {
	Version     string `json:"version"`
	Url         string `json:"url"`
	Description string `json:"description"`
	License     string `json:"License"`
}

// a license struct
// If license was only a string, we populate the identifier field, else if
// both were give, we use
type License struct {
	Identifier string `json:"identifier"`
	Url        string `json:"url"`
}

// info specific to an architecture
type ArchitectureInfo struct {
	Url         string `json:"url"`
	Hash        string `json:"hash"`
	Extract_dir string `json:"extract_dir"`
	Extract_to  string `json:"extract_to"`
}
