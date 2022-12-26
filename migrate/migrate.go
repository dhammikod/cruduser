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
	initializers.DB.AutoMigrate(&models.Resep{})
	initializers.DB.AutoMigrate(&models.Bahan{})
	initializers.DB.AutoMigrate(&models.Listbahan{})
}
