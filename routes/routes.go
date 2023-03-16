package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Start() {
	// Create a new Gin router
	r := gin.Default()

	// Define a route for the GET method
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	// Start the server on port 8080
	r.Run(":8080")
}
