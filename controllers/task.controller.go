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
	GetTaskById(primitive.ObjectID) (*models.Task, error)
	GetAllTasks() ([]*models.Task, error)
	CreateTask(*models.Task) error
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

func (tc *TaskControllerReceiver) GetTaskById(id primitive.ObjectID) (*models.Task, error) {
	var task *models.Task
	query := bson.D{
		bson.E{
			Key:   "_id",
			Value: id,
		},
	}
	if err := tc.taskcollection.FindOne(tc.ctx, query).Decode(&task); err != nil {
		return nil, err
	}

	return task, nil
}

func (tc *TaskControllerReceiver) GetAllTasks() ([]*models.Task, error) {
	var tasks []*models.Task
	query := bson.D{}
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

func (tc *TaskControllerReceiver) CreateTask(task *models.Task) error {
	createdAt := time.Now().Format(time.RFC3339)
	updatedAt := time.Now().Format(time.RFC3339)
	task.Completed = false
	task.CreatedAt = createdAt
	task.UpdatedAt = updatedAt
	_, err := tc.taskcollection.InsertOne(tc.ctx, task)
	return err
}

func (tc *TaskControllerReceiver) UpdateTask(id primitive.ObjectID, task *models.Task) error {
	updatedAt := time.Now().Format(time.RFC3339)
	task.UpdatedAt = updatedAt
	query := bson.D{
		bson.E{
			Key:   "_id",
			Value: id,
		},
	}
	update := bson.D{
		bson.E{
			Key:   "$set",
			Value: task,
		},
	}
	_, err := tc.taskcollection.UpdateOne(tc.ctx, query, update)
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
