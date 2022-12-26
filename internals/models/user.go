package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName string             `json:"first_name" bson:"first_name,omitempty"`
	LastName  string             `json:"last_name" bson:"last_name,omitempty"`
	Email     string             `json:"email" bson:"email,omitempty"`
	Password  string             `json:"password" bson:"password,omitempty"`
	CreatedAt string             `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt string             `json:"updated_at" bson:"updated_at,omitempty"`
	Token     string             `json:"token" bson:"token,omitempty"`
}
