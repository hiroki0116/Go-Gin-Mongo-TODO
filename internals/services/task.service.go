package services

import (
	"golang-nextjs-todo/internals/controllers"
	"golang-nextjs-todo/internals/models"
	"golang-nextjs-todo/internals/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ITaskService interface {
	GetTaskById(ctx *gin.Context)
	GetAllTasks(ctx *gin.Context)
	CreateTask(ctx *gin.Context)
	UpdateTask(ctx *gin.Context)
	DeleteTask(ctx *gin.Context)
}

type TaskService struct {
	TaskController controllers.ITaskController
}

func NewTask(taskcontroller controllers.ITaskController) ITaskService {
	return &TaskService{
		TaskController: taskcontroller,
	}
}

func (ts *TaskService) GetTaskById(ctx *gin.Context) {
	id := ctx.Param("id")
	taskId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		res := utils.NewHttpResponse(http.StatusBadRequest, err)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	task, err := ts.TaskController.GetTaskById(taskId)
	if err != nil {
		res := utils.NewHttpResponse(http.StatusBadRequest, err)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.NewHttpResponse(http.StatusOK, task)
	ctx.JSON(http.StatusOK, res)
}

func (ts *TaskService) GetAllTasks(ctx *gin.Context) {
	userId := utils.FetchUserFromCtx(ctx)
	tasks, err := ts.TaskController.GetAllTasks(userId)
	if err != nil {
		res := utils.NewHttpResponse(http.StatusBadRequest, err)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	if tasks == nil {
		// return empty array if no tasks found
		var emptyTasks []models.Task
		res := utils.NewHttpResponse(http.StatusOK, emptyTasks)
		ctx.JSON(http.StatusOK, res)
		return
	}
	res := utils.NewHttpResponse(http.StatusOK, tasks)
	ctx.JSON(http.StatusOK, res)
}

func (ts *TaskService) CreateTask(ctx *gin.Context) {
	userId := utils.FetchUserFromCtx(ctx)
	var task models.Task
	if err := ctx.ShouldBindJSON(&task); err != nil {
		res := utils.NewHttpResponse(http.StatusBadRequest, err)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	newTask, err := ts.TaskController.CreateTask(&task, userId)
	if err != nil {
		res := utils.NewHttpResponse(http.StatusBadRequest, err)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.NewHttpResponse(http.StatusCreated, newTask)
	ctx.JSON(http.StatusCreated, res)
}

func (ts *TaskService) UpdateTask(ctx *gin.Context) {
	id := ctx.Param("id")
	taskId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		res := utils.NewHttpResponse(http.StatusBadRequest, err)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	var task models.Task
	if err := ctx.ShouldBindJSON(&task); err != nil {
		res := utils.NewHttpResponse(http.StatusBadRequest, err)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	updatedTask, err := ts.TaskController.UpdateTask(taskId, &task)
	if err != nil {
		res := utils.NewHttpResponse(http.StatusBadRequest, err)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.NewHttpResponse(http.StatusOK, updatedTask)
	ctx.JSON(http.StatusOK, res)
}

func (ts *TaskService) DeleteTask(ctx *gin.Context) {
	id := ctx.Param("id")
	taskId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		res := utils.NewHttpResponse(http.StatusBadRequest, err)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if err := ts.TaskController.DeleteTask(taskId); err != nil {
		res := utils.NewHttpResponse(http.StatusBadRequest, err)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.NewHttpResponse(http.StatusOK, "Successfully deleted task")
	ctx.JSON(http.StatusOK, res)
}
