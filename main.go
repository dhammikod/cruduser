package main

import (
	"github.com/dhammikod/cruduser/controllers"
	"github.com/dhammikod/cruduser/initializers"
	"github.com/dhammikod/cruduser/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnecttoDB()
}

func main() {
	r := gin.Default()

	//register
	r.POST("/users", controllers.UserCreate)

	//login
	r.POST("/login", controllers.Login)

	//validation
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	//crud
	r.GET("/users", controllers.UsersIndex)
	r.GET("/users/:id", controllers.UsersShow)
	r.PUT("/users/:id", controllers.UsersUpdate)
	r.DELETE("/users/:id", controllers.UsersDelete)
	r.Run()
}
