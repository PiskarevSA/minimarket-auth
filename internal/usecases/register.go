package usecases

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"

	"github.com/PiskarevSA/minimarket-auth/internal/domain/entities"
	"github.com/PiskarevSA/minimarket-auth/internal/domain/objects"
	"github.com/PiskarevSA/minimarket-auth/internal/events"
	"github.com/PiskarevSA/minimarket-auth/internal/repo"
	"github.com/PiskarevSA/minimarket-auth/pkg/jwtmanager"
	"github.com/PiskarevSA/minimarket-auth/pkg/pgx/transactor"
)

type register struct {
	serviceName string
	accountRepo accountRepo
	outboxRepo  outboxRepo
	transactor  *transactor.Transactor
	jwtManager  *jwtmanager.JwtManager
}

func NewRegister(
	serviceName string,
	accountRepo accountRepo,
	outboxRepo outboxRepo,
	transactor *transactor.Transactor,
	jwtManager *jwtmanager.JwtManager,

) *register {
	return &register{
		serviceName: serviceName,
		accountRepo: accountRepo,
		outboxRepo:  outboxRepo,
		transactor:  transactor,
		jwtManager:  jwtManager,
	}
}

func (u *register) newAccount(
	op,
	login,
	password string,
	now time.Time,
) (entities.Account, error) {
	id := uuid.New()

	account, err := entities.NewAccount(id, login, password, now, now)
	if err != nil {
		var loginErr *objects.LoginError
		if errors.As(err, &loginErr) {
			return account, &ValidationError{
				Code:    "V1142",
				Field:   "login",
				Message: err.Error(),
			}
		}

		var passwordErr *objects.PasswordError
		if errors.As(err, &passwordErr) {
			return account, &ValidationError{
				Code:    "V1078",
				Field:   "password",
				Message: err.Error(),
			}
		}

		log.Error().
			Err(err).
			Str("op", op).
			Str("layer", "usecase").
			Msg("failed to validate input")

		return account, err
	}

	return account, nil
}

func (u *register) newEvent(
	op string,
	account entities.Account,
	now time.Time,
) (events.Event, error) {
	event, err := events.NewAccountRegistered(
		account.Id(),
		account.Login().String(),
		u.serviceName,
		now,
	)
	if err != nil {
		log.Error().
			Err(err).
			Str("op", op).
			Str("layer", "usecases").
			Msg("failed to marshal event")

		return event, err
	}

	return event, nil
}

func (u *register) Do(
	ctx context.Context,
	rawLogin string,
	rawPassword string,
) (uuid.UUID, string, string, error) {
	const op = "register"

	now := time.Now()
	account, err := u.newAccount(op, rawLogin, rawPassword, now)
	if err != nil {
		return uuid.Nil, "", "", err
	}

	event, err := u.newEvent(op, account, now)
	if err != nil {
		return uuid.Nil, "", "", err
	}

	pgxTxOpts := pgx.TxOptions{IsoLevel: pgx.ReadCommitted}
	err = u.transactor.Transact(
		ctx,
		pgxTxOpts,
		func(ctx context.Context) error {
			err = u.accountRepo.CreateAccountInTx(ctx, account)
			if err != nil {
				return err
			}

			err = u.outboxRepo.CreateOutboxInTx(ctx, event)
			if err != nil {
				return err
			}

			return nil
		},
	)
	if err != nil {
		if errors.Is(err, repo.ErrLoginAlreadyInUse) {
			return uuid.Nil, "", "", &BusinessError{
				Code: "D1002",
				Message: fmt.Sprintf(
					"login %s already in use ",
					account.Login().String(),
				),
			}
		}

		log.Error().
			Err(err).
			Str("op", op).
			Str("layer", "repo").
			Msg("failed to create account")

		return uuid.Nil, "", "", err
	}

	return uuid.Nil, "", "", nil
}
