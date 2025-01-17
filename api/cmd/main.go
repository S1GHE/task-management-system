package main

import (
	"log"
	"task-management-system-api/internal/app"
)

func main() {
	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
