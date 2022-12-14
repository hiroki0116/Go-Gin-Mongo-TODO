// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type NewTask struct {
	Title string `json:"title"`
}

type SignupInput struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type Task struct {
	ID            primitive.ObjectID `json:"_id"`
	Title         string             `json:"title"`
	Completed     bool               `json:"completed"`
	CompletedDate string             `json:"completed_date"`
	CreatedAt     string             `json:"created_at"`
	UpdatedAt     string             `json:"updated_at"`
	UserID        primitive.ObjectID `json:"user_id"`
}

type UpdateTask struct {
	ID        primitive.ObjectID `json:"_id"`
	Title     string             `json:"title"`
	Completed bool               `json:"completed"`
}

type UpdateUser struct {
	ID        primitive.ObjectID `json:"_id"`
	Email     string             `json:"email"`
	FirstName string             `json:"first_name"`
	LastName  string             `json:"last_name"`
}

type User struct {
	ID        primitive.ObjectID `json:"_id"`
	FirstName string             `json:"first_name"`
	LastName  string             `json:"last_name"`
	Email     string             `json:"email"`
	Password  string             `json:"password"`
	CreatedAt string             `json:"created_at"`
	UpdatedAt string             `json:"updated_at"`
	Token     string             `json:"token"`
}
