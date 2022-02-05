package joneva

import "fmt"

func (j *Joneva) Slice(key string) ([]interface{}, error) {
	valueInterface := j.Get(key)
	if valueInterface == nil {
		return nil, fmt.Errorf("%v: %w", key, ErrNotFound)
	}

	valueSlice, ok := j.Get(key).([]interface{})
	if !ok {
		return nil, fmt.Errorf("%v: %w", key, ErrTypeMismatch)
	}

	return valueSlice, nil
}

func (j *Joneva) MustSlice(key string) []interface{} {
	valueSlice, err := j.Slice(key)
	if err != nil {
		panic(err)
	}

	return valueSlice
}
