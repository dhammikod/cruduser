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

	//crud user
	r.GET("/users", controllers.UsersIndex)
	r.GET("/users/:id", controllers.UsersShow)
	r.PUT("/users/:id", controllers.UsersUpdate)
	r.DELETE("/users/:id", controllers.UsersDelete)

	//crud resep
	r.POST("/resep", controllers.ResepCreate)
	r.GET("/resep", controllers.ResepIndex)
	r.GET("/resep/:id", controllers.ResepShow)
	r.PUT("/resep/:id", controllers.ResepUpdate)
	r.DELETE("/resep/:id", controllers.ResepDelete)

	//crud bahan
	r.POST("/bahan", controllers.BahanCreate)
	r.GET("/bahan", controllers.BahanIndex)
	r.GET("/bahan/:id", controllers.BahanShow)
	r.PUT("/bahan/:id", controllers.BahanUpdate)
	r.DELETE("/bahan/:id", controllers.BahanDelete)

	//crud listbahan
	r.POST("/listbahan", controllers.ListBahanCreate)
	r.GET("/listbahan", controllers.ListBahanIndex)
	r.GET("/listbahan/:id", controllers.ListBahanShow)
	r.PUT("/listbahan/:id", controllers.ListBahanUpdate)
	r.DELETE("/listbahan/:id", controllers.ListBahanDelete)

	//crud saved recipe
	r.POST("/savedrecipe", controllers.SavedRecipeCreate)
	r.GET("/savedrecipe", controllers.SavedRecipeIndex)
	r.GET("/savedrecipe/:id", controllers.SavedRecipeShow)
	r.POST("/deletesavedresep", controllers.SavedRecipeDelete)
	r.Run()
}
