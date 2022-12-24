package gql_controllers

import (
	"context"
	"errors"
	"golang-nextjs-todo/graph/model"
	"golang-nextjs-todo/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GqlGetAllTasks(ctx context.Context, gql_taskcollection *mongo.Collection, userId primitive.ObjectID) ([]*models.Task, error) {
	var tasks []*models.Task
	query := bson.D{
		bson.E{
			Key:   "user_id",
			Value: userId,
		},
	}
	cursor, err := gql_taskcollection.Find(ctx, query)
	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		var task *models.Task
		if err := cursor.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func GqlGetTask(ctx context.Context, gql_taskcollection *mongo.Collection, id primitive.ObjectID) (*models.Task, error) {
	var task *models.Task
	query := bson.D{
		bson.E{
			Key:   "_id",
			Value: id,
		},
	}
	if err := gql_taskcollection.FindOne(ctx, query).Decode(&task); err != nil {
		return nil, err
	}
	return task, nil
}

func CreateTask(ctx context.Context, gql_taskcollection *mongo.Collection, task *models.Task) (*models.Task, error) {
	result, err := gql_taskcollection.InsertOne(ctx, task)
	if err != nil {
		return nil, err
	}
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, errors.New("failed to fetch task id when creating task")
	}
	task.ID = oid
	return task, nil
}

func DeleteTask(ctx context.Context, gql_taskcollection *mongo.Collection, id primitive.ObjectID) error {
	query := bson.D{
		bson.E{
			Key:   "_id",
			Value: id,
		},
	}
	_, err := gql_taskcollection.DeleteOne(ctx, query)
	return err
}

func UpdateTask(ctx context.Context, gql_taskcollection *mongo.Collection, task model.UpdateTask) error {
	filter := bson.D{
		bson.E{
			Key:   "_id",
			Value: task.ID,
		},
	}

	update := bson.D{
		bson.E{
			Key: "$set",
			Value: bson.D{
				bson.E{
					Key:   "title",
					Value: task.Title,
				},
				bson.E{
					Key:   "completed",
					Value: task.Completed,
				},
				bson.E{
					Key:   "completed_date",
					Value: time.Now(),
				},
			},
		},
	}
	_, err := gql_taskcollection.UpdateOne(ctx, filter, update)
	return err
}
