package utils

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FetchUserFromCtx(ctx *gin.Context) primitive.ObjectID {
	uId, _ := ctx.Get("id")
	userId, err := primitive.ObjectIDFromHex(uId.(primitive.ObjectID).Hex())
	if err != nil {
		res := NewHttpResponse(http.StatusBadRequest, err)
		ctx.JSON(http.StatusBadRequest, res)
		return primitive.NilObjectID
	}
	return userId
}

func GenerateJWTToken(id primitive.ObjectID) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": id,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
