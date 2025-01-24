package main

import (
	"context"
	"embed"
	"encoding/json"
	"os"
	"os/signal"

	"github.com/feynmaz/pkg/logger"
	"github.com/fungicibus/inventory/config"
	v1 "github.com/fungicibus/inventory/internal/api/v1"
	"github.com/fungicibus/inventory/internal/server"
	"github.com/fungicibus/inventory/internal/storage/migrations"
	mockStorage "github.com/fungicibus/inventory/internal/storage/mock"
	"github.com/fungicibus/inventory/internal/storage/pg"
	"github.com/jackc/pgx/v5/stdlib"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func main() {
	log := logger.New()

	cfg, err := config.GetDefault()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to get config")
	}
	log.SetLevel(cfg.LogLevel)

	prettyJSON, err := json.MarshalIndent(cfg, "", "    ")
	if err != nil {
		log.Fatal().Err(err).Msg("failed to marshal config")
	}
	log.Debug().Msgf("Config: %s", string(prettyJSON))

	pg, err := pg.New(cfg.Postgres)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create postgres adapter")
	}
	_ = pg

	conn := stdlib.OpenDBFromPool(pg.RwPool())
	if err := migrations.Up(embedMigrations, conn); err != nil {
		log.Fatal().Err(err).Msg("failed to up migrations")
	}

	mockStorage := mockStorage.New()
	v1 := v1.New(cfg, log, mockStorage)

	server := server.New(cfg, log, v1.GetHandler())

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	go func() {
		err := server.Run(ctx)
		if err != nil {
			log.Fatal().Err(err).Msg("server error")
		}
	}()

	<-ctx.Done()
	pg.Close()
	server.Shutdown()
}
