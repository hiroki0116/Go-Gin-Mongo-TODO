package main

import (
	"context"
	"golang-nextjs-todo/controllers"
	"golang-nextjs-todo/db"
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
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	// collections
	db.ConnectDB(os.Getenv("MONGO_URI"))
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
	server.GET("/api/v1/healthcheck", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "ok"})
	})
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	server.Use(gin.Logger())
}

func main() {
	defer db.MongoDB.Disconnect(ctx)
	basepath := server.Group("/api/v1")

	userroute.UserRoutes(basepath)
	taskroute.TaskRoutes(basepath)
	log.Fatalln(server.Run(":8000"))
}
