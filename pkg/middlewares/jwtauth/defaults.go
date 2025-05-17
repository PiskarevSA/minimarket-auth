package jwtauth

import (
	"context"

	"github.com/golang-jwt/jwt/v5"
)

var (
	DefaultSigningMethod = jwt.SigningMethodHS256
	DefaultParseFn       = func(ctx context.Context, ja *JwtAuth) jwt.Keyfunc {
		return func(t *jwt.Token) (any, error) {
			return ja.VerifyKey, nil
		}
	}
)
