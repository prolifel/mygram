package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	UserID  uint   `gorm:"type:int"`
	PhotoID uint   `gorm:"type:int"`
	Message string `gorm:"type:varchar(200)"`
}
