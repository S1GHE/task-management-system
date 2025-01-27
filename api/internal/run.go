package internal

import (
	"net/http"
	"task-management-system-api/internal/applications"
	"time"
)

func Run() error {
	handler := applications.NewHTTPServer()

	if err := startServer("8000", handler.SetupHTTPServer()); err != nil {
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
