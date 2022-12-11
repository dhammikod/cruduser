package main

import (
	"github.com/dhammikod/cruduser/controllers"
	"github.com/dhammikod/cruduser/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnecttoDB()
}

func main() {
	loginregister()
	api()
}

func loginregister() {

}

func api() {
	user()
}

func user() {
	r := gin.Default()

	r.POST("/users", controllers.UserCreate)
	r.GET("/users", controllers.UsersIndex)
	r.GET("/users/:id", controllers.UsersShow)
	r.PUT("/users/:id", controllers.UsersUpdate)
	r.DELETE("/users/:id", controllers.UsersDelete)
	r.Run()
}
