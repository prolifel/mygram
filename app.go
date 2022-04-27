package main

import (
	"mygram/config"
	"mygram/controllers"
	"mygram/middlewares"

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
		users.POST("/register", databaseConnection.CreateUser)
		users.POST("/login", databaseConnection.UserLogin)
		users.PUT("/", middlewares.Authentication(), databaseConnection.UpdateUser)
		users.DELETE("/", middlewares.Authentication(), databaseConnection.DeleteUser)
	}

	photos := router.Group("/photos")
	{
		photos.POST("/", middlewares.Authentication(), databaseConnection.CreatePhoto)
		photos.GET("/", middlewares.Authentication(), databaseConnection.GetPhotos)
		photos.PUT("/", middlewares.Authentication(), databaseConnection.UpdatePhoto)
		photos.DELETE("/", middlewares.Authentication(), databaseConnection.DeletePhoto)
	}

	router.Run(":3000")
}
