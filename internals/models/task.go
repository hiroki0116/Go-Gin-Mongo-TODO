package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title         string             `json:"title,omitempty" bson:"title,omitempty" binding:"max=250"`
	Completed     bool               `json:"completed,omitempty" bson:"completed,omitempty"`
	CompletedDate string             `json:"completed_date,omitempty" bson:"completed_date,omitempty"`
	CreatedAt     string             `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt     string             `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	UserID        primitive.ObjectID `json:"user_id,omitempty" bson:"user_id,omitempty"`
}
