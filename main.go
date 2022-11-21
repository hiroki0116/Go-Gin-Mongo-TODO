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
	usercontroller controllers.UserController
	userservice    services.UserService
	route          routes.UserRoutes
	ctx            context.Context
	err            error
)

func init() {
	ctx = context.TODO()
	mongoconn := options.Client().ApplyURI("mongodb://localhost:27017")
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
	usercontroller = controllers.NewUserContoller(usercollection, ctx)
	userservice = services.New(usercontroller)
	route = routes.NewRoute(userservice)
	server = gin.Default()
}

func main() {
	defer mongoclient.Disconnect(ctx)
	basepath := server.Group("/api/v1")

	route.UserRoutes(basepath)

	log.Fatalln(server.Run(":8000"))
}
