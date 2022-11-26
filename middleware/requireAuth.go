package middleware

import (
	"fmt"
	"golang-nextjs-todo/controllers"
	"golang-nextjs-todo/utils"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RequireAuth struct {
	UserController controllers.UserController
}

func NewRequireAuth(usercontroller controllers.UserController) RequireAuth {
	return RequireAuth{
		UserController: usercontroller,
	}
}

func (ra *RequireAuth) SetJWT(ctx *gin.Context) {
	tokenString, err := ctx.Cookie("Authorization")
	if err != nil {
		res := utils.NewHttpResponse(http.StatusUnauthorized, err)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return
	}

	// Decode and validate the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check token expiration
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			res := utils.NewHttpResponse(http.StatusUnauthorized, "Token expired")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
			return
		}

		userId, err := primitive.ObjectIDFromHex(claims["sub"].(string))
		if err != nil {
			res := utils.NewHttpResponse(http.StatusInternalServerError, err)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
			return
		}
		user, err := ra.UserController.GetUserById(userId)
		if err != nil {
			res := utils.NewHttpResponse(http.StatusUnauthorized, err)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
			return
		}

		ctx.Set("id", user.ID)
		ctx.Set("email", user.Email)

		ctx.Next()
	} else {
		res := utils.NewHttpResponse(http.StatusUnauthorized, err)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return
	}
}
