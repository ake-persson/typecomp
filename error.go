package typecomp

import (
	"errors"
)

var (
	// ErrNotSameKind values are not of the same type.
	ErrNotSameKind = errors.New("not same kind")

	// ErrKindNotSupported unsuported type.
	ErrKindNotSupported = errors.New("kind not supported")
)
