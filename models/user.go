package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"type:varchar(100);uniqueIndex" binding:"required"`
	Email    string `json:"email" gorm:"type:varchar(100);uniqueIndex" binding:"required,email"`
	Password string `json:"password" gorm:"type:varchar(100)" binding:"required,min=6"`
	Age      int    `json:"age" gorm:"type:int" binding:"required,min=8"`
}
