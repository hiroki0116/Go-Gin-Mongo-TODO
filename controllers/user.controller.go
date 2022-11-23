package controllers

import (
	"context"
	"errors"
	"golang-nextjs-todo/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserController interface {
	CreateUser(*models.User) error
	GetUserById(primitive.ObjectID) (*models.User, error)
	GetAllUsers() ([]*models.User, error)
	UpdateUser(primitive.ObjectID, *models.User) error
	DeleteUser(primitive.ObjectID) error
}

type UserControllerReceiver struct {
	usercollection *mongo.Collection
	ctx            context.Context
}

func NewUserContoller(usercollection *mongo.Collection, ctx context.Context) *UserControllerReceiver {
	return &UserControllerReceiver{
		usercollection: usercollection,
		ctx:            ctx,
	}
}

func (uc *UserControllerReceiver) CreateUser(user *models.User) error {
	// TO DO: 1. Add JWT token
	// TO DO: 2. Has password
	createdAt := time.Now().Format(time.RFC3339)
	updatedAt := time.Now().Format(time.RFC3339)
	user.CreatedAt = createdAt
	user.UpdatedAt = updatedAt
	_, err := uc.usercollection.InsertOne(uc.ctx, user)
	return err
}

func (uc *UserControllerReceiver) GetUserById(id primitive.ObjectID) (*models.User, error) {
	var user *models.User
	query := bson.D{
		bson.E{
			Key:   "_id",
			Value: id,
		},
	}
	err := uc.usercollection.FindOne(uc.ctx, query).Decode(&user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *UserControllerReceiver) GetAllUsers() ([]*models.User, error) {
	var users []*models.User
	cursor, err := uc.usercollection.Find(uc.ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}

	for cursor.Next(uc.ctx) {
		var user *models.User
		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (uc *UserControllerReceiver) UpdateUser(id primitive.ObjectID, user *models.User) error {
	filter := bson.D{
		bson.E{
			Key:   "_id",
			Value: id,
		},
	}

	update := bson.D{
		bson.E{
			Key: "$set",
			Value: bson.D{
				bson.E{
					Key:   "first_name",
					Value: user.FirstName,
				},
				bson.E{
					Key:   "last_name",
					Value: user.LastName,
				},
				bson.E{
					Key:   "email",
					Value: user.Email,
				},
				bson.E{
					Key:   "updated_at",
					Value: time.Now().Format(time.RFC3339),
				},
			},
		},
	}

	if result, _ := uc.usercollection.UpdateOne(uc.ctx, filter, update); result.MatchedCount != 1 {
		return errors.New("failed to update user. User not found")
	}

	return nil
}

func (uc *UserControllerReceiver) DeleteUser(id primitive.ObjectID) error {
	filter := bson.D{
		bson.E{
			Key:   "_id",
			Value: id,
		},
	}

	if result, _ := uc.usercollection.DeleteOne(uc.ctx, filter); result.DeletedCount != 1 {
		return errors.New("no matched document found for user delete")
	}
	return nil
}
