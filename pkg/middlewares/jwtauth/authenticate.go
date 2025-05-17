package jwtauth

import (
	"context"
	"errors"
	"net/http"

	"github.com/golang-jwt/jwt/v5"

	"github.com/PiskarevSA/minimarket-auth/pkg/ctxkey"
)

var JwtCtxKey = ctxkey.New("jwtauth.jwt", &jwt.Token{})

func parseTokenString(
	ctx context.Context,
	ja *JwtAuth,
	tokenString string,
) (*jwt.Token, error) {
	var (
		token *jwt.Token
		err   error
	)

	if ja.Claims != nil {
		token, err = jwt.ParseWithClaims(
			tokenString,
			ja.Claims(),
			ja.ParseFn(ctx, ja),
		)
	} else {
		token, err = jwt.Parse(tokenString, ja.ParseFn(ctx, ja))
	}

	if err != nil {
		if errors.Is(err, jwt.ErrTokenMalformed) ||
			errors.Is(err, jwt.ErrTokenUnverifiable) {
			return nil, ErrInvalidJwt
		}

		if errors.Is(err, jwt.ErrTokenNotValidYet) {
			return nil, ErrJwtNotYetValid
		}

		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrJwtExpired
		}

		return nil, ErrFailedToParseJwt
	}

	return token, nil
}

func validateToken(ctx context.Context, ja *JwtAuth, token *jwt.Token) error {
	if !token.Valid {
		return ErrInvalidJwt
	}

	if token.Method != ja.SigningMethod {
		return ErrUnsupportedSigningMethod
	}

	for _, validator := range ja.ValidatorFns {
		err := validator(ctx, token)
		if err != nil {
			return err
		}
	}

	return nil
}

func Authenticate(
	ja *JwtAuth,
	extractor Extractor,
) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(rw http.ResponseWriter, req *http.Request) {
			tokenString := extractor(req)
			if tokenString == "" {
				http.Error(rw, ErrMissingJwt.Error(), http.StatusUnauthorized)

				return
			}

			ctx := req.Context()

			token, err := parseTokenString(ctx, ja, tokenString)
			if err != nil {
				http.Error(rw, err.Error(), http.StatusUnauthorized)

				return
			}

			err = validateToken(ctx, ja, token)
			if err != nil {
				http.Error(rw, err.Error(), http.StatusUnauthorized)

				return
			}

			ctx = JwtCtxKey.WithValue(ctx, token)
			req = req.WithContext(ctx)

			next.ServeHTTP(rw, req)
		}

		return http.HandlerFunc(fn)
	}
}
