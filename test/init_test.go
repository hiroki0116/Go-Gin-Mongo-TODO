package test

import (
	"context"
	"golang-nextjs-todo/controllers"
	"golang-nextjs-todo/db"
	"golang-nextjs-todo/middleware"
	"golang-nextjs-todo/routes"
	"golang-nextjs-todo/services"
	"log"
	"os"
	"testing"

	unitTest "github.com/Valiben/gin_unit_test"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	server         *gin.Engine
	usercollection *mongo.Collection
	taskcollection *mongo.Collection
	usercontroller controllers.UserController
	taskcontroller controllers.TaskController
	userservice    services.UserService
	taskservice    services.TaskService
	requireauth    middleware.RequireAuth
	userroute      routes.UserRoutes
	taskroute      routes.TaskRoutes
	ctx            context.Context
)

func init() {
	ctx = context.TODO()
	err := godotenv.Load("test.env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	// database connection
	db.ConnectDB("MONGO_URI_TEST")
	// collections
	usercollection = db.MongoDB.Database("golangTodos").Collection("users")
	taskcollection = db.MongoDB.Database("golangTodos").Collection("tasks")
	// controllers
	usercontroller = controllers.NewUserController(usercollection, ctx)
	taskcontroller = controllers.NewTaskController(taskcollection, ctx)
	// services layer
	userservice = services.NewUser(usercontroller)
	taskservice = services.NewTask(taskcontroller)
	// middelware
	requireauth = middleware.NewRequireAuth(usercontroller)
	// routes
	userroute = routes.NewUserRoute(userservice, requireauth)
	taskroute = routes.NewTaskRoute(taskservice, requireauth)
	// server
	server = gin.Default()
}

func TestMain(m *testing.M) {
	defer db.MongoDB.Disconnect(ctx)
	// Set endpointes
	basepath := server.Group("/api/v1")
	userroute.UserRoutes(basepath)
	taskroute.TaskRoutes(basepath)
	unitTest.SetRouter(server)

	// Populate sample data
	PopulateUserSampleData(usercollection, ctx)

	newLog := log.New(os.Stdout, "", log.Llongfile|log.Ldate|log.Ltime)
	unitTest.SetLog(newLog)
	exitVal := m.Run()

	// Delte populated sample data
	DeleteUserData(usercollection, ctx)

	log.Println("Everything below run after ALL test")
	os.Exit(exitVal)
}
