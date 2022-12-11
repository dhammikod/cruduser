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
