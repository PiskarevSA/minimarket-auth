package signin

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Usecase struct {
	storage            Storage
	jwtSignKey         any
	jwtSigningMethod   jwt.SigningMethod
	accessTokenExpiry  time.Duration
	refreshTokenExpiry time.Duration
}

func New(
	storage Storage,
	signKey any,
	signingMethod jwt.SigningMethod,
	accessTokenExpiry time.Duration,
	refreshTokenExpiry time.Duration,
) *Usecase {
	return &Usecase{
		storage:            storage,
		jwtSignKey:         signKey,
		jwtSigningMethod:   signingMethod,
		accessTokenExpiry:  accessTokenExpiry,
		refreshTokenExpiry: refreshTokenExpiry,
	}
}
