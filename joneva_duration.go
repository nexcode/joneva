package joneva

import (
	"fmt"
	"time"
)

// Duration получает значение ключа как time.Duration.
// Поддерживает стандартные форматы Go: "300ms", "1.5h", "2h45m", "30s", "5m" и т.д.
// Также поддерживает числа как наносекунды.
func (j *Joneva) Duration(key string) (time.Duration, error) {
	valueInterface := j.Get(key)
	if valueInterface == nil {
		return 0, fmt.Errorf("%v: %w", key, ErrNotFound)
	}

	switch v := valueInterface.(type) {
	case string:
		return time.ParseDuration(v)
	case float64:
		// Числа интерпретируются как наносекунды
		return time.Duration(v), nil
	default:
		return 0, fmt.Errorf("%v: %w", key, ErrTypeMismatch)
	}
}

// MustDuration получает значение ключа как time.Duration или паникует при ошибке.
func (j *Joneva) MustDuration(key string) time.Duration {
	duration, err := j.Duration(key)
	if err != nil {
		panic(err)
	}
	return duration
}

// SetDuration устанавливает значение ключа как время в виде строки.
func (j *Joneva) SetDuration(key string, value time.Duration) {
	j.Set(key, value.String())
}
