package controllers

import (
	"mygram/models"

	"github.com/gin-gonic/gin"
)

func (databaseConnection *DatabaseConnection) GetUsers(c *gin.Context) {
	var users []models.User
	databaseConnection.DB.Find(&users)
	c.JSON(200, users)
}

func (databaseConnection *DatabaseConnection) CreateUser(c *gin.Context) {
	var (
		user   models.User
		result gin.H
	)

	// Binding data
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		c.Abort()
		return
	}

	// Insert user data
	if err := databaseConnection.DB.Create(&user).Error; err != nil {
		result = gin.H{
			"error": err.Error(),
		}
		c.JSON(400, result)
		return
	}

	c.JSON(201, user)
}
