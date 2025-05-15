package signup

import (
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/PiskarevSA/minimarket-points/pkg/pgx/transactor"
)

type Usecase struct {
	serviceName        string
	storage            Storage
	transactor         *transactor.Transactor
	jwtSignKey         any
	jwtSigningMethod   jwt.SigningMethod
	accessTokenExpiry  time.Duration
	refreshTokenExpiry time.Duration
}

func New(
	serviceName string,
	storage Storage,
	transactor *transactor.Transactor,
	signKey any,
	signingMethod jwt.SigningMethod,
	accessTokenExpiry time.Duration,
	refreshTokenExpiry time.Duration,
) *Usecase {
	return &Usecase{
		serviceName:        serviceName,
		storage:            storage,
		transactor:         transactor,
		jwtSignKey:         signKey,
		jwtSigningMethod:   signingMethod,
		accessTokenExpiry:  accessTokenExpiry,
		refreshTokenExpiry: refreshTokenExpiry,
	}
}
