package test

import (
	"fmt"
	"golang-nextjs-todo/models"
	"log"
	"os"
	"testing"
	"time"

	unitTest "github.com/Valiben/gin_unit_test"
	"github.com/Valiben/gin_unit_test/utils"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
)

func TestGetAllUsers(t *testing.T) {
	type HTTPResponse struct {
		Status  int           `json:"status"`
		Success bool          `json:"success"`
		Message string        `json:"message"`
		Data    []models.User `json:"data"`
	}
	var res HTTPResponse

	if err := unitTest.TestHandlerUnMarshalResp(utils.GET, "/api/v1/users/", "json", nil, &res); err != nil {
		t.Errorf("TestGetAllUsers: %v\n", err)
		return
	}

	if res.Success != true {
		t.Errorf("TestGetUserById: expected success, got %v\n", res.Success)
		return
	}

	if len(res.Data) != 10 {
		t.Errorf("TestGetAllUsers: expected 10 users, got %v\n", len(res.Data))
		return
	}

	t.Log("passed")
}

func TestGetUserById(t *testing.T) {
	type HTTPResponse struct {
		Status  int         `json:"status"`
		Success bool        `json:"success"`
		Message string      `json:"message"`
		Data    models.User `json:"data"`
	}
	var res HTTPResponse

	// get one test user id
	var users []*models.User
	cursor, err := usercollection.Find(ctx, bson.D{{}})
	if err != nil {
		log.Fatal("Error getting sample users: ", err)
		return
	}

	for cursor.Next(ctx) {
		var user *models.User
		err := cursor.Decode(&user)
		if err != nil {
			log.Fatal("Error decoding sample users: ", err)
			return
		}
		users = append(users, user)
	}

	userId := users[0].ID.Hex()
	// set customized request headers
	// Generate jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userId,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		log.Fatal("Error in generating JWT token: ", err)
		return
	}

	unitTest.AddHeader("authorization", fmt.Sprintf("Bearer %s", tokenString))
	if err := unitTest.TestHandlerUnMarshalResp(utils.GET, fmt.Sprintf("/api/v1/users/%v", userId), "json", nil, &res); err != nil {
		t.Errorf("TestGetUserById: %v/n", err)
		return
	}

	if res.Success != true {
		t.Errorf("TestGetUserById: expected success, got %v\n", res.Success)
		return
	}

	if res.Data.ID.Hex() != userId {
		t.Errorf("TestGetUserById: expected user id %v, got %v\n", userId, res.Data.ID.Hex())
		return
	}
	t.Log("passed")
}
