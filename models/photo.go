package models

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Photo struct {
	gorm.Model
	Title     string `json:"title" gorm:"type:varchar(100)" validate:"required"`
	Caption   string `json:"caption" gorm:"type:varchar(200)"`
	Photo_url string `json:"photo_url" gorm:"type:varchar(200)" validate:"required"`
	UserID    uint   `json:"user_id" gorm:"type:int"`
}

var photoValidate *validator.Validate

func (photo *Photo) BeforeSave(tx *gorm.DB) (err error) {
	photoValidate = validator.New()
	err = photoValidate.Struct(photo)
	return
}
