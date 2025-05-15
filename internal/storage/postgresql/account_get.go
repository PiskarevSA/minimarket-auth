package postgresql

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"

	"github.com/PiskarevSA/minimarket-points/internal/domain/entities"
	"github.com/PiskarevSA/minimarket-points/internal/domain/objects"
	"github.com/PiskarevSA/minimarket-points/internal/storage"
)

func (s *Storage) GetAccountCredentials(
	ctx context.Context,
	login objects.Login) (entities.Credentials, error) {
	creds, err := s.querier.GetAccountCredentials(ctx, login.String())
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entities.NilCredentials, storage.ErrAccountCredentialsNotFound
		}

		return entities.NilCredentials, err
	}

	return entities.Credentials{
		UserId:       creds.UserId,
		PasswordHash: []byte(creds.PasswordHash),
	}, nil
}
