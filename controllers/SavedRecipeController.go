package controllers

import (
	"net/http"

	"github.com/dhammikod/cruduser/initializers"
	"github.com/dhammikod/cruduser/models"
	"github.com/gin-gonic/gin"
)

func Existornot(c *gin.Context) {
	var body struct {
		User_id  int
		Resep_id int
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "400",
			"error":  "Failed to read body",
		})

		return
	}

	var savedada bool
	initializers.DB.Raw("SELECT EXISTS(SELECT * FROM `saved_recipes` WHERE resep_id = ? AND user_id = ?)", body.Resep_id, body.User_id).Scan(&savedada)
	var userada bool
	initializers.DB.Raw("SELECT EXISTS(SELECT * FROM `users` WHERE id = ?)", body.User_id).Scan(&userada)
	var recipeada bool
	initializers.DB.Raw("SELECT EXISTS(SELECT * FROM `reseps` WHERE id = ?)", body.Resep_id).Scan(&recipeada)

	c.JSON(http.StatusOK, gin.H{
		"statusresep":       recipeada,
		"statususer":        userada,
		"statussavedrecipe": savedada,
	})
}

// err := initializers.DB.Unscoped().
// 	Exec("SELECT EXISTS(SELECT * FROM `saved_recipes` WHERE resep_id = ? AND user_id = ?)", body.Resep_id, body.User_id)

func SavedRecipeCreate(c *gin.Context) {
	//get data
	var body struct {
		User_id  int
		Resep_id int
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "400",
			"error":  "Failed to read body",
		})

		return
	}

	var savedada bool
	initializers.DB.Raw("SELECT EXISTS(SELECT * FROM `saved_recipes` WHERE resep_id = ? AND user_id = ?)", body.Resep_id, body.User_id).Scan(&savedada)
	var userada bool
	initializers.DB.Raw("SELECT EXISTS(SELECT * FROM `users` WHERE id = ?)", body.User_id).Scan(&userada)
	var recipeada bool
	initializers.DB.Raw("SELECT EXISTS(SELECT * FROM `reseps` WHERE id = ?)", body.Resep_id).Scan(&recipeada)

	if userada && recipeada {
		if savedada {
			//delete
			initializers.DB.Unscoped().
				Exec("DELETE FROM `saved_recipes` WHERE resep_id = ? AND user_id = ?", body.Resep_id, body.User_id)

			c.JSON(http.StatusOK, gin.H{
				"status": "200",
				"msg":    "savedrecipe deleted",
				// "statusupda": err,
			})
		} else {
			//creating savedrecipe
			initializers.DB.Table("saved_recipes").Create(map[string]interface{}{
				"resep_id": body.Resep_id, "user_id": body.User_id,
			})

			c.JSON(http.StatusOK, gin.H{
				"status": "200",
				"msg":    "savedrecipecreated",
				// "statusupda": err,
			})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "400",
			"error":  "No key found",
			// "statusupda": err,
		})

		return
	}
}

func SavedRecipeIndex(c *gin.Context) {
	//get posts
	var users []models.User
	initializers.DB.Preload("Resep").
		Preload("Resep.User").Preload("Resep.Listbahan").Preload("Resep.Listbahan.Bahan").
		Find(&users)

	//respond to the posts
	c.JSON(200, gin.H{
		"status":        "200",
		"saved_recipes": users,
	})
}

func SavedRecipeShow(c *gin.Context) {
	//get id
	id := c.Param("id")
	//get posts
	var user models.User

	initializers.DB.Preload("Resep").
		Preload("Resep.User").Preload("Resep.Listbahan").Preload("Resep.Listbahan.Bahan").
		First(&user, id)

	//respond to the posts
	c.JSON(200, gin.H{
		"status":       "200",
		"Saved_recipe": user,
	})
}

func SavedRecipeDelete(c *gin.Context) {
	var body struct {
		User_id  int
		Resep_id int
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "400",
			"error":  "Failed to read body",
		})

		return
	}

	//delete
	result := initializers.DB.Unscoped().
		Exec("DELETE FROM `saved_recipes` WHERE resep_id = ? AND user_id = ?", body.Resep_id, body.User_id)

	//return value
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "400",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "200",
	})
	return
}
