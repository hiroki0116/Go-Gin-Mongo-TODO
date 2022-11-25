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
	userroute.POST("/signup", ur.UserService.SignUp)
	userroute.POST("/login", ur.UserService.Login)
	userroute.GET("/:id", ur.UserService.GetUserById)
	userroute.GET("/", ur.UserService.GetAllUsers)
	userroute.PATCH("/:id", ur.UserService.UpdateUser)
	userroute.DELETE("/:id", ur.UserService.DeleteUser)
}
