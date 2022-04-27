package controllers

import (
	"gorm.io/gorm"
)

type DatabaseConnection struct {
	DB *gorm.DB
}
