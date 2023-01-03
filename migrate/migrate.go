package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image/png"
	"io/ioutil"
	"log"

	// "net/http"
	"os"

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
	hash, err := bcrypt.GenerateFromPassword([]byte("cc"), 10)
	if err != nil {
		fmt.Println("ada error")
	}
	initializers.DB.Create(&models.User{Name: "Dhammiko", Email: "Dhammiko@gmail.com", Password: string(hash), Profile_picture: "profle picture", Notification: false})
	initializers.DB.Create(&models.User{Name: "hagen", Email: "hagen@gmail.com", Password: string(hash), Profile_picture: "profle picture", Notification: false})
	initializers.DB.Create(&models.User{Name: "siapa", Email: "siapa1@gmail.com", Password: string(hash), Profile_picture: "profle picture", Notification: false})
	initializers.DB.Create(&models.User{Name: "ludwig", Email: "ludwig@gmail.com", Password: string(hash), Profile_picture: "profle picture", Notification: false})
	initializers.DB.Create(&models.User{Name: "magnus", Email: "magnus@gmail.com", Password: string(hash), Profile_picture: "profle picture", Notification: false})
	initializers.DB.Create(&models.User{Name: "cc", Email: "cc", Password: string(hash), Profile_picture: "profle picture", Notification: false})

	//creating reseps
	nasigoreng := getgambar("rmimg/nasigoreng.png")
	initializers.DB.Create(&models.Resep{Created_by: 1, Foto: nasigoreng, Judul: "nasi goreng", Steps: "Cicak", Portionsize: 3, Description: "nasi goreng khusus mahasiswa on budget", Timetaken: "30 menit", Totalcal: 1000, Rating: 3, Jumlahrating: 0})
	bebekbakar := getgambar("rmimg/bebekgoreng.png")
	initializers.DB.Create(&models.Resep{Created_by: 2, Foto: bebekbakar, Judul: "bebek bakar", Steps: "Cicak", Portionsize: 1, Description: "aves mewah", Timetaken: "3 jam", Rating: 5, Jumlahrating: 0, Totalcal: 1000})
	miegoreng := getgambar("rmimg/miegoreng.png")
	initializers.DB.Create(&models.Resep{Created_by: 1, Foto: miegoreng, Judul: "mie goreng", Steps: "Cicak", Portionsize: 1, Description: "mie goreng khusus mahasiswa on budget", Timetaken: "30 menit", Totalcal: 500, Rating: 1, Jumlahrating: 0})
	indomiemewah := getgambar("rmimg/indomiemewah.png")
	initializers.DB.Create(&models.Resep{Created_by: 3, Foto: indomiemewah, Judul: "indomie mewah", Steps: "Cicak", Portionsize: 1, Description: "mie goreng indomie bernutrisi", Timetaken: "15 menit", Totalcal: 1000, Rating: 5, Jumlahrating: 0})
	fuyunghai := getgambar("rmimg/fuyunghai.png")
	initializers.DB.Create(&models.Resep{Created_by: 4, Foto: fuyunghai, Judul: "fu yung hai", Steps: "Cicak", Portionsize: 2, Description: "chineese food yang mudah dimasak", Timetaken: "15 menit", Totalcal: 700, Rating: 4, Jumlahrating: 0})

	//create bahan
	kentang := getgambar("rmimg/kentang.png")
	initializers.DB.Create(&models.Bahan{Namabahan: "kentang", Foto: kentang})
	nasi := getgambar("rmimg/nasi.png")
	initializers.DB.Create(&models.Bahan{Namabahan: "nasi", Foto: nasi})
	mie := getgambar("rmimg/mie.png")
	initializers.DB.Create(&models.Bahan{Namabahan: "mie", Foto: mie})
	bebek := getgambar("rmimg/bebek.png")
	initializers.DB.Create(&models.Bahan{Namabahan: "bebek", Foto: bebek})
	indomie := getgambar("rmimg/indomie.png")
	initializers.DB.Create(&models.Bahan{Namabahan: "indomie", Foto: indomie})
	telur := getgambar("rmimg/telur.png")
	initializers.DB.Create(&models.Bahan{Namabahan: "telur", Foto: telur})
	sosis := getgambar("rmimg/sosis.png")
	initializers.DB.Create(&models.Bahan{Namabahan: "sosis", Foto: sosis})
	readgambar(1)
	readgambar(2)
	readgambar(3)
	readgambar(4)
	readgambar(5)
	readgambar(6)
	readgambar(7)

	//create list bahan
	initializers.DB.Create(&models.Listbahan{Resep_id: 1, Bahan_id: 2, Jumlahbahan: "2"})
	initializers.DB.Create(&models.Listbahan{Resep_id: 1, Bahan_id: 1, Jumlahbahan: "1"})
	initializers.DB.Create(&models.Listbahan{Resep_id: 1, Bahan_id: 6, Jumlahbahan: "2"})
	initializers.DB.Create(&models.Listbahan{Resep_id: 1, Bahan_id: 7, Jumlahbahan: "2"})
	initializers.DB.Create(&models.Listbahan{Resep_id: 2, Bahan_id: 4, Jumlahbahan: "1"})
	initializers.DB.Create(&models.Listbahan{Resep_id: 3, Bahan_id: 3, Jumlahbahan: "2"})
	initializers.DB.Create(&models.Listbahan{Resep_id: 3, Bahan_id: 6, Jumlahbahan: "2"})
	initializers.DB.Create(&models.Listbahan{Resep_id: 3, Bahan_id: 7, Jumlahbahan: "2"})
	initializers.DB.Create(&models.Listbahan{Resep_id: 4, Bahan_id: 5, Jumlahbahan: "2"})
	initializers.DB.Create(&models.Listbahan{Resep_id: 4, Bahan_id: 6, Jumlahbahan: "1"})
	initializers.DB.Create(&models.Listbahan{Resep_id: 4, Bahan_id: 7, Jumlahbahan: "1"})
	initializers.DB.Create(&models.Listbahan{Resep_id: 5, Bahan_id: 1, Jumlahbahan: "2"})
	initializers.DB.Create(&models.Listbahan{Resep_id: 5, Bahan_id: 6, Jumlahbahan: "2"})
	initializers.DB.Create(&models.Listbahan{Resep_id: 5, Bahan_id: 7, Jumlahbahan: "2"})

	//create saved recipe
	initializers.DB.Table("saved_recipes").Create(map[string]interface{}{"resep_id": 1, "user_id": 4})
	initializers.DB.Table("saved_recipes").Create(map[string]interface{}{"resep_id": 1, "user_id": 5})
	initializers.DB.Table("saved_recipes").Create(map[string]interface{}{"resep_id": 2, "user_id": 5})
	initializers.DB.Table("saved_recipes").Create(map[string]interface{}{"resep_id": 3, "user_id": 5})
	initializers.DB.Table("saved_recipes").Create(map[string]interface{}{"resep_id": 4, "user_id": 5})
	initializers.DB.Table("saved_recipes").Create(map[string]interface{}{"resep_id": 4, "user_id": 1})
}

