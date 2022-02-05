package joneva

import "fmt"

func (j *Joneva) Bool(key string) (bool, error) {
	valueInterface := j.Get(key)
	if valueInterface == nil {
		return false, fmt.Errorf("%v: %w", key, ErrNotFound)
	}

	valueBool, ok := j.Get(key).(bool)
	if !ok {
		return false, fmt.Errorf("%v: %w", key, ErrTypeMismatch)
	}

	return valueBool, nil
}

func (j *Joneva) MustBool(key string) bool {
	valueBool, err := j.Bool(key)
	if err != nil {
		panic(err)
	}

	return valueBool
}
