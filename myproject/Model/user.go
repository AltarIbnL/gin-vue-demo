package Model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	//ID        int
	Name      string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"type:varchar(11);unique";`
	Password  string `gorm:"default:01234567890;size:255;not null"`
}
