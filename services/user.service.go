package services

import (
	"golang-nextjs-todo/controllers"
	"golang-nextjs-todo/models"
	"golang-nextjs-todo/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	UserController controllers.UserController
}

func NewUser(usercontroller controllers.UserController) UserService {
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
		res := utils.NewHttpResponse(http.StatusBadRequest, err)
		ctx.JSON(http.StatusBadRequest, res)
	}

	res := utils.NewHttpResponse(http.StatusOK, user)
	ctx.JSON(http.StatusOK, res)
}

func (us *UserService) GetAllUsers(ctx *gin.Context) {
	users, err := us.UserController.GetAllUsers()
	if err != nil {
		res := utils.NewHttpResponse(http.StatusBadRequest, err)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.NewHttpResponse(http.StatusOK, users)
	ctx.JSON(http.StatusOK, res)
}

func (us *UserService) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		res := utils.NewHttpResponse(http.StatusBadRequest, err)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	err := us.UserController.CreateUser(&user)
	if err != nil {
		res := utils.NewHttpResponse(http.StatusBadRequest, err)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.NewHttpResponse(http.StatusCreated, user)
	ctx.JSON(http.StatusCreated, res)
}

func (us *UserService) UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")
	var user models.User
	userId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		res := utils.NewHttpResponse(http.StatusBadRequest, err)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if err := ctx.ShouldBindJSON(&user); err != nil {
		res := utils.NewHttpResponse(http.StatusBadRequest, err)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if err := us.UserController.UpdateUser(userId, &user); err != nil {
		res := utils.NewHttpResponse(http.StatusBadRequest, err)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.NewHttpResponse(http.StatusCreated, user)
	ctx.JSON(http.StatusCreated, res)
}

func (us *UserService) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	userId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		res := utils.NewHttpResponse(http.StatusBadRequest, err)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if err := us.UserController.DeleteUser(userId); err != nil {
		res := utils.NewHttpResponse(http.StatusBadRequest, err)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.NewHttpResponse(http.StatusOK, "Successfully deleted")
	ctx.JSON(http.StatusCreated, res)
}
