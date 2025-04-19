package cond

import "errors"

var (
	ErrInvalidAMQPConfig error = errors.New("conductor err : invalid amqp config")
	ErrInvalidParams     error = errors.New("conductor err : invalid params")

	ErrInvalidCode      error = errors.New("conductor err : invalid code")
	ErrCodeRefreshBlock error = errors.New("conductor err : code refresh block")
)
