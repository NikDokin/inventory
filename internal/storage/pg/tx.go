package pg

import (
	"context"

	"github.com/fungicibus/inventory/internal/types"
	"github.com/jackc/pgx/v5"
)

type Tx struct {
	tx pgx.Tx
}

func (pg *Adapter) BeginTx(ctx context.Context) (types.Tx, error) {
	tx, err := pg.rwPool.Begin(ctx)
	return &Tx{
		tx: tx,
	}, err
}

func (tx Tx) Exec(ctx context.Context, sql string, arguments ...any) (int64, error) {
	commandTag, err := tx.tx.Exec(ctx, sql, arguments...)
	return commandTag.RowsAffected(), err
}

func (tx Tx) Commit(ctx context.Context) error {
	return tx.tx.Commit(ctx)
}
