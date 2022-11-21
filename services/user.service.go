package services

import (
	"golang-nextjs-todo/controllers"
	"golang-nextjs-todo/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	UserController controllers.UserController
}

func New(usercontroller controllers.UserController) UserService {
	return UserService{
		UserController: usercontroller,
	}
}

func (us *UserService) GetUserById(ctx *gin.Context) {
	id := ctx.Param("id")
	userId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}

	user, err := us.UserController.GetUserById(userId)
	if err != nil {
		res := utils.NewHttpResponse(http.StatusInternalServerError, err)
		ctx.JSON(http.StatusInternalServerError, res)
	}

	res := utils.NewHttpResponse(http.StatusOK, user)
	ctx.JSON(http.StatusOK, res)
}

func (us *UserService) GetAllUsers(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (us *UserService) CreateUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (us *UserService) UpdateUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (us *UserService) DeleteUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}
