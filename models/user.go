package models

import (
	"mygram/helpers"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"type:varchar(100);uniqueIndex" validate:"required"`
	Email    string `json:"email" gorm:"type:varchar(100);uniqueIndex" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Age      int    `json:"age" gorm:"type:int" validate:"required,min=8"`
}

var validate *validator.Validate

func (user *User) BeforeSave(tx *gorm.DB) (err error) {
	validate = validator.New()
	err = validate.Struct(user)
	return
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.Password = helpers.HashPassword(user.Password)

	err = nil
	return
}
