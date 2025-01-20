package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/feynmaz/pkg/http/middleware"
)

func (api *API) WriteJSON(w http.ResponseWriter, r *http.Request, response any) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		api.WriteError(w, r, WithError(fmt.Errorf("failed to encode response body: %w", err)))
	}
}

func (api *API) WriteError(w http.ResponseWriter, r *http.Request, opts ...ErrorOption) {
	// Apply configs
	config := &ErrorConfig{
		StatusCode: http.StatusInternalServerError,
		Detail:     "Server Error",
	}
	for _, opt := range opts {
		opt(config)
	}

	// Log
	requestId := middleware.GetRequestID(r.Context())
	logEvent := api.logger.Error().
		Int("statusCode", config.StatusCode).
		Str("detail", config.Detail).
		Str("id", requestId)

	if config.Error != nil {
		logEvent.Err(config.Error)
	}

	logEvent.Send()

	// Send http response
	now := time.Now().Format(time.RFC3339)
	response := ErrorResponse{
		Detail: config.Detail,
		Id:     requestId,
		Status: strconv.Itoa(config.StatusCode),
		Meta: ErrorMeta{
			Timestamp: &now,
		},
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(config.StatusCode)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		api.logger.Error().Err(err).Msg("failed to encode error response body")
		w.WriteHeader(http.StatusInternalServerError)
	}
}

type ErrorOption func(*ErrorConfig)
type ErrorConfig struct {
	Error      error  `json:"error"`
	StatusCode int    `json:"statusCode"`
	Detail     string `json:"detail"`
}

func WithError(err error) ErrorOption {
	return func(c *ErrorConfig) {
		c.Error = err
	}
}

func WithStatusCode(statusCode int) ErrorOption {
	return func(c *ErrorConfig) {
		c.StatusCode = statusCode
	}
}

func WithDetail(detail string) ErrorOption {
	return func(c *ErrorConfig) {
		c.Detail = detail
	}
}
