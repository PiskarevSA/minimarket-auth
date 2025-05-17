package usecases

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"

	"github.com/PiskarevSA/minimarket-auth/internal/domain/objects"
	"github.com/PiskarevSA/minimarket-auth/internal/repo"
	"github.com/PiskarevSA/minimarket-auth/pkg/jwtmanager"
)

type login struct {
	serviceName string
	accountRepo accountRepo
	jwtManager  *jwtmanager.JwtManager
}

func NewLogin(
	serviceName string,
	accountRepo accountRepo,
	jwtManager *jwtmanager.JwtManager,
) *login {
	return &login{
		serviceName: serviceName,
		accountRepo: accountRepo,
		jwtManager:  jwtManager,
	}
}

func (u *login) newLoginAndPasswrod(
	op string,
	rawLogin string,
	rawPassword string,
) (objects.Login, objects.Password, error) {
	login, err := objects.NewLogin(rawLogin)
	if err != nil {
		var loginErr *objects.LoginError
		if errors.As(err, &loginErr) {
			return objects.NilLogin, objects.NilPassword, &ValidationError{
				Code:    "V1142",
				Field:   "login",
				Message: err.Error(),
			}
		}

		log.Error().
			Err(err).
			Str("op", op).
			Str("layer", "usecase").
			Msg("failed to validate input")

		return objects.NilLogin, objects.NilPassword, err
	}

	password, err := objects.NewPassword(rawPassword)
	if err != nil {
		var passwordErr *objects.PasswordError
		if errors.As(err, &passwordErr) {
			return objects.NilLogin, objects.NilPassword, &ValidationError{
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

		return objects.NilLogin, objects.NilPassword, err
	}

	return login, password, nil
}

func (u *login) Do(
	ctx context.Context,
	rawLogin string,
	rawPassword string,
) (uuid.UUID, string, string, error) {
	const op = "login"

	login, password, err := u.newLoginAndPasswrod(op, rawLogin, rawPassword)
	if err != nil {
		return uuid.Nil, "", "", err
	}

	userId, passwordHash, err := u.accountRepo.GetUserIdAndPasswordHash(
		ctx,
		login,
	)
	if err != nil {
		if errors.Is(err, repo.ErrLoginNotExists) {
			return uuid.Nil, "", "", &BusinessError{
				Code:    "D1026",
				Message: "invalid login or password",
			}
		}

		log.Error().
			Err(err).
			Str("op", op).
			Str("layer", "repo").
			Msg("failed to get user id and password hash")

		return uuid.Nil, "", "", err
	}

	match, err := password.IsHashMatch(passwordHash)
	if err != nil {
		log.Error().
			Err(err).
			Str("op", op).
			Str("layer", "usecases").
			Msg("failed to match password and its hash")

		return uuid.Nil, "", "", err
	}

	if !match {
		return uuid.Nil, "", "", &BusinessError{
			Code:    "D1026",
			Message: "invalid login or password",
		}
	}

	return userId, "", "", err
}
