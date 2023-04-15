package middlewares

import "errors"

var (
	// 4xx
	ErrBadRequest   = errors.New("bad request")
	ErrNotFound     = errors.New("not found")
	ErrUnauthorized = errors.New("unauthorized")
	ErrConflict     = errors.New("conflict")

	// 5xx
	ErrRepository = errors.New("repository")
)
