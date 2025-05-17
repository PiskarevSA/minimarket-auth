package postgresql

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgconn"

	"github.com/PiskarevSA/minimarket-auth/internal/domain/entities"
	"github.com/PiskarevSA/minimarket-auth/internal/gen/sqlc/postgresql"
	"github.com/PiskarevSA/minimarket-auth/internal/repo"
	"github.com/PiskarevSA/minimarket-auth/pkg/pgcodes"
	"github.com/PiskarevSA/minimarket-auth/pkg/pgx/transactor"
)

func (r *PostgreSql) CreateAccountInTx(
	ctx context.Context,
	account entities.Account,
) error {
	tx := transactor.TxCtxKey.Value(ctx)

	err := r.querier.WithTx(tx).CreateAccount(
		ctx,
		postgresql.CreateAccountParams{
			Id:           account.Id(),
			Login:        account.Login().String(),
			PasswordHash: string(account.PasswordHash()),
			CreatedAt:    account.CreatedAt(),
			UpdatedAt:    account.UpdatedAt(),
		},
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) ||
			pgcodes.IsUniqueViolation(pgErr.Code) {
			return repo.ErrLoginAlreadyInUse
		}

		return err
	}

	return nil
}
