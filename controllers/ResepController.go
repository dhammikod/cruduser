package controllers

import (
	"encoding/base64"
	"net/http"

	"github.com/dhammikod/cruduser/initializers"
	"github.com/dhammikod/cruduser/models"
	"github.com/gin-gonic/gin"
)

func ResepCreate(c *gin.Context) {
	//get data
	var body struct {
		Created_by   int
		Rating       float32
		jumlahrating int
		Description  string
		Judul        string
		Portionsize  int
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "400",
			"error":  "Failed to read body",
		})

		return
	}

	//create a user
	resep := models.Resep{Created_by: body.Created_by, Judul: body.Judul, Foto: nil, Video: nil, Portionsize: body.Portionsize, Description: body.Description, Rating: body.Rating}
	result := initializers.DB.Create(&resep)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"status": "400",
		})
		return
	}
	//return user
	c.JSON(200, gin.H{
		"status": "200",
	})
}

func ResepIndex(c *gin.Context) {
	//get posts
	var reseps []models.Resep
	initializers.DB.Preload("User").Preload("Listbahan").Preload("Listbahan.Bahan").Find(&reseps)

	//respond to the posts
	c.JSON(200, gin.H{
		"status": "200",
		"Resep":  reseps,
	})
}

func ResepShow(c *gin.Context) {
	//get id
	id := c.Param("id")
	//get posts
	var resep models.Resep
	initializers.DB.Preload("User").Preload("Listbahan").Preload("Listbahan.Bahan").First(&resep, id)

	//respond to the posts
	c.JSON(200, gin.H{
		"status": "200",
		"resep":  resep,
	})
}

func ResepUpdate(c *gin.Context) {
	//get id
	id := c.Param("id")

	//get data off req body
	var body struct {
		Rating      float32
		Description string
		Judul       string
		Portionsize int
		Foto        string
		Video       string
	}
	c.Bind(&body)

	//decode the foto
	imageData, err := base64.StdEncoding.DecodeString(body.Foto)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"status": "500",
		})
		return
	}

	//decode the video
	videoData, err := base64.StdEncoding.DecodeString(body.Video)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"status": "500",
		})
		return
	}
	//find the post to update
	var resep models.Resep
	initializers.DB.First(&resep, id)

	//update
	initializers.DB.Model(&resep).Updates(models.Resep{
		Rating:      body.Rating,
		Description: body.Description,
		Judul:       body.Judul,
		Portionsize: body.Portionsize,
		Foto:        imageData,
		Video:       videoData,
	})

	//return updated value
	c.JSON(http.StatusOK, gin.H{
		"status": "200",
	})
}

func ResepDelete(c *gin.Context) {
	//get id
	id := c.Param("id")

	//delete
	initializers.DB.Delete(&models.Resep{}, id)
	//return value
	c.JSON(http.StatusOK, gin.H{
		"status": "200",
	})
	return
}
