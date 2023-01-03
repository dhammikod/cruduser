package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name            string
	Email           string
	Password        string
	Profile_picture string
	Notification    bool
	Resep           []Resep `gorm:"many2many:Saved_Recipes;"`
}

type Resep struct {
	gorm.Model
	Description  string
	Timetaken    string
	Totalcal     int
	Judul        string
	Portionsize  int
	Foto         string `gorm:"type:longblob;default:null"`
	Video        string `gorm:"type:longblob;default:null"`
	Steps        string `gorm:"type:text"`
	Created_by   int
	User         User `gorm:"foreignKey:Created_by"`
	Rating       float32
	Jumlahrating int
	Users        []User      `gorm:"many2many:Saved_Recipes;"`
	Listbahan    []Listbahan `gorm:"foreignKey:Resep_id"`
}

type Bahan struct {
	gorm.Model
	Namabahan string
	Foto      string
}

type Listbahan struct {
	gorm.Model
	Bahan_id    int
	Bahan       Bahan `gorm:"foreignKey:Bahan_id"`
	Jumlahbahan string
	Resep_id    uint
}

type Savedrecipe struct {
	gorm.Model
	User_id  int
	User     User `gorm:"foreignKey:User_id"`
	Resep_id int
	Resep    Resep `gorm:"foreignKey:Resep_id"`
}
