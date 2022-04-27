package main

import (
	"mygram/config"
	"mygram/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.DBInit()
	databaseConnection := &controllers.DatabaseConnection{DB: db}

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "mygram served successfully 🚀",
		})
	})

	users := router.Group("/users")
	{
		users.GET("/", databaseConnection.GetUsers)
		users.POST("/register", databaseConnection.CreateUser)
	}

	router.Run(":3000")
}
