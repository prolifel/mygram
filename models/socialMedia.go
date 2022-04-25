package models

type SocialMedia struct {
	ID               uint   `gorm:"primary_key"`
	Name             string `gorm:"type:varchar(100)"`
	Social_media_url string `gorm:"type:varchar(200)"`
	UserID           uint   `gorm:"type:int"`
}
