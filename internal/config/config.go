package config

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var conf config

type config struct {
	LogLevel              string
	ServerAddr            string
	ServerReadTimeout     time.Duration
	ServerWriteTimeout    time.Duration
	ServerIdleTimeout     time.Duration
	PostgreSqlAddr        string
	PostgreSqlUser        string
	PostgreSqlPassword    string
	PostgreSqlDb          string
	PostgreSqlSslMode     bool
	JwtSigningKeyFilePath string
	JwtSigningMethod      jwt.SigningMethod
	JwtAccessTokenTtl     time.Duration
	JwtRefreshTokenTtl    time.Duration
}

func Config() *config {
	return &conf
}

func LogLevel() string {
	return conf.LogLevel
}

func ServerAddr() string {
	return conf.ServerAddr
}

func ServerReadTimeout() time.Duration {
	return conf.ServerReadTimeout
}

func ServerWriteTimeout() time.Duration {
	return conf.ServerWriteTimeout
}

func ServerIdleTimeout() time.Duration {
	return conf.ServerIdleTimeout
}

func PostgreSqlConnUrl() string {
	connUrl := fmt.Sprintf(
		"postgres://%s:%s@%s/%s",
		conf.PostgreSqlUser,
		conf.PostgreSqlPassword,
		conf.PostgreSqlAddr,
		conf.PostgreSqlDb,
	)

	if conf.PostgreSqlSslMode {
		connUrl += "?sslmode=disable"
	}

	return connUrl
}

func JwtSigningKeyFilePath() string {
	return conf.JwtSigningKeyFilePath
}

func JwtSigningMethod() jwt.SigningMethod {
	return conf.JwtSigningMethod
}

func JwtAlgo() jwt.SigningMethod {
	return jwt.SigningMethodES256
}

func JwtAccessTokenTtl() time.Duration {
	return conf.JwtAccessTokenTtl
}

func JwtRefreshTokenTtl() time.Duration {
	return conf.JwtRefreshTokenTtl
}
