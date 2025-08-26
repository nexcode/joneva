package joneva

import (
	"fmt"
	"time"
)

// Duration gets the key value as time.Duration.
// Supports standard Go formats: "300ms", "1.5h", "2h45m", "30s", "5m", etc.
// Also supports numbers like nanoseconds.
func (j *Joneva) Duration(key string) (time.Duration, error) {
	valueInterface := j.Get(key)
	if valueInterface == nil {
		return 0, fmt.Errorf("%v: %w", key, ErrNotFound)
	}

	switch v := valueInterface.(type) {
	case string:
		return time.ParseDuration(v)
	case float64:
		return time.Duration(v), nil
	default:
		return 0, fmt.Errorf("%v: %w", key, ErrTypeMismatch)
	}
}

// MustDuration gets the key value as time.Duration or panics on error.
func (j *Joneva) MustDuration(key string) time.Duration {
	duration, err := j.Duration(key)
	if err != nil {
		panic(err)
	}
	return duration
}

// SetDuration sets the value of a key as a duration in string format.
func (j *Joneva) SetDuration(key string, value time.Duration) {
	j.Set(key, value.String())
}
