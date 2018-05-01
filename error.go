package cmp

import (
	"errors"
)

var (
	ErrNotSameKind      = errors.New("not same kind")
	ErrKindNotSupported = errors.New("kind not supported")
)
