package pg

import (
	"context"
	"fmt"

	"github.com/fungicibus/inventory/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Adapter struct {
	roPool *pgxpool.Pool
	rwPool *pgxpool.Pool
}

func New(cfg config.Postgres) (Adapter, error) {
	ctx, cancel := context.WithTimeout(context.Background(), cfg.ConnectTimeout)
	defer cancel()

	roPool, err := pgxpool.New(ctx, cfg.RoDSN)
	if err != nil {
		return Adapter{}, fmt.Errorf("failed to create ro pool: %w", err)
	}
	if err := roPool.Ping(ctx); err != nil {
		return Adapter{}, fmt.Errorf("failed to ping ro pool: %w", err)
	}

	rwPool, err := pgxpool.New(ctx, cfg.RwDSN)
	if err != nil {
		return Adapter{}, fmt.Errorf("failed to create rw pool: %w", err)
	}
	if err := rwPool.Ping(ctx); err != nil {
		return Adapter{}, fmt.Errorf("failed to ping rw pool: %w", err)
	}

	return Adapter{
		roPool: roPool,
		rwPool: rwPool,
	}, nil
}

func (s *Adapter) Close() {
	s.roPool.Close()
	s.rwPool.Close()
}
