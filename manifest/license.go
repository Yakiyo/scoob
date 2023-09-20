package manifest

import (
	"errors"

	json "github.com/json-iterator/go"
)

// a license struct
// If license was only a string, we populate the identifier field, else if
// both were give, we use
type License struct {
	Identifier string `json:"identifier"`
	Url        string `json:"url"`
}

func (l *License) UnmarshalJSON(data []byte) error {
	var v interface{}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	switch v := v.(type) {
	case string:
		*l = License{Identifier: v}
	case map[string]interface{}:
		err := json.Unmarshal(data, l)
		if err != nil {
			return err
		}
	default:
		return errors.New("invalid value for license, must be string or map")
	}
	return nil
}
