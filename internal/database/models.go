// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"database/sql"
)

type Location struct {
	ID          int64          `db:"id" json:"id"`
	Latitude    int64          `db:"latitude" json:"latitude"`
	Longitude   int64          `db:"longitude" json:"longitude"`
	Address     string         `db:"address" json:"address"`
	Category    sql.NullString `db:"category" json:"category"`
	Description sql.NullString `db:"description" json:"description"`
	CreatedAt   sql.NullTime   `db:"created_at" json:"created_at"`
	UpdatedAt   sql.NullTime   `db:"updated_at" json:"updated_at"`
}