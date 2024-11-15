package models

import "time"

type Location struct {
	ID          int64     `json:"id"`
	Latitude    float64   `json:"latitude"`
	Longitude   float64   `json:"longitude"`
	Address     string    `json:"address"`
	Category    string    `json:"category"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
