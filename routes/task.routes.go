package routes

import (
	"golang-nextjs-todo/services"

	"github.com/gin-gonic/gin"
)

type TaskRoutes struct {
	TaskService services.TaskService
}

func NewTaskRoute(taskservice services.TaskService) TaskRoutes {
	return TaskRoutes{
		TaskService: taskservice,
	}
}

func (tr *TaskRoutes) TaskRoutes(rg *gin.RouterGroup) {
	userroute := rg.Group("/tasks")
	userroute.POST("/", tr.TaskService.CreateTask)
	userroute.GET("/:id", tr.TaskService.GetTaskById)
	userroute.GET("/", tr.TaskService.GetAllTasks)
	userroute.PATCH("/:id", tr.TaskService.UpdateTask)
	userroute.DELETE("/:id", tr.TaskService.DeleteTask)
}
