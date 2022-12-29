package controllers

import (
	"encoding/base64"
	"net/http"

	"github.com/dhammikod/cruduser/initializers"
	"github.com/dhammikod/cruduser/models"
	"github.com/gin-gonic/gin"
)

func BahanCreate(c *gin.Context) {
	//get data
	var body struct {
		Foto      string
		Namabahan string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "400",
			"error":  "Failed to read body",
		})

		return
	}

	//create a user
	imageData, err := base64.StdEncoding.DecodeString(body.Foto)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"status": "500",
		})
		return
	}

	bahan := models.Bahan{Namabahan: body.Namabahan, Foto: imageData}
	result := initializers.DB.Create(&bahan)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "400",
			"error":  "Something went wrong",
		})
		return
	}
	//return user
	c.JSON(http.StatusOK, gin.H{
		"status": "200",
	})
}

func BahanIndex(c *gin.Context) {
	//get posts
	var bahans []models.Bahan
	initializers.DB.Find(&bahans)

	//respond to the posts
	c.JSON(200, gin.H{
		"status": "200",
		"Bahan":  bahans,
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
		"status": "200",
		"Bahan":  bahan,
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
