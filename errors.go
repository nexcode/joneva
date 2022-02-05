package joneva

import "errors"

var (
	ErrNotFound     = errors.New("not found")
	ErrTypeMismatch = errors.New("type mismatch")
)
