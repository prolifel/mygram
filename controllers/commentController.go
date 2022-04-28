package controllers

import (
	"mygram/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// POST /comments
func (databaseConnection *DatabaseConnection) CreateComment(c *gin.Context) {
	var (
		comment models.Comment
		result  gin.H
	)

	userData := c.MustGet("userData").(jwt.MapClaims)

	// Binding data
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		c.Abort()
		return
	}

	comment.UserID = uint(userData["id"].(float64))

	// Insert comment data
	if err := databaseConnection.DB.Omit(clause.Associations).Create(&comment).Error; err != nil {
		result = gin.H{
			"error": err.Error(),
		}
		c.JSON(400, result)
		return
	}

	c.JSON(201, gin.H{
		"id":         comment.ID,
		"message":    comment.Message,
		"photo_id":   comment.PhotoID,
		"user_id":    comment.UserID,
		"created_at": comment.CreatedAt,
	})
}

// GET /comments
func (databaseConnection *DatabaseConnection) GetComments(c *gin.Context) {
	var (
		comments    []models.Comment
		commentsAPI []models.APIComment
		result      gin.H
	)

	userData := c.MustGet("userData").(jwt.MapClaims)

	if err := databaseConnection.DB.Debug().Model(&comments).Where("user_id = ?", uint(userData["id"].(float64))).Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Model(&models.User{}).Find(&models.APIUser{})
	}).Preload("Photo", func(db *gorm.DB) *gorm.DB {
		return db.Model(&models.Photo{}).Find(&models.APIPhoto{})
	}).Find(&commentsAPI).Error; err != nil {
		result = gin.H{
			"error": err.Error(),
		}
		c.JSON(404, result)
		return
	}

	c.JSON(200, gin.H{
		"comments": commentsAPI,
	})
}

// PUT /comments/:commentId
func (databaseConnection *DatabaseConnection) UpdateComment(c *gin.Context) {
	var (
		comment models.Comment
		result  gin.H
	)

	userData := c.MustGet("userData").(jwt.MapClaims)
	id, _ := strconv.ParseInt(c.Query("commentId"), 10, 64)

	if err := databaseConnection.DB.Where("id = ?", id).First(&comment).Error; err != nil {
		result = gin.H{
			"error": "comment not found",
		}
		c.JSON(404, result)
		return
	}

	if int64(userData["id"].(float64)) != int64(comment.UserID) {
		result = gin.H{
			"error": "you're not allowed to do this",
		}
		c.JSON(401, result)
		return
	}

	if err := c.ShouldBindJSON(&comment); err != nil {
		result = gin.H{
			"error": err.Error(),
		}
		c.JSON(400, result)
		return
	}

	databaseConnection.DB.Save(&comment)

	if err := databaseConnection.DB.Preload("Photo").Where("id = ?", id).First(&comment).Error; err != nil {
		result = gin.H{
			"error": err.Error(),
		}
		c.JSON(404, result)
		return
	}

	result = gin.H{
		"id":         comment.ID,
		"title":      comment.Photo.Title,
		"caption":    comment.Photo.Caption,
		"photo_url":  comment.Photo.Photo_url,
		"user_id":    comment.UserID,
		"updated_at": comment.UpdatedAt,
	}
	c.JSON(200, result)
}

// DELETE /comment/:commentId
func (databaseConnection *DatabaseConnection) DeleteComment(c *gin.Context) {
	var (
		comment models.Comment
		result  gin.H
	)

	userData := c.MustGet("userData").(jwt.MapClaims)
	id, _ := strconv.ParseInt(c.Query("commentId"), 10, 64)

	if err := databaseConnection.DB.Where("id = ?", id).First(&comment).Error; err != nil {
		result = gin.H{
			"error": "comment not found",
		}
		c.JSON(404, result)
		return
	}

	if int64(userData["id"].(float64)) != int64(comment.UserID) {
		result = gin.H{
			"error": "you're not allowed to do this",
		}
		c.JSON(401, result)
		return
	}

	databaseConnection.DB.Delete(&comment)
	result = gin.H{
		"message": "Your comment has been successfully deleted",
	}
	c.JSON(200, result)
}
