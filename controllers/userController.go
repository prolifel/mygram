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
