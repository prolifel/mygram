package models

import (
	"mygram/helpers"
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type User struct {
	ID        uint       `json:"id,omitempty" gorm:"primary_key"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeleteAt  *time.Time `json:"delete_at,omitempty"`
	Username  string     `json:"username,omitempty" gorm:"type:varchar(100);uniqueIndex" validate:"required"`
	Email     string     `json:"email,omitempty" gorm:"type:varchar(100);uniqueIndex" validate:"required,email"`
	Password  string     `json:"password,omitempty" validate:"required,min=6"`
	Age       int        `json:"age,omitempty" gorm:"type:int" validate:"required,min=8"`
}

type APIUser struct {
	ID       uint   `json:",omitempty"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

var userValidate *validator.Validate

func (user *User) BeforeSave(tx *gorm.DB) (err error) {
	userValidate = validator.New()
	err = userValidate.Struct(user)
	return
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.Password = helpers.HashPassword(user.Password)

	err = nil
	return
}
