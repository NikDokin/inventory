package types

import "context"

type Tx interface {
	Exec(ctx context.Context, sql string, arguments ...any) (int64, error)
	Commit(ctx context.Context) error
}
