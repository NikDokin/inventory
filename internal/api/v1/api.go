package v1

import (
	"net/http"

	"github.com/feynmaz/pkg/logger"
	"github.com/fungicibus/inventory/config"
)

type API struct {
	cfg    *config.Config
	logger *logger.Logger
	storage Storage
}

func New(cfg *config.Config, logger *logger.Logger, storage Storage) *API {
	return &API{
		cfg:    cfg,
		logger: logger,
		storage: storage,
	}
}

func (api *API) GetHandler() http.Handler {
	return Handler(api)
}
