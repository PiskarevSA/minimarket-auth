package postgresql

import (
	"context"
	"time"

	"github.com/PiskarevSA/minimarket-points/internal/events"
	"github.com/PiskarevSA/minimarket-points/internal/gen/sqlc/postgresql"
	"github.com/PiskarevSA/minimarket-points/pkg/pgx/transactor"
)

func (s *Storage) CreateOutboxInTx(
	ctx context.Context,
	event events.Event,
	payload []byte,
	createdAt time.Time,
	createdBy string,
) error {
	tx, ok := transactor.TxCtxKey.ValueOk(ctx)
	if !ok {
		return transactor.ErrNoTxInCtx
	}

	query := s.querier.WithTx(tx)

	return query.CreateOutbox(
		ctx,
		postgresql.CreateOutboxParams{
			Event:     event.String(),
			Payload:   payload,
			CreatedAt: createdAt,
			CreatedBy: createdBy,
			UpdatedAt: createdAt,
			UpdatedBy: createdBy,
		},
	)
}
