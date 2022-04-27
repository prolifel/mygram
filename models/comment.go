package models

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	UserID  uint   `json:"user_id" gorm:"type:int"`
	PhotoID uint   `json:"photo_id" gorm:"type:int"`
	Message string `json:"message" gorm:"type:varchar(200)" validate:"required"`
}

var commentValidate *validator.Validate

func (comment *Comment) BeforeSave(tx *gorm.DB) (err error) {
	commentValidate = validator.New()
	err = commentValidate.Struct(comment)
	return
}
