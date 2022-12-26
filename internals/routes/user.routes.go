package routes

import (
	"golang-nextjs-todo/internals/middleware"
	"golang-nextjs-todo/internals/services"

	"github.com/gin-gonic/gin"
)

type UserRoutes struct {
	UserService services.IUserService
	RequireAuth middleware.RequireAuth
}

func NewUserRoute(userservice services.IUserService, requireauth middleware.RequireAuth) UserRoutes {
	return UserRoutes{
		UserService: userservice,
		RequireAuth: requireauth,
	}
}

func (ur *UserRoutes) UserRoutes(rg *gin.RouterGroup) {
	userroute := rg.Group("/users")
	userroute.POST("/signup", ur.UserService.SignUp)
	userroute.POST("/login", ur.UserService.Login)
	userroute.GET("/:id", ur.RequireAuth.SetJWT, ur.UserService.GetUserById)
	userroute.GET("/", ur.UserService.GetAllUsers)
	userroute.PUT("/:id", ur.RequireAuth.SetJWT, ur.UserService.UpdateUser)
	userroute.DELETE("/:id", ur.RequireAuth.SetJWT, ur.UserService.DeleteUser)
}
