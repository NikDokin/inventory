package v1

import (
	"net/http"

	"github.com/fungicibus/inventory/config"
	"github.com/fungicibus/inventory/internal/logger"
)

type API struct {
	cfg    *config.Config
	logger *logger.Logger
}

func New(cfg *config.Config, logger *logger.Logger) *API {
	return &API{
		cfg:    cfg,
		logger: logger,
	}
}

func (api *API) GetHandler() http.Handler {
	return Handler(api)
}
