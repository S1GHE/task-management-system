package app

import (
	"task-management-system-api/internal/handler"
	"task-management-system-api/internal/infra/logging"
	"task-management-system-api/internal/server"
)

func Start() error {
	log := logging.MustLog("dev")
	handlers := handler.New(log)
	srv := server.New("8000", handlers.InitHandler(), log)

	if err := srv.Run(); err != nil {
		return err
	}

	return nil
}
