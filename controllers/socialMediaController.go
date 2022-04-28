package controllers

import (
	"mygram/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// POST /socialmedias
func (databaseConnection *DatabaseConnection) CreateSocialMedia(c *gin.Context) {
	var (
		socialMedia models.SocialMedia
		result      gin.H
	)

	userData := c.MustGet("userData").(jwt.MapClaims)

	// Binding data
	if err := c.ShouldBindJSON(&socialMedia); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		c.Abort()
		return
	}

	socialMedia.UserID = uint(userData["id"].(float64))

	// Insert socialMedia data
	if err := databaseConnection.DB.Omit(clause.Associations).Create(&socialMedia).Error; err != nil {
		result = gin.H{
			"error": err.Error(),
		}
		c.JSON(400, result)
		return
	}

	c.JSON(201, gin.H{
		"id":               socialMedia.ID,
		"name":             socialMedia.Name,
		"social_media_url": socialMedia.Social_media_url,
		"user_id":          socialMedia.UserID,
		"created_at":       socialMedia.CreatedAt,
	})
}

// GET /socialmedias
func (databaseConnection *DatabaseConnection) GetSocialMedias(c *gin.Context) {
	var (
		socialmedias []models.SocialMedia
		result       gin.H
	)

	userData := c.MustGet("userData").(jwt.MapClaims)

	if err := databaseConnection.DB.Debug().Model(&socialmedias).Where("user_id = ?", uint(userData["id"].(float64))).Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Model(&models.User{}).Find(&models.APIUser{})
	}).Find(&socialmedias).Error; err != nil {
		result = gin.H{
			"error": err.Error(),
		}
		c.JSON(404, result)
		return
	}

	c.JSON(200, gin.H{
		"socialmedias": socialmedias,
	})
}

// PUT /socialmedias/:socialMediaId
func (databaseConnection *DatabaseConnection) UpdateSocialMedia(c *gin.Context) {
	var (
		socialMedia models.SocialMedia
		result      gin.H
	)

	userData := c.MustGet("userData").(jwt.MapClaims)
	id, _ := strconv.ParseInt(c.Query("socialMediaId"), 10, 64)

	if err := databaseConnection.DB.Where("id = ?", id).First(&socialMedia).Error; err != nil {
		result = gin.H{
			"error": "socialMedia not found",
		}
		c.JSON(404, result)
		return
	}

	if int64(userData["id"].(float64)) != int64(socialMedia.UserID) {
		result = gin.H{
			"error": "you're not allowed to do this",
		}
		c.JSON(401, result)
		return
	}

	if err := c.ShouldBindJSON(&socialMedia); err != nil {
		result = gin.H{
			"error": err.Error(),
		}
		c.JSON(400, result)
		return
	}

	databaseConnection.DB.Save(&socialMedia)

	if err := databaseConnection.DB.Where("id = ?", id).First(&socialMedia).Error; err != nil {
		result = gin.H{
			"error": err.Error(),
		}
		c.JSON(404, result)
		return
	}

	result = gin.H{
		"id":               socialMedia.ID,
		"name":             socialMedia.Name,
		"social_media_url": socialMedia.Social_media_url,
		"user_id":          socialMedia.UserID,
		"updated_at":       socialMedia.UpdatedAt,
	}
	c.JSON(200, result)
}

// DELETE /socialmedias/:socialMediaId
func (databaseConnection *DatabaseConnection) DeleteSocialMedia(c *gin.Context) {
	var (
		socialMedia models.SocialMedia
		result      gin.H
	)

	userData := c.MustGet("userData").(jwt.MapClaims)
	id, _ := strconv.ParseInt(c.Query("socialMediaId"), 10, 64)

	if err := databaseConnection.DB.Where("id = ?", id).First(&socialMedia).Error; err != nil {
		result = gin.H{
			"error": "socialMedia not found",
		}
		c.JSON(404, result)
		return
	}

	if int64(userData["id"].(float64)) != int64(socialMedia.UserID) {
		result = gin.H{
			"error": "you're not allowed to do this",
		}
		c.JSON(401, result)
		return
	}

	databaseConnection.DB.Delete(&socialMedia)
	result = gin.H{
		"message": "Your social media has been successfully deleted",
	}
	c.JSON(200, result)
}
