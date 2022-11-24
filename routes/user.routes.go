package routes

import (
	"golang-nextjs-todo/services"

	"github.com/gin-gonic/gin"
)

type UserRoutes struct {
	UserService services.UserService
}

func NewUserRoute(userservice services.UserService) UserRoutes {
	return UserRoutes{
		UserService: userservice,
	}
}

func (ur *UserRoutes) UserRoutes(rg *gin.RouterGroup) {
	userroute := rg.Group("/users")
	userroute.POST("/", ur.UserService.CreateUser)
	userroute.GET("/:id", ur.UserService.GetUserById)
	userroute.GET("/", ur.UserService.GetAllUsers)
	userroute.PATCH("/:id", ur.UserService.UpdateUser)
	userroute.DELETE("/:id", ur.UserService.DeleteUser)
}
