package signin

import (
	"context"
	"errors"
	"time"

	"github.com/PiskarevSA/minimarket-points/internal/storage"
	"github.com/PiskarevSA/minimarket-points/internal/usecases"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type Result struct {
	UserId       uuid.UUID
	Login        string
	AccessToken  string
	RefreshToken string
}

var nilResult = Result{}

func (u *Usecase) Do(
	ctx context.Context,
	rawLogin string,
	rawPassword string,
) (Result, error) {
	const op = "account.signin"

	args, err := newArgs(rawLogin, rawPassword)
	if err != nil {
		return nilResult, err
	}

	creds, err := u.storage.GetAccountCredentials(
		ctx,
		args.Login,
	)
	if err != nil {
		if errors.Is(err, storage.ErrAccountCredentialsNotFound) {
			return nilResult,
				&usecases.BusinessError{
					Code:    "D1932",
					Message: "invalid login or password",
				}
		}

		log.Error().
			Err(err).
			Str("op", op).
			Str("layer", "storage").
			Msg("failed to get account credentials")

		return nilResult, err
	}

	ok, err := args.Password.IsMatchHash(creds.PasswordHash)
	if err != nil {
		log.Error().
			Err(err).
			Str("op", op).
			Str("layer", "usecases").
			Msg("failed to match password and hash")

		return nilResult, err
	}

	if !ok {
		return nilResult,
			&usecases.BusinessError{
				Code:    "D1932",
				Message: "invalid login or password",
			}
	}

	userId := creds.UserId
	now := time.Now()

	accessToken, refreshToken, err := usecases.CreateTokenPair(
		userId.String(),
		u.jwtSignKey,
		u.jwtSigningMethod,
		u.accessTokenExpiry,
		u.refreshTokenExpiry,
		now,
	)
	if err != nil {
		log.Error().
			Err(err).
			Str("op", op).
			Str("layer", "usecases").
			Msg("failed to create token pair")

		return nilResult, err
	}

	return Result{
		UserId:       userId,
		Login:        args.Login.String(),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
