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

	users := r.Group("/users")
	{
		//register
		users.POST("/", controllers.UserCreate)

		//crud user
		users.GET("/", middleware.RequireAuth, controllers.UsersIndex)
		users.GET("/:id", middleware.RequireAuth, controllers.UsersShow)
		users.PUT("/:id", middleware.RequireAuth, controllers.UsersUpdate)
		users.DELETE("/:id", middleware.RequireAuth, controllers.UsersDelete)
	}

	//login
	r.POST("/login", controllers.Login)

	//tes validation
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	//crud resep
	reseps := r.Group("/resep").Use(middleware.RequireAuth)
	{
		reseps.POST("/", controllers.ResepCreate)
		reseps.GET("/", controllers.ResepIndex)
		reseps.GET("/:id", controllers.ResepShow)
		reseps.GET("/user/:id", controllers.Resepbyuser)
		reseps.POST("/bahan/", controllers.Resepbybahan)
		reseps.PUT("/:id", controllers.ResepUpdate)
		reseps.DELETE("/:id", controllers.ResepDelete)
	}

	bahans := r.Group("/bahan").Use(middleware.RequireAuth)
	{
		//crud bahan
		bahans.POST("/", controllers.BahanCreate)
		bahans.GET("/", controllers.BahanIndex)
		bahans.GET("/:id", controllers.BahanShow)
		bahans.PUT("/:id", controllers.BahanUpdate)
		bahans.DELETE("/:id", controllers.BahanDelete)
	}

	listbahan := r.Group("/listbahan").Use(middleware.RequireAuth)
	{
		//crud listbahan
		listbahan.POST("/", controllers.ListBahanCreate)
		listbahan.GET("/", controllers.ListBahanIndex)
		listbahan.GET("/:id", controllers.ListBahanShow)
		listbahan.PUT("/:id", controllers.ListBahanUpdate)
		listbahan.DELETE("/:id", controllers.ListBahanDelete)
	}

	savedrecipes := r.Group("/savedrecipe").Use(middleware.RequireAuth)
	{
		//crud saved recipe
		savedrecipes.POST("/", controllers.SavedRecipeCreate)
		savedrecipes.GET("/", controllers.SavedRecipeIndex)
		savedrecipes.GET("/:id", controllers.SavedRecipeShow)
		savedrecipes.POST("/deletesavedresep", controllers.SavedRecipeDelete)
	}

	r.Run()
}
