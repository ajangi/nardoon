package models

import "time"

// Role : this struct is for the role model in this project
type Role struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
