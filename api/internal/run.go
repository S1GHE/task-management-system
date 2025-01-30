package internal

import (
	"log/slog"
	"net/http"
	"task-management-system-api/internal/applications"
	"task-management-system-api/pkg/logger"
	"time"
)

func Run(cfg *Config) error {
	handler := applications.NewHTTPServer()
	logger.Setup()

	slog.Info("SERVER START")
	if err := startServer(cfg.Server.Port, handler.SetupHTTPServer()); err != nil {
		return err
	}

	return nil
}

func startServer(port string, handler http.Handler) error {
	server := &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	if err := server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
