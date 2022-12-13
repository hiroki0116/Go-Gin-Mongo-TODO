package services

import (
	"golang-nextjs-todo/controllers"
	"golang-nextjs-todo/models"
	"golang-nextjs-todo/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskService struct {
	TaskController controllers.TaskController
}

func NewTask(taskcontroller controllers.TaskController) TaskService {
	return TaskService{
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

	userId := utils.FetchUserFromCtx(ctx)

	task, err := ts.TaskController.GetTaskById(taskId, userId)
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
	if err := ts.TaskController.CreateTask(&task, userId); err != nil {
		res := utils.NewHttpResponse(http.StatusBadRequest, err)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.NewHttpResponse(http.StatusCreated, "Successfully created task")
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

	if err := ts.TaskController.UpdateTask(taskId, &task); err != nil {
		res := utils.NewHttpResponse(http.StatusBadRequest, err)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.NewHttpResponse(http.StatusOK, "Successfully updated task")
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
