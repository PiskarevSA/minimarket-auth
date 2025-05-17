package jwtauth

import "errors"

var (
	ErrMissingJwt               = errors.New("missing jwt")
	ErrInvalidJwt               = errors.New("invalid jwt")
	ErrJwtNotYetValid           = errors.New("jwt not yet valid")
	ErrJwtExpired               = errors.New("jwt expired")
	ErrFailedToParseJwt         = errors.New("failed to parse jwt")
	ErrUnsupportedSigningMethod = errors.New("unsupported signing method")
)
