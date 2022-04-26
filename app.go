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

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/users", databaseConnection.GetUsers)

	router.Run(":3000")
}
