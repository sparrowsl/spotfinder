package server

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"spotfinder/internal/database"
	"spotfinder/internal/server/handlers"
)

type Application struct {
	db     *database.Queries
	port   int
	routes http.Handler
}

func NewServer() *http.Server {
	port := flag.Int("port", 5000, "Port to run the application on (default is 5000)")

	app := &Application{
		db:     database.New(newDB()),
		port:   *port,
		routes: handlers.RegisterRoutes(),
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.port),
		Handler:      app.routes,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Minute,
	}

	return server
}
