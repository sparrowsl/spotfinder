package server

import (
	"database/sql"
	"fmt"
	"os"

	_ "modernc.org/sqlite"
)

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
