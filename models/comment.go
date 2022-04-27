package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Comment struct {
	// gorm.Model
	ID        uint       `json:"id,omitempty" gorm:"primary_key"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeleteAt  *time.Time `json:"delete_at,omitempty"`
	UserID    uint       `json:"user_id" gorm:"type:int"`
	PhotoID   uint       `json:"photo_id" gorm:"type:int"`
	Message   string     `json:"message" gorm:"type:varchar(200)" validate:"required"`
	User      User       `json:"user" validate:"-"`
	Photo     Photo      `json:"photo" validate:"-"`
}

type APIComment struct {
	ID        uint       `json:"id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	UserID    uint       `json:"user_id,omitempty"`
	PhotoID   uint       `json:"photo_id,omitempty"`
	Message   string     `json:"message,omitempty"`
	User      User       `json:"user,omitempty"`
	Photo     Photo      `json:"photo,omitempty"`
}

var commentValidate *validator.Validate

func (comment *Comment) BeforeSave(tx *gorm.DB) (err error) {
	commentValidate = validator.New()
	err = commentValidate.Struct(comment)
	return
}
