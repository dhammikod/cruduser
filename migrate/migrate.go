package main

import (
	"fmt"

	"github.com/dhammikod/cruduser/initializers"
	"github.com/dhammikod/cruduser/models"

	"golang.org/x/crypto/bcrypt"
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

	//creating Users
	hash, err := bcrypt.GenerateFromPassword([]byte("password123"), 10)
	if err != nil {
		fmt.Println("ada error")
	}
	initializers.DB.Create(&models.User{Name: "Dhammiko", Email: "Dhammiko@gmail.com", Password: string(hash), No_telp: "082233123", Profile_picture: "profle picture", Notification: false})
	initializers.DB.Create(&models.User{Name: "hagen", Email: "hagen@gmail.com", Password: string(hash), No_telp: "324523212", Profile_picture: "profle picture", Notification: false})
	initializers.DB.Create(&models.User{Name: "siapa", Email: "siapa1@gmail.com", Password: string(hash), No_telp: "812361273", Profile_picture: "profle picture", Notification: false})
	initializers.DB.Create(&models.User{Name: "ludwig", Email: "ludwig@gmail.com", Password: string(hash), No_telp: "657322311", Profile_picture: "profle picture", Notification: false})
	initializers.DB.Create(&models.User{Name: "magnus", Email: "magnus@gmail.com", Password: string(hash), No_telp: "856362211", Profile_picture: "profle picture", Notification: false})

	//creating reseps
	initializers.DB.Create(&models.Resep{Created_by: 1, Judul: "nasi goreng", Steps: "Cicak", Portionsize: 3, Description: "nasi goreng khusus mahasiswa on budget", Rating: 3, Jumlahrating: 0})
	initializers.DB.Create(&models.Resep{Created_by: 2, Judul: "bebek bakar", Steps: "Cicak", Portionsize: 1, Description: "aves mewah", Rating: 5, Jumlahrating: 0})
	initializers.DB.Create(&models.Resep{Created_by: 1, Judul: "mie goreng", Steps: "Cicak", Portionsize: 1, Description: "mie goreng khusus mahasiswa on budget", Rating: 1, Jumlahrating: 0})
	initializers.DB.Create(&models.Resep{Created_by: 3, Judul: "indomie mewah", Steps: "Cicak", Portionsize: 1, Description: "mie goreng indomie bernutrisi", Rating: 5, Jumlahrating: 0})
	initializers.DB.Create(&models.Resep{Created_by: 4, Judul: "fu yung hai", Steps: "Cicak", Portionsize: 2, Description: "chineese food yang mudah dimasak", Rating: 4, Jumlahrating: 0})

	//create bahan
	initializers.DB.Create(&models.Bahan{Namabahan: "kentang"})
	initializers.DB.Create(&models.Bahan{Namabahan: "nasi"})
	initializers.DB.Create(&models.Bahan{Namabahan: "mie"})
	initializers.DB.Create(&models.Bahan{Namabahan: "bebek"})
	initializers.DB.Create(&models.Bahan{Namabahan: "indomie"})
	initializers.DB.Create(&models.Bahan{Namabahan: "telur"})
	initializers.DB.Create(&models.Bahan{Namabahan: "sosis"})

	//create list bahan
	initializers.DB.Create(&models.Listbahan{Resep_id: 1, Bahan_id: 2, Jumlahbahan: 2})
	initializers.DB.Create(&models.Listbahan{Resep_id: 1, Bahan_id: 1, Jumlahbahan: 1})
	initializers.DB.Create(&models.Listbahan{Resep_id: 1, Bahan_id: 6, Jumlahbahan: 2})
	initializers.DB.Create(&models.Listbahan{Resep_id: 1, Bahan_id: 7, Jumlahbahan: 2})
	initializers.DB.Create(&models.Listbahan{Resep_id: 2, Bahan_id: 4, Jumlahbahan: 1})
	initializers.DB.Create(&models.Listbahan{Resep_id: 3, Bahan_id: 3, Jumlahbahan: 2})
	initializers.DB.Create(&models.Listbahan{Resep_id: 3, Bahan_id: 6, Jumlahbahan: 2})
	initializers.DB.Create(&models.Listbahan{Resep_id: 3, Bahan_id: 7, Jumlahbahan: 2})
	initializers.DB.Create(&models.Listbahan{Resep_id: 4, Bahan_id: 5, Jumlahbahan: 2})
	initializers.DB.Create(&models.Listbahan{Resep_id: 4, Bahan_id: 6, Jumlahbahan: 1})
	initializers.DB.Create(&models.Listbahan{Resep_id: 4, Bahan_id: 7, Jumlahbahan: 1})
	initializers.DB.Create(&models.Listbahan{Resep_id: 5, Bahan_id: 1, Jumlahbahan: 2})
	initializers.DB.Create(&models.Listbahan{Resep_id: 5, Bahan_id: 6, Jumlahbahan: 2})
	initializers.DB.Create(&models.Listbahan{Resep_id: 5, Bahan_id: 7, Jumlahbahan: 2})

	//create saved recipe
	initializers.DB.Table("saved_recipes").Create(map[string]interface{}{"resep_id": 1, "user_id": 4})
	initializers.DB.Table("saved_recipes").Create(map[string]interface{}{"resep_id": 1, "user_id": 5})
	initializers.DB.Table("saved_recipes").Create(map[string]interface{}{"resep_id": 2, "user_id": 5})
	initializers.DB.Table("saved_recipes").Create(map[string]interface{}{"resep_id": 3, "user_id": 5})
	initializers.DB.Table("saved_recipes").Create(map[string]interface{}{"resep_id": 4, "user_id": 5})
	initializers.DB.Table("saved_recipes").Create(map[string]interface{}{"resep_id": 4, "user_id": 1})

}
