package transactor

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"

	"github.com/PiskarevSA/minimarket-points/pkg/ctxkey"
)

var (
	TxCtxKey     = ctxkey.New[pgx.Tx]("transactor.tx", nil)
	ErrNoTxInCtx = errors.New("no transaction in context")
)

func (t *Transactor) Transact(
	ctx context.Context,
	opts pgx.TxOptions,
	fn func(ctx context.Context) error,
) error {
	tx, err := t.driver.BeginTx(ctx, opts)
	if err != nil {
		return err
	}

	defer func() {
		p := recover()
		if p != nil {
			// For simplicity we assume that the rollback will be executed in
			// 100% cases. For production use it's needed to implement a
			// mechanism to handle rollback errors. Descriptions of strategies
			// that help to deal with such errors can be in this paper:
			// https://nthu-datalab.github.io/db/slides/10_Transaction_Recovery.pdf
			tx.Rollback(ctx)
			panic(p)
		}
	}()

	txCtx := TxCtxKey.WithValue(ctx, tx)
	err = fn(txCtx)
	if err != nil {
		tx.Rollback(ctx)

		return err
	}

	return tx.Commit(ctx)
}
