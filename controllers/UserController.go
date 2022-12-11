package controllers

import (
	"net/http"

	"github.com/dhammikod/cruduser/initializers"
	"github.com/dhammikod/cruduser/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func UserCreate(c *gin.Context) {
	//get data
	var body struct {
		Name            string
		Email           string
		Password        string
		No_telp         string
		Profile_picture string
		Notification    bool
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}
	//hashing the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
	}
	//create a user

	user := models.User{Name: body.Name, Email: body.Email, Password: string(hash), No_telp: body.No_telp, Profile_picture: body.Profile_picture, Notification: body.Notification}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.Status(400)
		return
	}
	//return user
	c.JSON(200, gin.H{
		"user": user,
	})
}

func UsersIndex(c *gin.Context) {
	//get posts
	var users []models.User
	initializers.DB.Find(&users)

	//respond to the posts
	c.JSON(200, gin.H{
		"user": users,
	})
}

func UsersShow(c *gin.Context) {
	//get id
	id := c.Param("id")
	//get posts
	var user models.User
	initializers.DB.First(&user, id)

	//respond to the posts
	c.JSON(200, gin.H{
		"user": user,
	})
}

func UsersUpdate(c *gin.Context) {
	//get id
	id := c.Param("id")

	//get data off req body
	var body struct {
		Name            string
		Email           string
		Password        string
		No_telp         string
		Profile_picture string
		Notification    bool
	}
	c.Bind(&body)

	//find the post to update
	var user models.User
	initializers.DB.First(&user, id)

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
	}

	//update
	initializers.DB.Model(&user).Updates(models.User{
		Name:            body.Name,
		Email:           body.Email,
		Password:        string(hash),
		No_telp:         body.No_telp,
		Profile_picture: body.Profile_picture,
		Notification:    body.Notification,
	})

	//return updated value
	c.JSON(200, gin.H{
		"user": user,
	})
}

func UsersDelete(c *gin.Context) {
	//get id
	id := c.Param("id")

	//delete
	initializers.DB.Delete(&models.User{}, id)
	//return value
	c.Status(200)
}

func login(c *gin.Context) {
	// Get the email and pass off req body

	//look up requested user

	//compare sent in pass with saved user pass

	//generate a jwt token

	//send it back
}
