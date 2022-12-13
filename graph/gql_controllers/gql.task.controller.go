package gql_controllers

import (
	"context"
	"golang-nextjs-todo/graph/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GqlGetAllTasks(ctx context.Context, gql_taskcollection *mongo.Collection, userId primitive.ObjectID) ([]*model.Task, error) {
	var tasks []*model.Task
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
		var task *model.Task
		if err := cursor.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func GqlGetTask(ctx context.Context, gql_taskcollection *mongo.Collection, id primitive.ObjectID) (*model.Task, error) {
	var task *model.Task
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

func CreateTask(ctx context.Context, gql_taskcollection *mongo.Collection, task *model.Task) (*model.Task, error) {
	_, err := gql_taskcollection.InsertOne(ctx, task)
	if err != nil {
		return nil, err
	}
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
	oid, err := primitive.ObjectIDFromHex(task.ID)
	if err != nil {
		return err
	}
	filter := bson.D{
		bson.E{
			Key:   "_id",
			Value: oid,
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
	_, err = gql_taskcollection.UpdateOne(ctx, filter, update)
	return err
}