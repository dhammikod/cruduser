package main

import (
	"github.com/dhammikod/cruduser/initializers"
	"github.com/dhammikod/cruduser/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnecttoDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
}
