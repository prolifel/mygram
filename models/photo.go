package models

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Photo struct {
	gorm.Model
	Title     string `json:"title,omitempty" gorm:"type:varchar(100)" validate:"required"`
	Caption   string `json:"caption,omitempty" gorm:"type:varchar(200)"`
	Photo_url string `json:"photo_url,omitempty" gorm:"type:varchar(200)" validate:"required"`
	UserID    uint   `json:"user_id,omitempty" gorm:"type:int"`
	User      User   `json:"user,omitempty" validate:"-"`
}

type APIPhoto struct {
	ID        uint   `json:"id,omitempty"`
	Title     string `json:"title,omitempty"`
	Caption   string `json:"caption,omitempty"`
	Photo_url string `json:"photo_url,omitempty"`
	UserID    uint   `json:"user_id,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	User      User   `json:"user,omitempty"`
}

var photoValidate *validator.Validate

func (photo *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	photoValidate = validator.New()
	err = photoValidate.Struct(photo)
	return
}
