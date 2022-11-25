package main

import (
	"context"
	"fmt"
	"golang-nextjs-todo/controllers"
	"golang-nextjs-todo/routes"
	"golang-nextjs-todo/services"
	"log"

	"github.com/gin-gonic/gin"
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
	userroute      routes.UserRoutes
	taskroute      routes.TaskRoutes
	ctx            context.Context
	err            error
)

func init() {
	ctx = context.TODO()
	// database connection
	uri := "mongodb+srv://hirokiseino0116:Everythingis6@cluster0.e3ylkdo.mongodb.net/?retryWrites=true&w=majority"
	mongoconn := options.Client().ApplyURI(uri)
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
	// routes
	userroute = routes.NewUserRoute(userservice)
	taskroute = routes.NewTaskRoute(taskservice)
	// server
	server = gin.Default()
}

func main() {
	defer mongoclient.Disconnect(ctx)
	basepath := server.Group("/api/v1")

	userroute.UserRoutes(basepath)
	taskroute.TaskRoutes(basepath)

	log.Fatalln(server.Run(":8000"))
}
