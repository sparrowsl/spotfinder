package main

import (
	"log/slog"

	"spotfinder/internal/server"
)

func main() {
	server := server.NewServer()

	slog.Info("Server started on port: 5000")
	if err := server.ListenAndServe(); err != nil {
		slog.Error(err.Error())
	}
}
