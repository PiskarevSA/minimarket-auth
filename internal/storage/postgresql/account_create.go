package postgresql

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/PiskarevSA/minimarket-points/internal/domain/objects"
	"github.com/PiskarevSA/minimarket-points/internal/gen/sqlc/postgresql"
	"github.com/PiskarevSA/minimarket-points/internal/storage"
	"github.com/PiskarevSA/minimarket-points/pkg/pgcodes"
	"github.com/PiskarevSA/minimarket-points/pkg/pgx/transactor"
)

func (s *Storage) CreateAccountInTx(
	ctx context.Context,
	userId uuid.UUID,
	login objects.Login,
	passwordHash []byte,
	createdAt time.Time,
) error {
	tx, ok := transactor.TxCtxKey.ValueOk(ctx)
	if !ok {
		return transactor.ErrNoTxInCtx
	}

	query := s.querier.WithTx(tx)

	err := query.CreateAccount(
		ctx,
		postgresql.CreateAccountParams{
			Id:           userId,
			Login:        login.String(),
			PasswordHash: string(passwordHash),
			CreatedAt:    createdAt,
		},
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgcodes.IsUniqueViolation(pgErr.Code) {
				return storage.ErrLoginAlreadyExists
			}
		}

		return err
	}

	return nil
}
