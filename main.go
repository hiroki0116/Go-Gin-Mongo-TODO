package main

import (
	"context"
	"fmt"
	"golang-nextjs-todo/controllers"
	"golang-nextjs-todo/middleware"
	"golang-nextjs-todo/routes"
	"golang-nextjs-todo/services"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server         *gin.Engine
	mongoclient    *mongo.Client
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
	err            error
)

func init() {
	ctx = context.TODO()
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	// database connection
	mongoconn := options.Client().ApplyURI(string(os.Getenv("MONGO_URI")))
	mongoclient, err = mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Connected to MongoDB!!!!")
	// collections
	usercollection = mongoclient.Database("golangTodos").Collection("users")
	taskcollection = mongoclient.Database("golangTodos").Collection("tasks")
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
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST, GET, OPTIONS, PUT, DELETE"},
		AllowHeaders:     []string{"Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, X-Max,Set-Cookie"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
}

func main() {
	defer mongoclient.Disconnect(ctx)
	basepath := server.Group("/api/v1")

	userroute.UserRoutes(basepath)
	taskroute.TaskRoutes(basepath)

	log.Fatalln(server.Run(":8000"))
}
