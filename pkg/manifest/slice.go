package manifest

import (
	"errors"

	json "github.com/json-iterator/go"
	"github.com/samber/lo"
)

// A type that can be either `T` or a slice of `T`. This will be
// unmarshalled into a slice of `T`
type Vectorized[T any] []T

func (t *Vectorized[T]) UnmarshalJSON(data []byte) error {
	if len(data) < 1 {
		return nil
	}
	var i interface{}
	err := json.Unmarshal(data, &i)
	if err != nil {
		return err
	}
	switch i := i.(type) {
	case T:
		*t = []T{i}
	case []interface{}:
		*t = lo.Map(i, func(item interface{}, _ int) T {
			return item.(T)
		})
	default:
		return errors.New("received invalid type, neither of generic type T or slice T")
	}
	return nil
}
