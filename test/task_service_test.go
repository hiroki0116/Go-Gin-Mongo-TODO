package test

import (
	"fmt"
	"golang-nextjs-todo/internals/models"
	util "golang-nextjs-todo/internals/utils"
	"log"
	"testing"

	unitTest "github.com/Valiben/gin_unit_test"
	"github.com/Valiben/gin_unit_test/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func getSampleTasks() []*models.Task {
	// get one test user id
	var tasks []*models.Task
	cursor, err := taskcollection.Find(ctx, bson.D{{}})
	if err != nil {
		log.Fatal("Error getting sample tasks: ", err)
		return nil
	}

	for cursor.Next(ctx) {
		var task *models.Task
		err := cursor.Decode(&task)
		if err != nil {
			log.Fatal("Error decoding sample tasks: ", err)
			return nil
		}
		tasks = append(tasks, task)
	}
	return tasks
}

func setHeaderToken(userId primitive.ObjectID) error {
	tokenString, err := util.GenerateJWTToken(userId)
	if err != nil {
		return err
	}
	unitTest.AddHeader("authorization", fmt.Sprintf("Bearer %s", tokenString))
	return nil
}

func TestGeAllTasks(t *testing.T) {
	type HTTPResponse struct {
		Status  int           `json:"status"`
		Success bool          `json:"success"`
		Message string        `json:"message"`
		Data    []models.Task `json:"data"`
	}
	var res HTTPResponse

	users := getSampleTasks()

	if error := setHeaderToken(users[0].UserID); error != nil {
		t.Errorf("TestCreateTask: %v\n", error)
		return
	}
	if err := unitTest.TestHandlerUnMarshalResp(utils.GET, "/api/v1/tasks/", "json", nil, &res); err != nil {
		t.Errorf("TestGetAllTasks: %v\n", err)
		return
	}

	if !res.Success {
		t.Errorf("TestGetAllTasks: expected success, got %v\n", res.Success)
		return
	}

	if len(res.Data) == 0 {
		t.Errorf("TestGetAllTasks: expected 10 tasks, got %v\n", len(res.Data))
		return
	}

	t.Log("passed")

}

func TestCreateTask(t *testing.T) {
	type HTTPResponse struct {
		Status  int         `json:"status"`
		Success bool        `json:"success"`
		Message string      `json:"message"`
		Data    models.Task `json:"data"`
	}
	var res HTTPResponse

	type TaskParams struct {
		Title     string             `json:"title"`
		UserID    primitive.ObjectID `json:"user_id"`
		UpdatedAt string             `json:"updated_at"`
		CreatedAt string             `json:"created_at"`
	}

	users := getSampleTasks()
	if error := setHeaderToken(users[0].UserID); error != nil {
		t.Errorf("TestCreateTask: %v\n", error)
		return
	}

	params := TaskParams{
		Title:  "Created Test Task",
		UserID: users[0].UserID,
		// UpdatedAt: time.Now().Format(time.RFC3339),
		// CreatedAt: time.Now().Format(time.RFC3339),
	}

	if err := unitTest.TestHandlerUnMarshalResp(utils.POST, "/api/v1/tasks/", "json", params, &res); err != nil {
		t.Errorf("TestCreateTask: %v\n", err)
		return
	}

	if !res.Success {
		t.Errorf("TestCreateTask: expected success, got %v\n", res.Success)
		return
	}

	t.Log("passed")
}

func TestGetTaskById(t *testing.T) {
	type HTTPResponse struct {
		Status  int         `json:"status"`
		Success bool        `json:"success"`
		Message string      `json:"message"`
		Data    models.Task `json:"data"`
	}
	var res HTTPResponse
	// get one test user id
	tasks := getSampleTasks()
	tokenString, err := util.GenerateJWTToken(tasks[0].UserID)
	if err != nil {
		log.Fatal("Error in generating JWT token: ", err)
		return
	}
	unitTest.AddHeader("authorization", fmt.Sprintf("Bearer %s", tokenString))

	params := tasks[0].ID

	if err := unitTest.TestHandlerUnMarshalResp(utils.GET, fmt.Sprintf("/api/v1/tasks/%v", tasks[0].ID.Hex()), "json", params, &res); err != nil {
		t.Errorf("TestGetTaskById: %v/n", err)
		return
	}

	if !res.Success {
		t.Errorf("TestGetTaskById: expected success, got %v\n", res.Success)
		return
	}

	if res.Data.ID.Hex() != tasks[0].ID.Hex() {
		t.Errorf("TestGetTaskById: expected user id %v, got %v\n", tasks[0].ID.Hex(), res.Data.ID.Hex())
		return
	}
	t.Log("passed")
}

func TestUpdateTask(t *testing.T) {
	type HTTPResponse struct {
		Status  int         `json:"status"`
		Success bool        `json:"success"`
		Message string      `json:"message"`
		Data    models.Task `json:"data"`
	}

	type TaskParams struct {
		Title     string `json:"title"`
		Completed bool   `json:"completed,omitempty"`
	}

	tasks := getSampleTasks()

	var res HTTPResponse

	params := TaskParams{
		Title:     "Updated Test Task",
		Completed: true,
	}

	tokenString, err := util.GenerateJWTToken(tasks[0].UserID)
	if err != nil {
		log.Fatal("Error in generating JWT token: ", err)
		return
	}
	unitTest.AddHeader("authorization", fmt.Sprintf("Bearer %s", tokenString))
	if err := unitTest.TestHandlerUnMarshalResp(utils.PUT, fmt.Sprintf("/api/v1/tasks/%v", tasks[0].ID.Hex()), "json", params, &res); err != nil {
		t.Errorf("TestUpdateTask: %v/n", err)
		return
	}

	if !res.Success {
		t.Errorf("TestUpdateTask: expected success, got %v\n", res.Success)
		return
	}

	updatedTasks := getSampleTasks()
	if updatedTasks[0].Title != params.Title {
		t.Errorf("TestUpdateTask: expected title %v, got %v\n", params.Title, updatedTasks[0].Title)
		return
	}

	if updatedTasks[0].Completed != params.Completed {
		t.Errorf("TestUpdateTask: expected completed %v, got %v\n", params.Completed, updatedTasks[0].Completed)
		return
	}

	t.Log("passed")
}

func TestDeleteTask(t *testing.T) {
	type HTTPResponse struct {
		Status  int    `json:"status"`
		Success bool   `json:"success"`
		Message string `json:"message"`
		Data    string `json:"data"`
	}

	tasks := getSampleTasks()

	var res HTTPResponse
	tokenString, err := util.GenerateJWTToken(tasks[0].UserID)

	if err != nil {
		log.Fatal("Error in generating JWT token: ", err)
		return
	}
	unitTest.AddHeader("authorization", fmt.Sprintf("Bearer %s", tokenString))
	if err := unitTest.TestHandlerUnMarshalResp(utils.DELETE, fmt.Sprintf("/api/v1/tasks/%v", tasks[0].ID.Hex()), "json", nil, &res); err != nil {
		t.Errorf("TestUpdateTask: %v/n", err)
		return
	}

	if !res.Success {
		t.Errorf("TestUpdateTask: expected success, got %v\n", res.Success)
		return
	}

	afterDeletedTasks := getSampleTasks()
	if len(tasks)-len(afterDeletedTasks) != 1 {
		t.Errorf("TestUpdateTask: expected deleted task, got %v\n", len(tasks)-len(afterDeletedTasks))
		return
	}

	t.Log("passed")
}
