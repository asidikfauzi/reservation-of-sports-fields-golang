package routes

import (
	"asidikfauzi/reservation-of-sport-fields-golang/controllers"
	"github.com/gin-gonic/gin"
)

func Start() {
	// Create a new Gin router
	r := gin.Default()

	// Define a route for the GET method
	r.POST("/register", controllers.RegisterController)

	// Start the server on port 8080
	r.Run(":8080")
}
