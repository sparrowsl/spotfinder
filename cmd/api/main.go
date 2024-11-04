package main

import (
	"log/slog"

	"spotfinder/internal/server"
)

func main() {
	server := server.NewServer()

	if err := server.ListenAndServe(); err != nil {
		slog.Error(err.Error())
	}
}
