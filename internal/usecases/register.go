package usecases

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"

	"github.com/PiskarevSA/minimarket-auth/internal/domain/entities"
	"github.com/PiskarevSA/minimarket-auth/internal/domain/objects"
	"github.com/PiskarevSA/minimarket-auth/internal/repo"
	"github.com/PiskarevSA/minimarket-auth/pkg/jwtmanager"
)

type register struct {
	serviceName string
	accountRepo accountRepo
	jwtManager  *jwtmanager.JwtManager
}

func NewRegister(
	serviceName string,
	accountRepo accountRepo,
	jwtManager *jwtmanager.JwtManager,
) *register {
	return &register{
		serviceName: serviceName,
		accountRepo: accountRepo,
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

	err = u.accountRepo.CreateAccountInTx(ctx, account)
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
