package server

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/http/pprof"

	"github.com/feynmaz/pkg/http/middleware"
	"github.com/feynmaz/pkg/logger"
	"github.com/fungicibus/inventory/config"
	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Server struct {
	cfg    *config.Config
	logger *logger.Logger
	v1     http.Handler
}

func New(cfg *config.Config, logger *logger.Logger, v1 http.Handler) *Server {
	return &Server{
		cfg:    cfg,
		logger: logger,
		v1:     v1,
	}
}

func (s *Server) Run(ctx context.Context) error {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%v", s.cfg.Server.Port),
		BaseContext:  func(_ net.Listener) context.Context { return ctx },
		Handler:      s.getRouter(),
		ReadTimeout:  s.cfg.Server.ReadTimeout,
		WriteTimeout: s.cfg.Server.WriteTimeout,
	}

	s.logger.Info().Msgf("server started on port %d", s.cfg.Server.Port)
	return srv.ListenAndServe()
}

func (s *Server) getRouter() *chi.Mux {
	router := chi.NewMux()

	// Middleware
	router.Use(middleware.RequestIDMiddleware)
	router.Use(middleware.NewLoggingMiddleware(s.logger))
	// router.Use(s.TelemetryMiddleware)
	registerMetrics(s.cfg.AppVersion, router)

	// Profiler
	router.HandleFunc("/debug/pprof/", pprof.Index)
	router.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	router.HandleFunc("/debug/pprof/profile", pprof.Profile)
	router.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	router.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
	router.Handle("/debug/pprof/heap", pprof.Handler("heap"))
	router.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
	router.Handle("/debug/pprof/block", pprof.Handler("block"))

	// Metrics
	router.Handle("/metrics", promhttp.Handler())

	// API
	router.Mount("/api/v1", s.v1)

	// Swagger
	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))
	router.Get("/swagger/doc.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, s.cfg.OpenapiPath)
	})

	return router
}

func (s *Server) Shutdown() {
	s.logger.Info().Msg("graceful server shutdown")
}
