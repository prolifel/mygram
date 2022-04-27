package main

import (
	"mygram/config"
	"mygram/controllers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}

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
		users.POST("/login", databaseConnection.UserLogin)
	}

	router.Run(":3000")
}
