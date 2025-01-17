package app

import (
	"task-management-system-api/internal/handler"
	"task-management-system-api/internal/server"
)

func Start() error {
	handlers := handler.New()
	srv := server.New("8000", handlers.InitHandler())

	if err := srv.Run(); err != nil {
		return err
	}

	return nil
}
