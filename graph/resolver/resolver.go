package graph

import "golang-nextjs-todo/internals/controllers"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	TaskController controllers.ITaskController
	UserController controllers.IUserController
}
