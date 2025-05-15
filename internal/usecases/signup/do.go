package signup

import (
	"context"
	"errors"
	"time"

	"github.com/PiskarevSA/minimarket-points/internal/events"
	"github.com/PiskarevSA/minimarket-points/internal/storage"
	"github.com/PiskarevSA/minimarket-points/internal/usecases"
	json "github.com/bytedance/sonic"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"

	"github.com/jackc/pgx/v5"
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
	const op = "account.signup"

	args, err := newArgs(rawLogin, rawPassword)
	if err != nil {
		return nilResult, err
	}

	userId := uuid.New()
	now := time.Now().UTC()

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

	login := args.Login

	eventPayload, err := json.ConfigDefault.Marshal(
		events.AccountRegisteredPayload{
			UserId: userId.String(),
			Login:  login.String(),
		},
	)
	if err != nil {
		log.Error().
			Err(err).
			Str("op", op).
			Str("layer", "usecases").
			Msg("failed to marshal event")

		return nilResult, err
	}

	fn := func(ctx context.Context) error {
		err := u.storage.CreateAccountInTx(
			ctx,
			userId,
			login,
			args.Password.Hash(),
			now,
		)
		if err != nil {
			return err
		}

		return u.storage.CreateOutboxInTx(
			ctx,
			events.AccountRegistered,
			eventPayload,
			now,
			u.serviceName,
		)
	}

	pgxTxOpts := pgx.TxOptions{IsoLevel: pgx.ReadCommitted}

	err = u.transactor.Transact(ctx, pgxTxOpts, fn)
	if err != nil {
		if errors.Is(err, storage.ErrLoginAlreadyExists) {
			return nilResult,
				usecases.ErrLoginAlreadyExists(login.String())
		}

		log.Error().
			Err(err).
			Str("op", op).
			Str("layer", "storage").
			Msg("failed to create account")

		return nilResult, err
	}

	return Result{
		UserId:       userId,
		Login:        args.Login.String(),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
