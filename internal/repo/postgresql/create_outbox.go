package postgresql

import (
	"context"

	"github.com/PiskarevSA/minimarket-auth/internal/events"
	"github.com/PiskarevSA/minimarket-auth/internal/gen/sqlc/postgresql"
	"github.com/PiskarevSA/minimarket-auth/pkg/pgx/transactor"
)

func (r *PostgreSql) CreateOutboxInTx(
	ctx context.Context,
	event events.Event,
) error {
	tx := transactor.TxCtxKey.Value(ctx)

	return r.querier.WithTx(tx).CreateOutbox(
		ctx,
		postgresql.CreateOutboxParams{
			Eventname: event.GetName(),
			Payload:   event.GetPayload(),
			CreatedAt: event.GetCreatedAt(),
			CreatedBy: event.GetCreatedBy(),
			UpdatedAt: event.GetUpdatedAt(),
			UpdatedBy: event.GetUpdatedBy(),
		},
	)
}
