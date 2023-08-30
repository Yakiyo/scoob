package manifest

import (
	"errors"
	json "github.com/json-iterator/go"
	"github.com/samber/lo"
)

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
