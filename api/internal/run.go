package internal

import (
	"log/slog"
	"net/http"
	"task-management-system-api/internal/applications"
	"time"
)

func Run() error {
	handler := applications.NewHTTPServer()

	if err := startServer("8080", handler.SetupHTTPServer()); err != nil {
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

	slog.Info("Server start")
	if err := server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
