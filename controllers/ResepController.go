package controllers

import (
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
		Foto         string
		Video        string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	//create a user

	resep := models.Resep{Created_by: body.Created_by, Judul: body.Judul, Foto: body.Foto, Video: body.Video, Portionsize: body.Portionsize, Description: body.Description, Rating: body.Rating}
	result := initializers.DB.Create(&resep)

	if result.Error != nil {
		c.Status(400)
		return
	}
	//return user
	c.Status(200)
}

func ResepIndex(c *gin.Context) {
	//get posts
	var reseps []models.Resep
	initializers.DB.Preload("User").Preload("Listbahan").Preload("Listbahan.Bahan").Find(&reseps)

	//respond to the posts
	c.JSON(200, gin.H{
		"Resep": reseps,
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
		"resep": resep,
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

	//find the post to update
	var resep models.Resep
	initializers.DB.First(&resep, id)

	//update
	initializers.DB.Model(&resep).Updates(models.Resep{
		Rating:      body.Rating,
		Description: body.Description,
		Judul:       body.Judul,
		Portionsize: body.Portionsize,
		Foto:        body.Foto,
		Video:       body.Video,
	})

	//return updated value
	c.Status(200)
}

func ResepDelete(c *gin.Context) {
	//get id
	id := c.Param("id")

	//delete
	initializers.DB.Delete(&models.Resep{}, id)
	//return value
	c.Status(200)
}