func getgambar(link string) string {
	//getting the image from online url
	// resp, err := http.Get(link)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer resp.Body.Close()

	file, err := os.Open(link)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	// Read the image data from the response body
	imageData, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	// Encode the image data as a base64 string
	imageDataString := base64.StdEncoding.EncodeToString(imageData)

	//decoding the image to []byte
	// imageDatas, err := base64.StdEncoding.DecodeString(imageDataString)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	//converting the byte to have a .jpg file extension
	// image, err := png.Decode(bytes.NewReader(imageDatas))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Encode the image object to a []byte
	// buffer := new(bytes.Buffer)
	// err = png.Encode(buffer, image)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	//returning the data
	// return buffer.Bytes()
	return imageDataString
}

func readgambar(id int) {
	// Retrieve the image from the database
	var bahan models.Bahan
	initializers.DB.First(&bahan, id)

	//making directory
	// err := os.MkdirAll(os.ExpandEnv("$HOME/images"), 0755)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Create a file to write the image data to
	file, err := os.Create("images/" + bahan.Namabahan + ".png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// decoding the image to []byte
	imageDatas, err := base64.StdEncoding.DecodeString(bahan.Foto)
	if err != nil {
		log.Fatal(err)
	}

	// converting the byte to have a .jpg file extension
	image, err := png.Decode(bytes.NewReader(imageDatas))
	if err != nil {
		log.Fatal(err)
	}

	// Encode the image object to a []byte
	// buffer := new(bytes.Buffer)
	// err = png.Encode(buffer, image)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Decode the image data into an image.Image object
	// img, _, err := image.Decode(bytes.NewReader(bahan.Foto))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Write the image data to the file
	err = png.Encode(file, image)
	if err != nil {
		log.Fatal(err)
	}
}
