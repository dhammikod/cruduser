package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
	"log"
	"net/http"
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
	hash, err := bcrypt.GenerateFromPassword([]byte("password123"), 10)
	if err != nil {
		fmt.Println("ada error")
	}
	initializers.DB.Create(&models.User{Name: "Dhammiko", Email: "Dhammiko@gmail.com", Password: string(hash), Profile_picture: "profle picture", Notification: false})
	initializers.DB.Create(&models.User{Name: "hagen", Email: "hagen@gmail.com", Password: string(hash), Profile_picture: "profle picture", Notification: false})
	initializers.DB.Create(&models.User{Name: "siapa", Email: "siapa1@gmail.com", Password: string(hash), Profile_picture: "profle picture", Notification: false})
	initializers.DB.Create(&models.User{Name: "ludwig", Email: "ludwig@gmail.com", Password: string(hash), Profile_picture: "profle picture", Notification: false})
	initializers.DB.Create(&models.User{Name: "magnus", Email: "magnus@gmail.com", Password: string(hash), Profile_picture: "profle picture", Notification: false})

	//creating reseps
	initializers.DB.Create(&models.Resep{Created_by: 1, Judul: "nasi goreng", Steps: "Cicak", Portionsize: 3, Description: "nasi goreng khusus mahasiswa on budget", Timetaken: "30 menit", Totalcal: 1000, Rating: 3, Jumlahrating: 0})
	initializers.DB.Create(&models.Resep{Created_by: 2, Judul: "bebek bakar", Steps: "Cicak", Portionsize: 1, Description: "aves mewah", Timetaken: "3 jam", Rating: 5, Jumlahrating: 0, Totalcal: 1000})
	initializers.DB.Create(&models.Resep{Created_by: 1, Judul: "mie goreng", Steps: "Cicak", Portionsize: 1, Description: "mie goreng khusus mahasiswa on budget", Timetaken: "30 menit", Totalcal: 500, Rating: 1, Jumlahrating: 0})
	initializers.DB.Create(&models.Resep{Created_by: 3, Judul: "indomie mewah", Steps: "Cicak", Portionsize: 1, Description: "mie goreng indomie bernutrisi", Timetaken: "15 menit", Totalcal: 1000, Rating: 5, Jumlahrating: 0})
	initializers.DB.Create(&models.Resep{Created_by: 4, Judul: "fu yung hai", Steps: "Cicak", Portionsize: 2, Description: "chineese food yang mudah dimasak", Timetaken: "15 menit", Totalcal: 700, Rating: 4, Jumlahrating: 0})

	//create bahan
	kentang := getgambar("https://w7.pngwing.com/pngs/1018/700/png-transparent-potato-vegetable-food-fruit-potato-soup-food-baking.png")
	initializers.DB.Create(&models.Bahan{Namabahan: "kentang", Foto: kentang})
	nasi := getgambar("https://e7.pngegg.com/pngimages/622/651/png-clipart-rice-rice.png")
	initializers.DB.Create(&models.Bahan{Namabahan: "nasi", Foto: nasi})
	mie := getgambar("https://png.pngtree.com/png-vector/20210930/ourmid/pngtree-raw-noodles-healthy-pastry-nutritional-ingredients-png-image_3963645.png")
	initializers.DB.Create(&models.Bahan{Namabahan: "mie", Foto: mie})
	bebek := getgambar("https://w7.pngwing.com/pngs/146/121/png-transparent-duck-meat-chicken-leg-food-duck-animals-liver-animal-source-foods.png")
	initializers.DB.Create(&models.Bahan{Namabahan: "bebek", Foto: bebek})
	indomie := getgambar("https://w7.pngwing.com/pngs/318/116/png-transparent-indomie-instant-noodle-mie-goreng-fried-noodles-chow-mein-chicken-food-animals-recipe.png")
	initializers.DB.Create(&models.Bahan{Namabahan: "indomie", Foto: indomie})
	telur := getgambar("https://w7.pngwing.com/pngs/321/422/png-transparent-chicken-egg-yolk-food-egg-food-chicken-egg-yolk.png")
	initializers.DB.Create(&models.Bahan{Namabahan: "telur", Foto: telur})
	sosis := getgambar("https://e7.pngegg.com/pngimages/19/528/png-clipart-frankfurter-wurstchen-hot-dog-bockwurst-sausage-roll-hot-dog-food-beef.png")
	initializers.DB.Create(&models.Bahan{Namabahan: "sosis", Foto: sosis})
	readgambar(1)
	readgambar(2)
	readgambar(3)
	readgambar(4)
	readgambar(5)
	readgambar(6)
	readgambar(7)

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

func getgambar(link string) []byte {
	//getting the image from online url
	resp, err := http.Get(link)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Read the image data from the response body
	imageData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Encode the image data as a base64 string
	imageDataString := base64.StdEncoding.EncodeToString(imageData)

	//decoding the image to []byte
	imageDatas, err := base64.StdEncoding.DecodeString(imageDataString)
	if err != nil {
		log.Fatal(err)
	}

	//converting the byte to have a .jpg file extension
	image, err := png.Decode(bytes.NewReader(imageDatas))
	if err != nil {
		log.Fatal(err)
	}

	// Encode the image object to a []byte
	buffer := new(bytes.Buffer)
	err = png.Encode(buffer, image)
	if err != nil {
		log.Fatal(err)
	}

	//returning the data
	return buffer.Bytes()
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

	// Decode the image data into an image.Image object
	img, _, err := image.Decode(bytes.NewReader(bahan.Foto))
	if err != nil {
		log.Fatal(err)
	}

	// Write the image data to the file
	err = png.Encode(file, img)
	if err != nil {
		log.Fatal(err)
	}
}
