package cond

import "errors"

var (
	ErrInvalidParams       error = errors.New("conductor err : invalid params")
	ErrInvalidTemplate     error = errors.New("conductor err : invalid template")
	ErrInvalidSenderParams error = errors.New("conductor err : invalid sender params")

	ErrInvalidCode error = errors.New("conductor err : invalid code")
)
