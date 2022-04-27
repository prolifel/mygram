package controllers

import (
	"mygram/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// POST /photos
func (databaseConnection *DatabaseConnection) CreatePhoto(c *gin.Context) {
	var (
		photo  models.Photo
		result gin.H
	)

	userData := c.MustGet("userData").(jwt.MapClaims)

	// Binding data
	if err := c.ShouldBindJSON(&photo); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		c.Abort()
		return
	}

	photo.UserID = uint(userData["id"].(float64))

	// Insert photo data
	if err := databaseConnection.DB.Omit(clause.Associations).Create(&photo).Error; err != nil {
		result = gin.H{
			"error": err.Error(),
		}
		c.JSON(400, result)
		return
	}

	c.JSON(201, gin.H{
		"id":         photo.ID,
		"title":      photo.Title,
		"caption":    photo.Caption,
		"photo_url":  photo.Photo_url,
		"user_id":    photo.UserID,
		"created_at": photo.CreatedAt,
	})
}

// GET /photos
func (databaseConnection *DatabaseConnection) GetPhotos(c *gin.Context) {
	var (
		photos    []models.Photo
		photosAPI []models.APIPhoto
		result    gin.H
	)

	userData := c.MustGet("userData").(jwt.MapClaims)

	if err := databaseConnection.DB.Debug().Model(&photos).Where("user_id = ?", uint(userData["id"].(float64))).Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Model(&models.User{}).Find(&models.APIUser{})
	}).Find(&photosAPI).Error; err != nil {
		result = gin.H{
			"error": err.Error(),
		}
		c.JSON(404, result)
		return
	}

	c.JSON(200, gin.H{
		"photos": photosAPI,
	})
}

// PUT /photos/:photoId
func (databaseConnection *DatabaseConnection) UpdatePhoto(c *gin.Context) {
	var (
		photo  models.Photo
		result gin.H
	)

	userData := c.MustGet("userData").(jwt.MapClaims)
	id, _ := strconv.ParseInt(c.Query("photoId"), 10, 64)

	if err := databaseConnection.DB.Where("id = ?", id).First(&photo).Error; err != nil {
		result = gin.H{
			"error": "photo not found",
		}
		c.JSON(404, result)
		return
	}

	if int64(userData["id"].(float64)) != int64(photo.UserID) {
		result = gin.H{
			"error": "you're not allowed to do this",
		}
		c.JSON(401, result)
		return
	}

	if err := c.ShouldBindJSON(&photo); err != nil {
		result = gin.H{
			"error": err.Error(),
		}
		c.JSON(400, result)
		return
	}

	databaseConnection.DB.Save(&photo)
	result = gin.H{
		"id":         photo.ID,
		"title":      photo.Title,
		"caption":    photo.Caption,
		"photo_url":  photo.Photo_url,
		"user_id":    photo.UserID,
		"updated_at": photo.UpdatedAt,
	}
	c.JSON(200, result)
}

// DELETE /photos/
func (databaseConnection *DatabaseConnection) DeletePhoto(c *gin.Context) {
	var (
		photo  models.Photo
		result gin.H
	)

	userData := c.MustGet("userData").(jwt.MapClaims)
	id, _ := strconv.ParseInt(c.Query("photoId"), 10, 64)

	if err := databaseConnection.DB.Where("id = ?", id).First(&photo).Error; err != nil {
		result = gin.H{
			"error": "photo not found",
		}
		c.JSON(404, result)
		return
	}

	if int64(userData["id"].(float64)) != int64(photo.UserID) {
		result = gin.H{
			"error": "you're not allowed to do this",
		}
		c.JSON(401, result)
		return
	}

	databaseConnection.DB.Delete(&photo)
	result = gin.H{
		"message": "Your photo has been successfully deleted",
	}
	c.JSON(200, result)
}
