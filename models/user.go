package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(100);unique_index"`
	Email    string `gorm:"type:varchar(100);unique_index"`
	Password string `gorm:"type:varchar(100)"`
	Age      int    `gorm:"type:int"`
}
