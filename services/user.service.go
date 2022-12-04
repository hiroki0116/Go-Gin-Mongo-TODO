package services

import (
	"golang-nextjs-todo/controllers"
	"golang-nextjs-todo/models"
	"golang-nextjs-todo/utils"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
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
		res := utils.NewHttpResponse(http.StatusBadRequest, err)
		ctx.JSON(http.StatusBadRequest, res)
		return
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

func (us *UserService) SignUp(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		res := utils.NewHttpResponse(http.StatusBadRequest, err)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	// Hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		res := utils.NewHttpResponse(http.StatusInternalServerError, err)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}
	user.Password = string(hash)

	// check existing user
	if _, err := us.UserController.GetUserByEmail(user.Email); err == nil {
		res := utils.NewHttpResponse(http.StatusBadRequest, "User already exists")
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	userId, err := us.UserController.CreateUser(&user)
	if err != nil {
		res := utils.NewHttpResponse(http.StatusBadRequest, err)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	// Generate jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userId,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		res := utils.NewHttpResponse(http.StatusBadRequest, err)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	// Allow most cross-domain cookie-sharing
	ctx.SetSameSite(http.SameSiteNoneMode)
	// Set it in cookie
	ctx.SetCookie("token", tokenString, 3600*24*30, "/", "*", false, false)
	// Return response
	resUser := models.User{
		ID:        userId,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Token:     tokenString,
	}

	// Return response
	res := utils.NewHttpResponse(http.StatusCreated, resUser)
	ctx.JSON(http.StatusCreated, res)
}

func (us *UserService) Login(ctx *gin.Context) {

	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		res := utils.NewHttpResponse(http.StatusBadRequest, err)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	// Find user by email
	user, err := us.UserController.GetUserByEmail(body.Email)
	if err != nil {
		res := utils.NewHttpResponse(http.StatusBadRequest, err)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	// Compare password and hash
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		res := utils.NewHttpResponse(http.StatusBadRequest, err)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	// Generate jwt token
	tokenString, err := utils.GenerateJWTToken(user.ID)
	if err != nil {
		res := utils.NewHttpResponse(http.StatusBadRequest, err)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	ctx.SetSameSite(http.SameSiteLaxMode)
	// Set it in cookie
	ctx.SetCookie("token", tokenString, 3600*24*30, "/", "*", false, false)
	// Return response
	resUser := models.User{
		ID:        user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Token:     tokenString,
	}
	res := utils.NewHttpResponse(http.StatusOK, resUser)
	ctx.JSON(http.StatusOK, res)
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
