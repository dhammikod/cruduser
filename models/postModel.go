package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name            string
	Email           string
	Password        string
	No_telp         string
	Profile_picture string
	Notification    bool
}

type Resep struct {
	gorm.Model
	Rating      bool
	Descriptoin string
	Judul       string
	Portionsize int
	Foto        string
	Video       string
}

type Bahan struct {
	gorm.Model
	Namabahan string
}

type Requiredingredients struct {
	gorm.Model
	Bahan_id    int
	Bahan       Bahan `gorm:"foreignKey:Bahan_id"`
	Jumlahbahan int
}

type Listbahan struct {
	gorm.Model
	Req_id              int
	Requiredingredients Requiredingredients `gorm:"foreignKey:Req_id"`
	Resep_id            int
	Resep               Resep `gorm:"foreignKey:Resep_id"`
}

type Savedrecipe struct {
	gorm.Model
	User_id  int
	User     User `gorm:"foreignKey:User_id"`
	Resep_id int
	Resep    Resep `gorm:"foreignKey:Resep_id"`
}
