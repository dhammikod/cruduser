package controllers

import (
	"net/http"

	"github.com/dhammikod/cruduser/initializers"
	"github.com/dhammikod/cruduser/models"
	"github.com/gin-gonic/gin"
)

func BahanCreate(c *gin.Context) {
	//get data
	var body struct {
		Namabahan string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	//create a user

	bahan := models.Bahan{Namabahan: body.Namabahan}
	result := initializers.DB.Create(&bahan)

	if result.Error != nil {
		c.Status(400)
		return
	}
	//return user
	c.Status(200)
}

func BahanIndex(c *gin.Context) {
	//get posts
	var bahans []models.Bahan
	initializers.DB.Find(&bahans)

	//respond to the posts
	c.JSON(200, gin.H{
		"Resep": bahans,
	})
}

func BahanShow(c *gin.Context) {
	//get id
	id := c.Param("id")
	//get posts
	var bahan models.Bahan
	initializers.DB.First(&bahan, id)

	//respond to the posts
	c.JSON(200, gin.H{
		"bahan": bahan,
	})
}

func BahanUpdate(c *gin.Context) {
	//get id
	id := c.Param("id")

	//get data off req body
	var body struct {
		Namabahan string
	}
	c.Bind(&body)

	//find the post to update
	var bahan models.Bahan
	initializers.DB.First(&bahan, id)

	//update
	initializers.DB.Model(&bahan).Updates(models.Bahan{
		Namabahan: body.Namabahan,
	})

	//return updated value
	c.Status(200)
}

func BahanDelete(c *gin.Context) {
	//get id
	id := c.Param("id")

	//delete
	initializers.DB.Delete(&models.Bahan{}, id)
	//return value
	c.Status(200)
}
