package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	r := gin.Default()

	r.GET("/v1/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "UPDATE FERIII!",
		})
	})

	r.GET("/v1/greeting", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, UPDATE FERIII!",
		})
	})

	r.GET("/v1/about", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Ini adalah hands-on deployment VPS UPDATE FERIII",
		})
	})

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000" // Default port
	}
	fmt.Println("Server is running on port " + port)
	r.Run(":" + port)
}
