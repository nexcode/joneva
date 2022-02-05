package joneva

import "fmt"

func (j *Joneva) Float64(key string) (float64, error) {
	valueInterface := j.Get(key)
	if valueInterface == nil {
		return 0, fmt.Errorf("%v: %w", key, ErrNotFound)
	}

	valueFloat64, ok := j.Get(key).(float64)
	if !ok {
		return 0, fmt.Errorf("%v: %w", key, ErrTypeMismatch)
	}

	return valueFloat64, nil
}

func (j *Joneva) MustFloat64(key string) float64 {
	valueFloat64, err := j.Float64(key)
	if err != nil {
		panic(err)
	}

	return valueFloat64
}
