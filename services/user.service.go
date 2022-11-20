package services

import (
	"golang-nextjs-todo/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserService struct {
	UserController controllers.UserController
}

func New(usercontroller controllers.UserController) UserService {
	return UserService{
		UserController: usercontroller,
	}
}

func (us *UserService) CreateUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (us *UserService) GetUserById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (us *UserService) GetAllUsers(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (us *UserService) UpdateUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (us *UserService) DeleteUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}
