package controllers

import (
	"net/http"

	"github.com/dhammikod/cruduser/initializers"
	"github.com/dhammikod/cruduser/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

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

	err := initializers.DB.Unscoped().
		Table("saved_recipes").
		Where("resep_id = ? AND user_id = ?", body.Resep_id, body.User_id).
		First(&models.Savedrecipe{}).Error

	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "400",
			"result": "multiple record found",
		})

		return
	}

	//create a user
	result := initializers.DB.Table("saved_recipes").Create(map[string]interface{}{
		"resep_id": body.Resep_id, "user_id": body.User_id,
	})

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "400",
			"error":  "No key found",
		})

		return
	}
	//return user
	c.JSON(http.StatusBadRequest, gin.H{
		"status": "200",
	})
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
		"saved recipes": users,
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
		"Saved recipe": user,
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
