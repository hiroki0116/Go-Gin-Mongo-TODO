package test

import (
	"context"
	"fmt"
	"golang-nextjs-todo/models"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func PopulateUserSampleData(mongo *mongo.Collection, ctx context.Context) error {
	// create 10 users data
	for i := 0; i < 10; i++ {

		// hash password before saving to db
		beforeHash := fmt.Sprintf("password_%d", i)
		password, err := bcrypt.GenerateFromPassword([]byte(beforeHash), 10)
		if err != nil {
			log.Fatal("Error hashing password: ", err)
			return err
		}

		user := models.User{
			FirstName: fmt.Sprintf("first_name_%d", i),
			LastName:  fmt.Sprintf("last_name_%d", i),
			Email:     fmt.Sprintf("test_email%v@test.com", i),
			Password:  string(password),
		}

		if _, err = mongo.InsertOne(ctx, user); err != nil {
			log.Fatal("Error inserting new user: ", err)
			return err
		}
	}

	return nil
}

func DeleteUserData(mongo *mongo.Collection, ctx context.Context) error {
	filter := bson.D{{}}
	_, err := mongo.DeleteMany(ctx, filter)
	if err != nil {
		log.Fatal("Error deleting sample user data: ", err)
		return err
	}
	return nil
}
