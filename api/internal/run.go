package internal

import (
	"net/http"
	"task-management-system-api/internal/applications"
	"task-management-system-api/pkg/logger"
	"time"
)

func Run(cfg *Config) error {
	handler := applications.NewHTTPServer()
	logger.SetupLog(cfg.Mod)

	if err := startServer(cfg.Port, handler.SetupHTTPServer()); err != nil {
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
