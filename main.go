package main

import (
	"context"
	graph "golang-nextjs-todo/graph/generated"
	graph1 "golang-nextjs-todo/graph/resolver"
	"golang-nextjs-todo/internals/controllers"
	"golang-nextjs-todo/internals/db"
	"golang-nextjs-todo/internals/middleware"
	"golang-nextjs-todo/internals/routes"
	"golang-nextjs-todo/internals/services"
	"log"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	server         *gin.Engine
	usercollection *mongo.Collection
	taskcollection *mongo.Collection
	usercontroller controllers.IUserController
	taskcontroller controllers.ITaskController
	userservice    services.IUserService
	taskservice    services.ITaskService
	requireauth    middleware.RequireAuth
	userroute      routes.UserRoutes
	taskroute      routes.TaskRoutes
	ctx            context.Context
)

func init() {
	ctx = context.TODO()
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Some error occured. Err: %s", err)
		}
	}
	// collections
	db.ConnectDB("MONGO_URI")
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
	server.POST("/query", requireauth.SetJWT, graphqlHandler())
	server.GET("/", playgroundHandler())
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "https://go-next-tasks.vercel.app", "https://go-next-tasks-addzr2bzh-hiroki0116.vercel.app"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	server.Use(gin.Logger())
}

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph1.Resolver{TaskController: taskcontroller, UserController: usercontroller}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	defer db.MongoDB.Disconnect(ctx)
	basepath := server.Group("/api/v1")

	userroute.UserRoutes(basepath)
	taskroute.TaskRoutes(basepath)
	log.Fatalln(server.Run(":8080"))
}
