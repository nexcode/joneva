package joneva

import "fmt"

func (j *Joneva) String(key string) (string, error) {
	valueInterface := j.Get(key)
	if valueInterface == nil {
		return "", fmt.Errorf("%v: %w", key, ErrNotFound)
	}

	valueString, ok := valueInterface.(string)
	if !ok {
		return "", fmt.Errorf("%v: %w", key, ErrTypeMismatch)
	}

	return valueString, nil
}

func (j *Joneva) MustString(key string) string {
	valueString, err := j.String(key)
	if err != nil {
		panic(err)
	}

	return valueString
}
