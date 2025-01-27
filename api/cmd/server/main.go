package main

import (
	"log/slog"
	"os"
	"task-management-system-api/internal"
)

func main() {
	if err := internal.Run(); err != nil {
		slog.Error("Failed to run server", "err", err)
		os.Exit(1)
	}
}
