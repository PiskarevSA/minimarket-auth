package postgresql

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/PiskarevSA/minimarket-auth/internal/domain/objects"
	"github.com/PiskarevSA/minimarket-auth/internal/repo"
)

func (r *PostgreSql) GetUserIdAndPasswordHash(
	ctx context.Context,
	login objects.Login,
) (uuid.UUID, []byte, error) {
	result, err := r.querier.GetUserIdAndPasswordHash(
		ctx,
		login.String(),
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return uuid.Nil, nil, repo.ErrLoginNotExists
		}

		return uuid.Nil, nil, err
	}

	passwordHash := []byte(result.PasswordHash)

	return result.UserId, passwordHash, nil
}
