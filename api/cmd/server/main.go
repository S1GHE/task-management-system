package main

import (
	"log/slog"
	"os"
	"task-management-system-api/internal"
)

func main() {
	cfg := internal.LoadConfig()

	if err := internal.Run(cfg); err != nil {
		slog.Error("Failed to run server", "err", err)
		os.Exit(1)
	}
}
