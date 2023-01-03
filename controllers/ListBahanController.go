package controllers

import (
	"net/http"

	"github.com/dhammikod/cruduser/initializers"
	"github.com/dhammikod/cruduser/models"
	"github.com/gin-gonic/gin"
)

func ListBahanCreate(c *gin.Context) {
	//get data
	var body struct {
		Resep_id    uint
		Bahan_id    int
		Jumlahbahan string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "400",
			"error":  "Failed to read body",
		})

		return
	}

	//create a user

	listbahan := models.Listbahan{Resep_id: body.Resep_id, Bahan_id: body.Bahan_id, Jumlahbahan: body.Jumlahbahan}
	result := initializers.DB.Create(&listbahan)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "400",
			"error":  "something went wrong",
		})
		return
	}

	//return user
	c.JSON(http.StatusOK, gin.H{
		"status": "200",
	})
}

func ListBahanIndex(c *gin.Context) {
	//get posts
	var bahans []models.Listbahan
	initializers.DB.Preload("Resep").Preload("Bahan").Preload("User").Find(&bahans)

	//respond to the posts
	c.JSON(200, gin.H{
		"status":    "200",
		"listbahan": bahans,
	})
}

func ListBahanShow(c *gin.Context) {
	//get id
	id := c.Param("id")
	//get posts
	var bahan models.Listbahan
	initializers.DB.Preload("Resep").Preload("Bahan").First(&bahan, id)

	//respond to the posts
	c.JSON(200, gin.H{
		"status":    "200",
		"listbahan": bahan,
	})
}

func ListBahanUpdate(c *gin.Context) {
	//get id
	id := c.Param("id")

	//get data off req body
	var body struct {
		Resep_id    uint
		Bahan_id    int
		Jumlahbahan string
	}
	c.Bind(&body)

	//find the post to update
	var bahan models.Listbahan
	initializers.DB.First(&bahan, id)

	//update
	initializers.DB.Model(&bahan).Updates(models.Listbahan{
		Resep_id: body.Resep_id, Bahan_id: body.Bahan_id, Jumlahbahan: body.Jumlahbahan,
	})

	//return updated value
	c.JSON(200, gin.H{
		"status": "200",
	})
}

func ListBahanDelete(c *gin.Context) {
	//get id
	id := c.Param("id")

	//delete
	initializers.DB.Delete(&models.Listbahan{}, id)
	//return value
	c.JSON(200, gin.H{
		"status": "200",
	})
}
