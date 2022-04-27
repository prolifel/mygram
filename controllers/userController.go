package controllers

import (
	"mygram/helpers"
	"mygram/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// POST /users/register
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

	c.JSON(201, gin.H{
		"age":      user.Age,
		"email":    user.Email,
		"id":       user.ID,
		"username": user.Username,
	})
}

// POST /users/login
func (databaseConnection *DatabaseConnection) UserLogin(c *gin.Context) {
	var (
		user           models.User
		result         gin.H
		passwordString string
	)

	// Binding data
	if err := c.ShouldBindJSON(&user); err != nil {
		result = gin.H{
			"error": err.Error(),
		}
		c.JSON(400, result)
		return
	}

	passwordString = string(user.Password)

	// Check if user exists
	if err := databaseConnection.DB.Where("email = ?", user.Email).Take(&user).Error; err != nil {
		result = gin.H{
			"error": "user not found",
		}
		c.JSON(404, result)
		return
	}

	// Check if password is correct
	if err := helpers.ComparePassword([]byte(user.Password), []byte(passwordString)); !err {
		result = gin.H{
			"error": "password is incorrect",
		}
		c.JSON(401, result)
		return
	}

	// Create token
	token := helpers.GenerateToken(user.ID, user.Email)
	result = gin.H{
		"token": token,
	}
	c.JSON(200, result)
}

// PUT /users/:id
func (databaseConnection *DatabaseConnection) UpdateUser(c *gin.Context) {
	var (
		user   models.User
		result gin.H
	)

	userData := c.MustGet("userData").(jwt.MapClaims)
	id, _ := strconv.ParseInt(c.Query("userId"), 10, 64)

	if int64(userData["id"].(float64)) != id {
		result = gin.H{
			"error": "you're not allowed to do this",
		}
		c.JSON(401, result)
		return
	}

	if err := databaseConnection.DB.Where("id = ?", id).First(&user).Error; err != nil {
		result = gin.H{
			"error": "user not found",
		}
		c.JSON(404, result)
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		result = gin.H{
			"error": err.Error(),
		}
		c.JSON(400, result)
		return
	}

	databaseConnection.DB.Save(&user)
	result = gin.H{
		"id":         user.ID,
		"email":      user.Email,
		"username":   user.Username,
		"age":        user.Age,
		"updated_at": user.UpdatedAt,
	}
	c.JSON(200, result)
}

// Delete /users/
func (databaseConnection *DatabaseConnection) DeleteUser(c *gin.Context) {
	var (
		user   models.User
		result gin.H
	)

	userData := c.MustGet("userData").(jwt.MapClaims)
	userDataId := int64(userData["id"].(float64))

	if err := databaseConnection.DB.Where("id = ?", userDataId).First(&user).Error; err != nil {
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
