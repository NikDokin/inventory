package main

import (
	"context"
	"encoding/json"
	"os"
	"os/signal"

	"github.com/feynmaz/pkg/logger"
	"github.com/fungicibus/inventory/config"
	v1 "github.com/fungicibus/inventory/internal/api/v1"
	"github.com/fungicibus/inventory/internal/server"
	mockStorage "github.com/fungicibus/inventory/internal/storage/mock"
)

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
	log.Debug().Msgf("Config: /n%s", string(prettyJSON))

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
	server.Shutdown()
}
