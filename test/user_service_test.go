package test

import (
	"fmt"
	"golang-nextjs-todo/models"
	util "golang-nextjs-todo/utils"
	"log"
	"testing"

	unitTest "github.com/Valiben/gin_unit_test"
	"github.com/Valiben/gin_unit_test/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func getSampleUsers() []*models.User {
	// get one test user id
	var users []*models.User
	cursor, err := usercollection.Find(ctx, bson.D{{}})
	if err != nil {
		log.Fatal("Error getting sample users: ", err)
		return nil
	}

	for cursor.Next(ctx) {
		var user *models.User
		err := cursor.Decode(&user)
		if err != nil {
			log.Fatal("Error decoding sample users: ", err)
			return nil
		}
		users = append(users, user)
	}
	return users
}

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

	if !res.Success {
		t.Errorf("TestGetUserById: expected success, got %v\n", res.Success)
		return
	}

	if len(res.Data) == 0 {
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
	users := getSampleUsers()
	tokenString, err := util.GenerateJWTToken(users[0].ID)
	if err != nil {
		log.Fatal("Error in generating JWT token: ", err)
		return
	}
	unitTest.AddHeader("authorization", fmt.Sprintf("Bearer %s", tokenString))

	if err := unitTest.TestHandlerUnMarshalResp(utils.GET, fmt.Sprintf("/api/v1/users/%v", users[0].ID.Hex()), "json", nil, &res); err != nil {
		t.Errorf("TestGetUserById: %v/n", err)
		return
	}

	if !res.Success {
		t.Errorf("TestGetUserById: expected success, got %v\n", res.Success)
		return
	}

	if res.Data.ID.Hex() != users[0].ID.Hex() {
		t.Errorf("TestGetUserById: expected user id %v, got %v\n", users[0].ID.Hex(), res.Data.ID.Hex())
		return
	}
	t.Log("passed")
}

func TestLogin(t *testing.T) {
	type HTTPResponse struct {
		Status  int         `json:"status"`
		Success bool        `json:"success"`
		Message string      `json:"message"`
		Data    models.User `json:"data"`
	}

	type Params struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	params := Params{
		Email:    "test_email1@test.com",
		Password: "password_1",
	}

	var res HTTPResponse

	if err := unitTest.TestHandlerUnMarshalResp(utils.POST, "/api/v1/users/login", "json", params, &res); err != nil {
		t.Errorf("TestLogin: %v/n", err)
		return
	}

	if !res.Success {
		t.Errorf("TestLogin: expected success, got %v\n", res.Success)
		return
	}

	t.Log("passed")
}

func TestSignup(t *testing.T) {
	type HTTPResponse struct {
		Status  int         `json:"status"`
		Success bool        `json:"success"`
		Message string      `json:"message"`
		Data    models.User `json:"data"`
	}

	type UserParams struct {
		Email     string `json:"email"`
		Password  string `json:"password"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}

	var res HTTPResponse

	params := UserParams{
		Email:     "test_email11@test.com",
		Password:  "password_11",
		FirstName: "test_first_name_11",
		LastName:  "test_last_name_11",
	}

	if err := unitTest.TestHandlerUnMarshalResp(utils.POST, "/api/v1/users/signup", "json", params, &res); err != nil {
		t.Errorf("TestSignup: %v/n", err)
		return
	}

	if !res.Success {
		t.Errorf("TestSignup: expected success, got %v\n", res.Success)
		return
	}

	if res.Data.Email != params.Email {
		t.Errorf("TestSignup: expected email %v, got %v\n", params.Email, res.Data.Email)
		return
	}

	t.Log("passed")
}

func TestUpdateUser(t *testing.T) {
	type HTTPResponse struct {
		Status  int         `json:"status"`
		Success bool        `json:"success"`
		Message string      `json:"message"`
		Data    models.User `json:"data"`
	}

	type Params struct {
		Email    string `json:"email"`
		LastName string `json:"last_name"`
	}

	users := getSampleUsers()

	var res HTTPResponse

	params := Params{
		Email:    "test_email999@test.com",
		LastName: "last_name_999",
	}

	tokenString, err := util.GenerateJWTToken(users[0].ID)
	if err != nil {
		log.Fatal("Error in generating JWT token: ", err)
		return
	}
	unitTest.AddHeader("authorization", fmt.Sprintf("Bearer %s", tokenString))
	if err := unitTest.TestHandlerUnMarshalResp(utils.PUT, fmt.Sprintf("/api/v1/users/%v", users[0].ID.Hex()), "json", params, &res); err != nil {
		t.Errorf("TestSignup: %v/n", err)
		return
	}

	if !res.Success {
		t.Errorf("TestSignup: expected success, got %v\n", res.Success)
		return
	}

	if res.Data.Email != params.Email {
		t.Errorf("TestSignup: expected email %v, got %v\n", params.Email, res.Data.Email)
		return
	}

	t.Log("passed")
}

func TestDeleteUser(t *testing.T) {
	type HTTPResponse struct {
		Status  int    `json:"status"`
		Success bool   `json:"success"`
		Message string `json:"message"`
		Data    string `json:"data"`
	}

	users := getSampleUsers()
	beforeDeleteCount := len(users)

	var res HTTPResponse

	tokenString, err := util.GenerateJWTToken(users[0].ID)
	if err != nil {
		log.Fatal("Error in generating JWT token: ", err)
		return
	}
	unitTest.AddHeader("authorization", fmt.Sprintf("Bearer %s", tokenString))
	if err := unitTest.TestHandlerUnMarshalResp(utils.DELETE, fmt.Sprintf("/api/v1/users/%v", users[0].ID.Hex()), "json", nil, &res); err != nil {
		t.Errorf("TestSignup: %v/n", err)
		return
	}

	if !res.Success {
		t.Errorf("TestSignup: expected success, got %v\n", res.Success)
		return
	}

	users = getSampleUsers()
	afterDeleteCount := len(users)

	if beforeDeleteCount != afterDeleteCount+1 {
		t.Errorf("TestSignup: expected %v, got %v\n", beforeDeleteCount, afterDeleteCount)
		return
	}
	t.Log("passed")
}
