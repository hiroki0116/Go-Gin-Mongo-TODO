package routes

import (
	"golang-nextjs-todo/middleware"
	"golang-nextjs-todo/services"

	"github.com/gin-gonic/gin"
)

type TaskRoutes struct {
	TaskService services.ITaskService
	RequireAuth middleware.RequireAuth
}

func NewTaskRoute(taskservice services.ITaskService, requireauth middleware.RequireAuth) TaskRoutes {
	return TaskRoutes{
		TaskService: taskservice,
		RequireAuth: requireauth,
	}
}

func (tr *TaskRoutes) TaskRoutes(rg *gin.RouterGroup) {
	userroute := rg.Group("/tasks")
	userroute.POST("/", tr.RequireAuth.SetJWT, tr.TaskService.CreateTask)
	userroute.GET("/:id", tr.RequireAuth.SetJWT, tr.TaskService.GetTaskById)
	userroute.GET("/", tr.RequireAuth.SetJWT, tr.TaskService.GetAllTasks)
	userroute.PUT("/:id", tr.RequireAuth.SetJWT, tr.TaskService.UpdateTask)
	userroute.DELETE("/:id", tr.RequireAuth.SetJWT, tr.TaskService.DeleteTask)
}
