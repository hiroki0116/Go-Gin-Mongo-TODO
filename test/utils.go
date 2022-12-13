package test

import (
	"context"
	"fmt"
	"golang-nextjs-todo/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func PopulateTaskSampleData(usercollection *mongo.Collection, taskcollection *mongo.Collection, ctx context.Context) error {

	// create one user sample data
	beforeHash := "password_100"
	password, err := bcrypt.GenerateFromPassword([]byte(beforeHash), 10)
	if err != nil {
		log.Fatal("Error hashing password: ", err)
		return err
	}

	user := models.User{
		FirstName: "first_name_100",
		LastName:  "last_name_100",
		Email:     "test_email100@test.com",
		Password:  string(password),
	}

	result, err := usercollection.InsertOne(ctx, user)
	if err != nil {
		log.Fatal("Error inserting new sample user: ", err)
		return err
	}
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		log.Fatal("Error fetching user id: ", err)
		return err
	}

	// create 10 tasks data
	for i := 0; i < 10; i++ {
		task := models.Task{
			Title:     fmt.Sprintf("title_%d", i),
			Completed: false,
			CreatedAt: time.Now().Format(time.RFC3339),
			UserID:    oid,
		}

		if _, err := taskcollection.InsertOne(ctx, task); err != nil {
			log.Fatal("Error inserting new task: ", err)
			return err
		}
	}

	return nil
}

func DeleteSampleData(mongo *mongo.Collection, ctx context.Context) error {
	filter := bson.D{{}}
	_, err := mongo.DeleteMany(ctx, filter)
	if err != nil {
		log.Fatal("Error deleting sample data: ", err)
		return err
	}
	return nil
}
