package models

import "gorm.io/gorm"

type Photo struct {
	gorm.Model
	Title     string `gorm:"type:varchar(100)"`
	Caption   string `gorm:"type:varchar(200)"`
	Photo_url string `gorm:"type:varchar(200)"`
	UserID    uint   `gorm:"type:int"`
}
