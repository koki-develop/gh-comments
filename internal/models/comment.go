package models

import "time"

type User struct {
	Login string `json:"login"`
	ID    int64  `json:"id"`
}

type Comment struct {
	ID        int64     `json:"id"`
	Body      string    `json:"body"`
	User      User      `json:"user"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
