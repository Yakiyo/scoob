package manifest

import (
	"errors"

	json "github.com/json-iterator/go"
	"github.com/samber/lo"
)

// A struct representing a scoob manifest
type Manifest struct {
	Version      string              `json:"version"`
	Url          string              `json:"url"`
	Description  string              `json:"description"`
	License      License             `json:"License"`
	Architecture map[string]ArchInfo `json:"architecture"`
	Env_add_path []string            `json:"env_add_path"`
	Bin          Bin                 `json:"bin"`
	Hash         Hash                `json:"hash"`
	Extract_dir  string              `json:"extract_dir"`
	Extract_to   string              `json:"extract_to"`
}

// a license struct
// If license was only a string, we populate the identifier field, else if
// both were give, we use
type License struct {
	Identifier string `json:"identifier"`
	Url        string `json:"url"`
}

// info specific to an architecture
type ArchInfo struct {
	Url         string `json:"url"`
	Hash        Hash   `json:"hash"`
	Extract_dir string `json:"extract_dir"`
	Extract_to  string `json:"extract_to"`
}

// The bin key
// this can be one of the following:
//   - a string
//   - a slice of strings
//   - a slice of (slice of strings)
type Bin [][]string

func (b *Bin) UnmarshalJSON(data []byte) error {
	if len(data) < 1 {
		return errors.New("required key `bin` is empty, must be an array of string arrays or a string")
	}
	v := [][]string{}
	var i interface{}
	err := json.Unmarshal(data, &i)
	if err != nil {
		return err
	}
	switch d := i.(type) {
	case string:
		v = [][]string{{d}}
	case []interface{}:
		for _, inner := range d {
			switch inner := inner.(type) {
			case string:
				v = append(v, []string{inner})
			case []interface{}:
				v = append(v, lo.Map(inner, func(item interface{}, _ int) string { return item.(string) }))
			}
		}
	default:
		return errors.New("invalid value for bin received, must be one of string, []string or [][]string")
	}
	t := Bin(v)
	*b = t
	return nil
}

// can be a string or a slice of strings
type Hash []string
