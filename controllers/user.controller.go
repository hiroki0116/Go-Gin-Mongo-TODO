package controllers

import (
	"context"
	"golang-nextjs-todo/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserController interface {
	CreateUser(*models.User) error
	GetUserById(primitive.ObjectID) (*models.User, error)
	GetAllUsers() ([]*models.User, error)
	UpdateUser(primitive.ObjectID) error
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
	return nil
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
	return nil, nil
}

func (uc *UserControllerReceiver) UpdateUser(id primitive.ObjectID) error {
	return nil
}

func (uc *UserControllerReceiver) DeleteUser(id primitive.ObjectID) error {
	return nil
}
