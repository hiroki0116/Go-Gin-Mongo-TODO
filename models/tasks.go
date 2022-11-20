package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	ID            primitive.ObjectID `json:"id" bson:"_id"`
	Title         string             `json:"title" bson:"title" binding:"required,max=250"`
	Completed     bool               `json:"completed" bson:"completed"`
	CompletedDate string             `json:"completed_date" bson:"completed_date"`
	CreatedAt     string             `json:"created_at" bson:"created_at"`
	UpdatedAt     string             `json:"updated_at" bson:"updated_at"`
}

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	FirstName string             `json:"first_name" bson:"first_name"`
	LastName  string             `json:"last_name" bson:"last_name"`
	Email     string             `json:"email" bson:"email"`
	Password  string             `json:"password" bson:"password"`
	CreatedAt string             `json:"created_at" bson:"created_at"`
}
