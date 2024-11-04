package server

import (
	"database/sql"
	"flag"
	"fmt"
	"net/http"
	"time"

	"spotfinder/internal/server/handlers"
)

type Application struct {
	db     *sql.DB // TODO: change after running SQLC
	port   int
	routes http.Handler
}

func NewServer() *http.Server {
	port := flag.Int("port", 5000, "Port to run the application on (default is 5000)")

	app := &Application{
		db:     newDB(),
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
