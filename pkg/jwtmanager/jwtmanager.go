package jwtmanager

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtManager struct {
	signingKey    any
	signingMethod jwt.SigningMethod
	issuer        string
	accessTtl     time.Duration
	refreshTtl    time.Duration
}

func New(
	signingKey any,
	signingMethod jwt.SigningMethod,
	issuer string,
	accessTtl time.Duration,
	refreshTtl time.Duration,
) *JwtManager {
	return &JwtManager{
		signingKey:    signingKey,
		signingMethod: signingMethod,
		issuer:        issuer,
		accessTtl:     accessTtl,
		refreshTtl:    refreshTtl,
	}
}

func (m *JwtManager) IssueAccessToken(
	claims jwt.Claims,
	now time.Time,
) (string, error) {
	return m.issueToken(claims, now, m.accessTtl)
}

func (m *JwtManager) IssueRefreshToken(
	claims jwt.Claims,
	now time.Time,
) (string, error) {
	return m.issueToken(claims, now, m.refreshTtl)
}

func (m *JwtManager) IssueTokenPair(
	accessClaims jwt.Claims,
	refreshClaims jwt.Claims,
	now time.Time,
) (accessToken string, refreshToken string, err error) {
	accessToken, err = m.issueToken(accessClaims, now, m.accessTtl)
	if err != nil {
		return "", "", fmt.Errorf("failed to issue access token: %w", err)
	}

	refreshToken, err = m.issueToken(refreshClaims, now, m.refreshTtl)
	if err != nil {
		return "", "", fmt.Errorf("failed to issue refresh token: %w", err)
	}

	return accessToken, refreshToken, nil
}

func (m *JwtManager) issueToken(
	claims jwt.Claims,
	now time.Time,
	ttl time.Duration,
) (string, error) {
	rc, ok := claims.(interface {
		SetIssuer(string)
		SetIssuedAt(*jwt.NumericDate)
		SetExpiresAt(*jwt.NumericDate)
	})
	if ok {
		rc.SetIssuer(m.issuer)
		rc.SetIssuedAt(jwt.NewNumericDate(now.UTC()))
		rc.SetExpiresAt(jwt.NewNumericDate(now.UTC().Add(ttl)))
	}

	token := jwt.NewWithClaims(m.signingMethod, claims)
	tokenString, err := token.SignedString(m.signingKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
