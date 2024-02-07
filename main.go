package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// Create a Gin router with default middleware: logger and recovery (crash-free) middleware.
	router := gin.Default()

	// Define a GET route that returns a "Hello World" JSON response.
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})

	// Start the server on port 8080.
	router.Run(":9000")
}
