package joneva

import "fmt"

func (j *Joneva) Map(key string) (map[string]interface{}, error) {
	valueInterface := j.Get(key)
	if valueInterface == nil {
		return nil, fmt.Errorf("%v: %w", key, ErrNotFound)
	}

	valueMap, ok := j.Get(key).(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("%v: %w", key, ErrTypeMismatch)
	}

	return valueMap, nil
}

func (j *Joneva) MustMap(key string) map[string]interface{} {
	valueMap, err := j.Map(key)
	if err != nil {
		panic(err)
	}

	return valueMap
}
