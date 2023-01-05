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
	initializers.DB.Create(&models.Resep{Created_by: 1, Foto: nasigoreng, Judul: "nasi goreng", Steps: "1. Kukus nasi atau cari nasi yang sudah masak, 2. Potong potong sosis menjadi bagian kecil, 3. Goreng telur, 4. Panaskan teflon, 5. Tuang minyak ke dalam teflon, 6. Setlah minyak panas masukkan nasi sosis dan telur bersamaan, 7. Tambah kecap dan aduk sampai merata ", Portionsize: 3, Description: "nasi goreng khusus mahasiswa on budget", Timetaken: "30 menit", Totalcal: 2000, Rating: 3, Jumlahrating: 0})
	bebekbakar := getgambar("rmimg/bebekgoreng.png")
	initializers.DB.Create(&models.Resep{Created_by: 2, Foto: bebekbakar, Judul: "bebek bakar", Steps: "1. Bersihkan bebek, 2. Rebus bebek di dalam rendaman air bumbu selama 2 jam, 3. Keluarkan bebek kemudian berikan tepung dan telur, 4. panaskan minyak dengan jumlah sehingga bebeknya bisa menyelam, 5. Setelah minyak panas goreng bebek hingga matang,", Portionsize: 1, Description: "aves mewah", Timetaken: "3 jam", Rating: 5, Jumlahrating: 0, Totalcal: 1000})
	miegoreng := getgambar("rmimg/miegoreng.png")
	initializers.DB.Create(&models.Resep{Created_by: 1, Foto: miegoreng, Judul: "mie kering", Steps: "1. Goreng mie tipis sampai kering, 2. Siapkan kuah yang terdiri dari sawi kanji tomat micin air secukupnya, 3. Setelah kuah siap campurkan mie dan kuah, ", Portionsize: 1, Description: "mie goreng khusus mahasiswa on budget", Timetaken: "30 menit", Totalcal: 500, Rating: 1, Jumlahrating: 0})
	// indomiemewah := getgambar("rmimg/indomiemewah.png")
	// initializers.DB.Create(&models.Resep{Created_by: 3, Foto: indomiemewah, Judul: "indomie mewah", Steps: "Cicak", Portionsize: 1, Description: "mie goreng indomie bernutrisi", Timetaken: "15 menit", Totalcal: 1000, Rating: 5, Jumlahrating: 0})
	fuyunghai := getgambar("rmimg/fuyunghai.png")
	initializers.DB.Create(&models.Resep{Created_by: 2, Foto: fuyunghai, Judul: "fu yung hai", Steps: "1. Bersihkan sosis dan bawang, 2.Tumis bawang sampai setengah matang kemudian masukkan sosis ke dalam wajan, 3. setelah sosis dan bawang hampir matang masukkan telur dan aduk sampai telur matang, ", Portionsize: 2, Description: "chineese food yang mudah dimasak", Timetaken: "15 menit", Totalcal: 700, Rating: 4, Jumlahrating: 0})
	pecel := getgambar("rmimg/pecel.png")
	initializers.DB.Create(&models.Resep{Created_by: 2, Foto: pecel, Judul: "pecel", Steps: "1. Bersihkan semua sayuran yang ada, 2.Rebus sayuran yang ada, 3. goreng kacang dan kemudian blender sampai halus, 4. goreng tempe dan tahu sampai matang, 5. Ketika ingin disajikan campur bubuk kacang dengan sedikit air dan beberapa sambal jika ingin pedis lalu campur dengan lauk yang diinginkan", Portionsize: 2, Description: "makanan sehat yang enak", Timetaken: "2 jam", Totalcal: 1000, Rating: 5, Jumlahrating: 0})
	nasikari := getgambar("rmimg/nasikari.png")
	initializers.DB.Create(&models.Resep{Created_by: 5, Foto: nasikari, Judul: "nasi kari", Steps: "1. Kukus nasi atau cari nasi yang sudah masak, 2.Goreng sosis sampai masak, 3. Siapkan bumbu kari lalu ikuti instruksi dari bumbu kari tersebut, 4. Campur bubuk kari degan nasi dan sosis", Portionsize: 2, Description: "makanan jepang on Budget", Timetaken: "30 menit", Totalcal: 1200, Rating: 2, Jumlahrating: 0})
	coto := getgambar("rmimg/coto.png")
	initializers.DB.Create(&models.Resep{Created_by: 6, Foto: coto, Judul: "coto", Steps: "1. Bersihkan jeroan yang ada, 2.Jika menggunakan paru maka goreng yang lainnya di rebus sampai matang, 3. Rebus jahe serai garam gula dan micin bersamaan sampai rasanya enak, 4. Setelah airnya masak maka jeroan di masukkan ke dalam kuahnya", Portionsize: 2, Description: "Makanan khas sulawesi", Timetaken: "3 jam", Totalcal: 700, Rating: 9999, Jumlahrating: 0})

	//create bahan
	kentang := getgambar("rmimg/kentang.png")
	initializers.DB.Create(&models.Bahan{Namabahan: "kentang", Foto: kentang})
	nasi := getgambar("rmimg/nasi.png")
	initializers.DB.Create(&models.Bahan{Namabahan: "nasi", Foto: nasi})
	mie := getgambar("rmimg/mie.png")
	initializers.DB.Create(&models.Bahan{Namabahan: "mie", Foto: mie})
	bebek := getgambar("rmimg/bebek.png")
	initializers.DB.Create(&models.Bahan{Namabahan: "bebek", Foto: bebek})
	// indomie := getgambar("rmimg/indomie.png")
	// initializers.DB.Create(&models.Bahan{Namabahan: "indomie", Foto: indomie})
	telur := getgambar("rmimg/telur.png")
	initializers.DB.Create(&models.Bahan{Namabahan: "telur", Foto: telur})
	sosis := getgambar("rmimg/sosis.png")
	initializers.DB.Create(&models.Bahan{Namabahan: "sosis", Foto: sosis})
	minyak := getgambar("rmimg/minyak.png")
	initializers.DB.Create(&models.Bahan{Namabahan: "minyak", Foto: minyak})
	tepung := getgambar("rmimg/tepung.png")
	initializers.DB.Create(&models.Bahan{Namabahan: "tepung", Foto: tepung})
	mentega := getgambar("rmimg/mentega.png")
	initializers.DB.Create(&models.Bahan{Namabahan: "mentega", Foto: mentega})
	micin := getgambar("rmimg/micin.png")
	initializers.DB.Create(&models.Bahan{Namabahan: "micin", Foto: micin})
	sawi := getgambar("rmimg/sawi.png")
	initializers.DB.Create(&models.Bahan{Namabahan: "sawi", Foto: sawi})
	tepungkanji := getgambar("rmimg/tepungkanji.png")
	initializers.DB.Create(&models.Bahan{Namabahan: "kanji", Foto: tepungkanji})
	tomat := getgambar("rmimg/tomat.png")
	initializers.DB.Create(&models.Bahan{Namabahan: "tomat", Foto: tomat})
	keju := getgambar("rmimg/keju.png")
	initializers.DB.Create(&models.Bahan{Namabahan: "keju", Foto: keju})
	parsli := getgambar("rmimg/parsli.png")
	initializers.DB.Create(&models.Bahan{Namabahan: "parsli", Foto: parsli})
	sapi := getgambar("rmimg/sapi.png")
	initializers.DB.Create(&models.Bahan{Namabahan: "daging sapi", Foto: sapi})
	bawang := getgambar("rmimg/bawang.png")
	initializers.DB.Create(&models.Bahan{Namabahan: "bawang", Foto: bawang})
	taoge := getgambar("rmimg/taoge.png")
	initializers.DB.Create(&models.Bahan{Namabahan: "taoge", Foto: taoge})
	kacangpanjang := getgambar("rmimg/kacangpanjang.png")
	initializers.DB.Create(&models.Bahan{Namabahan: "kacang panjang", Foto: kacangpanjang})
	tahu := getgambar("rmimg/tahu.png")
	initializers.DB.Create(&models.Bahan{Namabahan: "tahu", Foto: tahu})
	tempe := getgambar("rmimg/tempe.png")
	initializers.DB.Create(&models.Bahan{Namabahan: "tempe", Foto: tempe})
	kacang := getgambar("rmimg/kacang.png")
	initializers.DB.Create(&models.Bahan{Namabahan: "kacang", Foto: kacang})
	kari := getgambar("rmimg/kari.png")
	initializers.DB.Create(&models.Bahan{Namabahan: "bubuk kari", Foto: kari})
	gula := getgambar("rmimg/gula.png")
	initializers.DB.Create(&models.Bahan{Namabahan: "gula", Foto: gula})
	garam := getgambar("rmimg/garam.png")
	initializers.DB.Create(&models.Bahan{Namabahan: "garam", Foto: garam})
	sambal := getgambar("rmimg/sambal.png")
	initializers.DB.Create(&models.Bahan{Namabahan: "sambal", Foto: sambal})

	//create list bahan
	initializers.DB.Create(&models.Listbahan{Resep_id: 1, Bahan_id: 2, Jumlahbahan: "2"})
	initializers.DB.Create(&models.Listbahan{Resep_id: 1, Bahan_id: 1, Jumlahbahan: "1"})
	initializers.DB.Create(&models.Listbahan{Resep_id: 1, Bahan_id: 6, Jumlahbahan: "2"})
	initializers.DB.Create(&models.Listbahan{Resep_id: 1, Bahan_id: 7, Jumlahbahan: "2"})
	initializers.DB.Create(&models.Listbahan{Resep_id: 2, Bahan_id: 4, Jumlahbahan: "1"})
	initializers.DB.Create(&models.Listbahan{Resep_id: 3, Bahan_id: 3, Jumlahbahan: "2"})
	initializers.DB.Create(&models.Listbahan{Resep_id: 3, Bahan_id: 6, Jumlahbahan: "2"})
	initializers.DB.Create(&models.Listbahan{Resep_id: 3, Bahan_id: 7, Jumlahbahan: "2"})
	// initializers.DB.Create(&models.Listbahan{Resep_id: 4, Bahan_id: 5, Jumlahbahan: "2"})
	// initializers.DB.Create(&models.Listbahan{Resep_id: 4, Bahan_id: 6, Jumlahbahan: "1"})
	// initializers.DB.Create(&models.Listbahan{Resep_id: 4, Bahan_id: 7, Jumlahbahan: "1"})
	initializers.DB.Create(&models.Listbahan{Resep_id: 4, Bahan_id: 1, Jumlahbahan: "2"})
	initializers.DB.Create(&models.Listbahan{Resep_id: 4, Bahan_id: 6, Jumlahbahan: "2"})
	initializers.DB.Create(&models.Listbahan{Resep_id: 4, Bahan_id: 7, Jumlahbahan: "2"})
	initializers.DB.Create(&models.Listbahan{Resep_id: 5, Bahan_id: 18, Jumlahbahan: "1 kg"})
	initializers.DB.Create(&models.Listbahan{Resep_id: 5, Bahan_id: 19, Jumlahbahan: "1 ikat"})
	initializers.DB.Create(&models.Listbahan{Resep_id: 5, Bahan_id: 20, Jumlahbahan: "500 gram"})
	initializers.DB.Create(&models.Listbahan{Resep_id: 5, Bahan_id: 21, Jumlahbahan: "500 gram"})
	initializers.DB.Create(&models.Listbahan{Resep_id: 5, Bahan_id: 22, Jumlahbahan: "4 kangkang"})
	initializers.DB.Create(&models.Listbahan{Resep_id: 6, Bahan_id: 2, Jumlahbahan: "1 kg"})
	initializers.DB.Create(&models.Listbahan{Resep_id: 6, Bahan_id: 6, Jumlahbahan: "5 pcs"})
	initializers.DB.Create(&models.Listbahan{Resep_id: 6, Bahan_id: 23, Jumlahbahan: "1 saset"})
	initializers.DB.Create(&models.Listbahan{Resep_id: 7, Bahan_id: 16, Jumlahbahan: "1 kg"})
	initializers.DB.Create(&models.Listbahan{Resep_id: 7, Bahan_id: 24, Jumlahbahan: "sesuka hati"})
	initializers.DB.Create(&models.Listbahan{Resep_id: 7, Bahan_id: 25, Jumlahbahan: "sesuka hati"})

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
