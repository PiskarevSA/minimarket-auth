package jwtauth

import (
	"context"

	"github.com/golang-jwt/jwt/v5"
)

type (
	ParseFn     func(context.Context, *JwtAuth) jwt.Keyfunc
	ValidatorFn func(context.Context, *jwt.Token) error
)

type JwtAuth struct {
	SigningMethod jwt.SigningMethod
	SigningKey    any
	VerifyKey     any
	Claims        func() jwt.Claims
	ParseFn       ParseFn
	ValidatorFns  []ValidatorFn
}

func New(signingKey any, opts ...Option) *JwtAuth {
	ja := &JwtAuth{
		SigningMethod: DefaultSigningMethod,
		SigningKey:    signingKey,
		ParseFn:       DefaultParseFn,
		ValidatorFns:  make([]ValidatorFn, 0),
	}

	for _, opt := range opts {
		opt(ja)
	}

	return ja
}
