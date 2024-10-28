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

func (j *Joneva) SliceBool(key string) ([]bool, error) {
	valueInterface := j.Get(key)
	if valueInterface == nil {
		return nil, fmt.Errorf("%v: %w", key, ErrNotFound)
	}

	valueSlice, ok := j.Get(key).([]interface{})
	if !ok {
		return nil, fmt.Errorf("%v: %w", key, ErrTypeMismatch)
	}

	valueSliceBool := make([]bool, len(valueSlice))
	for key, value := range valueSlice {
		if valueSliceBool[key], ok = value.(bool); !ok {
			return nil, fmt.Errorf("%v: %w", key, ErrTypeMismatch)
		}
	}

	return valueSliceBool, nil
}

func (j *Joneva) MustSliceBool(key string) []bool {
	valueSliceBool, err := j.SliceBool(key)
	if err != nil {
		panic(err)
	}

	return valueSliceBool
}

func (j *Joneva) SliceFloat64(key string) ([]float64, error) {
	valueInterface := j.Get(key)
	if valueInterface == nil {
		return nil, fmt.Errorf("%v: %w", key, ErrNotFound)
	}

	valueSlice, ok := j.Get(key).([]interface{})
	if !ok {
		return nil, fmt.Errorf("%v: %w", key, ErrTypeMismatch)
	}

	valueSliceFloat64 := make([]float64, len(valueSlice))
	for key, value := range valueSlice {
		if valueSliceFloat64[key], ok = value.(float64); !ok {
			return nil, fmt.Errorf("%v: %w", key, ErrTypeMismatch)
		}
	}

	return valueSliceFloat64, nil
}

func (j *Joneva) MustSliceFloat64(key string) []float64 {
	valueSliceFloat64, err := j.SliceFloat64(key)
	if err != nil {
		panic(err)
	}

	return valueSliceFloat64
}

func (j *Joneva) SliceString(key string) ([]string, error) {
	valueInterface := j.Get(key)
	if valueInterface == nil {
		return nil, fmt.Errorf("%v: %w", key, ErrNotFound)
	}

	valueSlice, ok := j.Get(key).([]interface{})
	if !ok {
		return nil, fmt.Errorf("%v: %w", key, ErrTypeMismatch)
	}

	valueSliceString := make([]string, len(valueSlice))
	for key, value := range valueSlice {
		if valueSliceString[key], ok = value.(string); !ok {
			return nil, fmt.Errorf("%v: %w", key, ErrTypeMismatch)
		}
	}

	return valueSliceString, nil
}

func (j *Joneva) MustSliceString(key string) []string {
	valueSliceString, err := j.SliceString(key)
	if err != nil {
		panic(err)
	}

	return valueSliceString
}
