package server

import (
	"database/sql"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"spotfinder/internal/database"
	"spotfinder/internal/server/handlers"

	_ "modernc.org/sqlite"
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

var dbName = "spotfinder.db"

func newDB() *sql.DB {
	// TODO: accept database from flags
	// TODO: check if accepted name is an actual database that ends in ".db"

	db, err := sql.Open("sqlite", dbName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Check if connection to db is open/valid
	if err := db.Ping(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	return db
}
