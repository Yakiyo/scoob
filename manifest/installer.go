package manifest

// The `installer` used in Installer key and Uninstaller
type Installer struct {
	File   string             `json:"file"` // for installer, defaults to url, required for uninstaller
	Script Vectorized[string] `json:"script"`
	Args   Vectorized[string] `json:"args"`
	Keep   string             `json:"keep"` // "true" (string, not bool) to keep file, else remove. Ignored for uninstaller
}
