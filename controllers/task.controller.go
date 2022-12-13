package controllers

import (
	"context"
	"golang-nextjs-todo/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskController interface {
	GetTaskById(primitive.ObjectID, primitive.ObjectID) (*models.Task, error)
	GetAllTasks(primitive.ObjectID) ([]*models.Task, error)
	CreateTask(*models.Task, primitive.ObjectID) error
	UpdateTask(primitive.ObjectID, *models.Task) error
	DeleteTask(primitive.ObjectID) error
}

type TaskControllerReceiver struct {
	taskcollection *mongo.Collection
	ctx            context.Context
}

func NewTaskController(taskcollection *mongo.Collection, ctx context.Context) *TaskControllerReceiver {
	return &TaskControllerReceiver{
		taskcollection: taskcollection,
		ctx:            ctx,
	}
}

func (tc *TaskControllerReceiver) GetTaskById(id, userId primitive.ObjectID) (*models.Task, error) {
	var task *models.Task
	query := bson.D{
		bson.E{
			Key:   "_id",
			Value: id,
		},
		bson.E{
			Key:   "user_id",
			Value: userId,
		},
	}
	if err := tc.taskcollection.FindOne(tc.ctx, query).Decode(&task); err != nil {
		return nil, err
	}

	return task, nil
}

func (tc *TaskControllerReceiver) GetAllTasks(userId primitive.ObjectID) ([]*models.Task, error) {
	var tasks []*models.Task
	query := bson.D{
		bson.E{
			Key:   "user_id",
			Value: userId,
		},
	}
	cursor, err := tc.taskcollection.Find(tc.ctx, query)
	if err != nil {
		return nil, err
	}
	for cursor.Next(tc.ctx) {
		var task *models.Task
		if err := cursor.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (tc *TaskControllerReceiver) CreateTask(task *models.Task, userId primitive.ObjectID) error {
	createdAt := time.Now().Format(time.RFC3339)
	updatedAt := time.Now().Format(time.RFC3339)
	task.Completed = false
	task.CreatedAt = createdAt
	task.UpdatedAt = updatedAt
	task.UserID = userId
	_, err := tc.taskcollection.InsertOne(tc.ctx, task)
	return err
}

func (tc *TaskControllerReceiver) UpdateTask(id primitive.ObjectID, task *models.Task) error {
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
					Key:   "title",
					Value: task.Title,
				},
				bson.E{
					Key:   "completed",
					Value: task.Completed,
				},
				bson.E{
					Key:   "completed_date",
					Value: time.Now().Format(time.RFC3339),
				},
			},
		},
	}
	_, err := tc.taskcollection.UpdateOne(tc.ctx, filter, update)
	return err
}

func (tc *TaskControllerReceiver) DeleteTask(id primitive.ObjectID) error {
	query := bson.D{
		bson.E{
			Key:   "_id",
			Value: id,
		},
	}
	_, err := tc.taskcollection.DeleteOne(tc.ctx, query)
	return err
}
