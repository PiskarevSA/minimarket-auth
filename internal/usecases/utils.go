package usecases

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateAccessToken(
	userId string,
	signKey any,
	signingMethod jwt.SigningMethod,
	expiry time.Duration,
	now time.Time,
) (string, error) {
	iat := now.UTC().Unix()
	exp := now.Add(time.Hour * time.Duration(expiry)).Unix()

	token := jwt.NewWithClaims(
		signingMethod,
		jwt.MapClaims{
			"iss": "auth-service",
			"sub": userId,
			"exp": exp,
			"iat": iat,
		},
	)

	tokenString, err := token.SignedString(signKey)
	if err != nil {
		return "", fmt.Errorf(
			"failed to create access token: %w",
			err,
		)
	}

	return tokenString, err
}

func CreateRefreshToken(
	userId string,
	signKey any,
	signingMethod jwt.SigningMethod,
	expiry time.Duration,
	now time.Time,
) (string, error) {
	iat := now.UTC().Unix()
	exp := now.Add(time.Hour * time.Duration(expiry)).Unix()

	token := jwt.NewWithClaims(
		signingMethod,
		jwt.MapClaims{
			"iss": "auth-service",
			"sub": userId,
			"exp": exp,
			"iat": iat,
		},
	)

	tokenString, err := token.SignedString(signKey)
	if err != nil {
		return "", fmt.Errorf(
			"failed to create refresh token: %w",
			err,
		)
	}

	return tokenString, err
}

func CreateTokenPair(
	userId string,
	signKey any,
	signingMethod jwt.SigningMethod,
	accessTokenExpiry time.Duration,
	refreshTokenExpiry time.Duration,
	now time.Time,
) (string, string, error) {
	accessToken, err := CreateAccessToken(
		userId,
		signKey,
		signingMethod,
		accessTokenExpiry,
		now,
	)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := CreateRefreshToken(
		userId,
		signKey,
		signingMethod,
		refreshTokenExpiry,
		now,
	)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
