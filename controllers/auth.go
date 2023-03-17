package controllers

import (
	"asidikfauzi/reservation-of-sport-fields-golang/helpers"
	"asidikfauzi/reservation-of-sport-fields-golang/lib/database"
	"asidikfauzi/reservation-of-sport-fields-golang/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func RegisterController(c *gin.Context) {
	var userModel models.Users_register

	if err := c.BindJSON(&userModel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Error registration users",
			"status":  err.Error(),
		})
		return
	}

	if err := models.Validate.Struct(userModel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Error registration users",
			"status":  err.Error(),
		})
		return
	}

	hashPassword, _ := helpers.HashPassword(userModel.Password)

	var users models.Users
	users.ID = helpers.Uuid()
	users.Username = userModel.Username
	users.Email = userModel.Email
	users.Password = hashPassword
	users.Role_id = userModel.Role_id
	users.IsActive = true
	users.CreatedAt = time.Now()

	_, err := database.InsertUsers(users)

	if err != nil {
		// Return an HTTP error with details of the error
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    500,
			"message": "Error registration users",
			"status":  err.Error(),
		})
		return

	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success register users",
		"status":  "success",
	})

}
