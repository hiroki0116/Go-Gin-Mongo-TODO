package test

import (
	"golang-nextjs-todo/models"
	"testing"

	unitTest "github.com/Valiben/gin_unit_test"
	"github.com/Valiben/gin_unit_test/utils"
)

func TestGetAllUsers(t *testing.T) {
	type HTTPResponse struct {
		Status  int           `json:"status"`
		Success bool          `json:"success"`
		Message string        `json:"message"`
		Data    []models.User `json:"data"`
	}
	var res HTTPResponse

	err := unitTest.TestHandlerUnMarshalResp(
		utils.GET, "/api/v1/users/", "json", nil, &res,
	)

	if err != nil {
		t.Errorf("TestGetAllUsers: %v\n", err)
		return
	}
	t.Log("passed")
}
