package controllers

import (
	"mygram/helpers"
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

func (databaseConnection *DatabaseConnection) UserLogin(c *gin.Context) {
	var (
		user           models.User
		result         gin.H
		passwordString string
	)

	if err := c.ShouldBindJSON(&user); err != nil {
		result = gin.H{
			"error": err.Error(),
		}
		c.JSON(400, result)
		return
	}

	passwordString = string(user.Password)

	if err := databaseConnection.DB.Where("email = ?", user.Email).Take(&user).Error; err != nil {
		result = gin.H{
			"error": "user not found",
		}
		c.JSON(404, result)
		return
	}

	if err := helpers.ComparePassword([]byte(user.Password), []byte(passwordString)); !err {
		result = gin.H{
			"error": "password is incorrect",
		}
		c.JSON(401, result)
		return
	}

	token := helpers.GenerateToken(user.ID, user.Email)
	result = gin.H{
		"token": token,
	}
	c.JSON(200, result)
}

func (databaseConnection *DatabaseConnection) GetUser(c *gin.Context) {
	var user models.User
	id := c.Params.ByName("id")
	if err := databaseConnection.DB.Where("id = ?", id).First(&user).Error; err != nil {
		c.JSON(404, gin.H{
			"error": "user not found",
		})
		return
	}
	c.JSON(200, user)
}

func (databaseConnection *DatabaseConnection) UpdateUser(c *gin.Context) {
	var (
		user   models.User
		result gin.H
	)

	id := c.Params.ByName("id")
	if err := databaseConnection.DB.Where("id = ?", id).First(&user).Error; err != nil {
		result = gin.H{
			"error": "user not found",
		}
		c.JSON(404, result)
		return
	}

	if err := c.BindJSON(&user); err != nil {
		result = gin.H{
			"error": err.Error(),
		}
		c.JSON(400, result)
		return
	}

	databaseConnection.DB.Save(&user)
	c.JSON(200, user)
}

func (databaseConnection *DatabaseConnection) DeleteUser(c *gin.Context) {
	var (
		user   models.User
		result gin.H
	)

	id := c.Params.ByName("id")
	if err := databaseConnection.DB.Where("id = ?", id).First(&user).Error; err != nil {
		result = gin.H{
			"error": "user not found",
		}
		c.JSON(404, result)
		return
	}
	databaseConnection.DB.Delete(&user)
	result = gin.H{
		"message": "user deleted",
	}
	c.JSON(200, result)
}
