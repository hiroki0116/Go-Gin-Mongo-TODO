// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type NewTask struct {
	Title string `json:"title"`
}

type Task struct {
	ID            string    `json:"_id"`
	Title         string    `json:"title"`
	Completed     bool      `json:"completed"`
	CompletedDate time.Time `json:"completed_date"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	UserID        string    `json:"user_id"`
}

type UpdateTask struct {
	ID        string `json:"_id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
