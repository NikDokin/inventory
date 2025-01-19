package middleware

import (
	"net/http"

	"github.com/feynmaz/pkg/logger"
)

func NewLoggingMiddleware(logger *logger.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestId := GetRequestID(r.Context())
			next.ServeHTTP(w, r)
			logger.Info().Str("requestID", requestId).Msgf("%s %s", r.Method, r.URL.Path)
		})
	}
}
