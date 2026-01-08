package auth

import "errors"

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrTokenInvalid       = errors.New("token invalid")
	ErrUnauthorized       = errors.New("unauthorized")
)
