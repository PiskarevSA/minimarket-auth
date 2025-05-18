package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	chimiddlewares "github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v3"

	"github.com/PiskarevSA/minimarket-auth/internal/config"
	"github.com/PiskarevSA/minimarket-auth/internal/gen/oapi"
	"github.com/PiskarevSA/minimarket-auth/internal/handlers"
	"github.com/PiskarevSA/minimarket-auth/internal/repo/postgresql"
	"github.com/PiskarevSA/minimarket-auth/internal/usecases"
	"github.com/PiskarevSA/minimarket-auth/pkg/jwtmanager"
	pkgmiddlewares "github.com/PiskarevSA/minimarket-auth/pkg/middlewares"
)

var run = &cli.Command{
	Name:  "run",
	Usage: "Запуск minimarket-auth сервиса",
	Flags: runFlags,
	Action: func(ctx context.Context, cmd *cli.Command) error {
		const serviceName = "minimarket-auth"

		postgreSqlUrl := config.PostgreSqlConnUrl()
		pgxPool, err := pgxpool.New(ctx, postgreSqlUrl)
		if err != nil {
			log.Fatal().
				Err(err).
				Msg("failed to connect to postgresql")
		}

		router := chi.NewRouter()
		router.Use(chimiddlewares.Recoverer)
		router.Use(pkgmiddlewares.Decompress)

		postgreSql := postgresql.New(pgxPool)

		jwtSigningKeyBytes, err := os.ReadFile(config.JwtSigningKeyFilePath())
		if err != nil {
			log.Fatal().Err(err).Send()
		}
		jwtSigningKey := string(jwtSigningKeyBytes)

		jwtManager := jwtmanager.New(
			jwtSigningKey,
			config.JwtSigningMethod(),
			serviceName,
			config.JwtAccessTokenTtl(),
			config.JwtRefreshTokenTtl(),
		)

		login := usecases.NewLogin(
			serviceName,
			postgreSql,
			jwtManager)

		register := usecases.NewRegister(
			serviceName,
			postgreSql,
			jwtManager,
		)

		handlers := handlers.New(login, register)
		strictHandler := oapi.NewStrictHandler(handlers, nil)

		server := http.Server{
			Addr:         config.ServerAddr(),
			Handler:      oapi.HandlerFromMux(strictHandler, router),
			ReadTimeout:  config.ServerReadTimeout(),
			WriteTimeout: config.ServerWriteTimeout(),
			IdleTimeout:  config.ServerWriteTimeout(),
		}

		log.Info().
			Str("addr", config.ServerAddr()).
			Msg("listening server...")

		go func() {
			err := server.ListenAndServe()
			if err != nil && !errors.Is(err, http.ErrServerClosed) {
				log.Fatal().
					Err(err).
					Msg("failed to run server")
			}
		}()

		<-ctx.Done()

		log.Info().Msg("server stopped")

		return nil
	},
}

var runFlags = []cli.Flag{
	&cli.StringFlag{
		Name:        "server.addr",
		Value:       "127.0.0.1:8461",
		Destination: &config.Config().ServerAddr,
	},

	&cli.DurationFlag{
		Name:        "server.read-timeout",
		Value:       5 * time.Second,
		Usage:       "Server read timeout",
		Destination: &config.Config().ServerReadTimeout,
	},

	&cli.DurationFlag{
		Name:        "server.write-timeout",
		Value:       10 * time.Second,
		Usage:       "Server write timeout",
		Destination: &config.Config().ServerWriteTimeout,
	},

	&cli.DurationFlag{
		Name:        "server.idle-timeout",
		Value:       120 * time.Second,
		Usage:       "Server idle timeout",
		Destination: &config.Config().ServerIdleTimeout,
	},

	&cli.StringFlag{
		Name:        "postgresql.addr",
		Value:       "127.0.0.1:5432",
		Destination: &config.Config().PostgreSqlAddr,
	},

	&cli.StringFlag{
		Name:        "postgresql.user",
		Value:       "user",
		Destination: &config.Config().PostgreSqlUser,
	},

	&cli.StringFlag{
		Name:        "postgresql.password",
		Value:       "password",
		Destination: &config.Config().PostgreSqlPassword,
	},

	&cli.StringFlag{
		Name:        "postgresql.db",
		Value:       "postgres",
		Destination: &config.Config().PostgreSqlDb,
	},

	&cli.BoolFlag{
		Name:        "postgresql.sslmode",
		Value:       false,
		Destination: &config.Config().PostgreSqlSslMode,
	},

	&cli.StringFlag{
		Name:        "jwt.signkeyfilepath",
		Value:       "jwt.pem",
		Usage:       "path to file with sign key for ES256 algorithm",
		Destination: &config.Config().JwtSigningKeyFilePath,
	},

	&cli.StringFlag{
		Name:  "jwt.signing-method",
		Value: "ES256",
	},

	&cli.DurationFlag{
		Name:        "jwt.refresh-ttl",
		Value:       30 * 24 * time.Hour,
		Destination: &config.Config().JwtRefreshTokenTtl,
	},

	&cli.DurationFlag{
		Name:        "jwt.access-ttl",
		Value:       15 * time.Minute,
		Destination: &config.Config().JwtAccessTokenTtl,
	},
}
