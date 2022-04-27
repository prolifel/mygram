package models

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type SocialMedia struct {
	ID               uint   `json:"id" gorm:"primary_key"`
	Name             string `json:"name" gorm:"type:varchar(100)" validate:"required"`
	Social_media_url string `json:"social_media_url" gorm:"type:varchar(200)" validate:"required"`
	UserID           uint   `json:"user_id" gorm:"type:int"`
}

var socialMediaValidate *validator.Validate

func (socialMedia *SocialMedia) BeforeSave(tx *gorm.DB) (err error) {
	socialMediaValidate = validator.New()
	err = socialMediaValidate.Struct(socialMedia)
	return
}
