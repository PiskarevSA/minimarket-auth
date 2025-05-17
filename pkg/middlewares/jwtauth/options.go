package jwtauth

import "github.com/golang-jwt/jwt/v5"

type Option func(*JwtAuth)

func WitSigningMethod(method jwt.SigningMethod) Option {
	return func(a *JwtAuth) {
		a.SigningMethod = method
	}
}

func WithSigningKey(signingKey any) Option {
	return func(a *JwtAuth) {
		a.SigningKey = signingKey
	}
}

func WithVerifyKey(key any) Option {
	return func(a *JwtAuth) {
		a.VerifyKey = key
	}
}

func WithClaims(f func() jwt.Claims) Option {
	return func(o *JwtAuth) {
		o.Claims = f
	}
}

func WithValidator(fn ValidatorFn) Option {
	return func(a *JwtAuth) {
		a.ValidatorFns = append(a.ValidatorFns, fn)
	}
}
