package controllers

import (
	"asidikfauzi/reservation-of-sport-fields-golang/lib/database"
	"asidikfauzi/reservation-of-sport-fields-golang/lib/utils"
	"asidikfauzi/reservation-of-sport-fields-golang/models"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"time"
)

func RegisterController(c *gin.Context) {
	var userModel models.Users_register

	if err := c.ShouldBindJSON(&userModel); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]utils.ErrorMessageEmpty, len(ve))
			for i, fe := range ve {
				out[i] = utils.ErrorMessageEmpty{Field: fe.Field(), Message: utils.GetErrorMessageEmpty(fe)}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code":    400,
				"status":  "Bad Request",
				"message": out,
			})
		}
		return
	}

	//if err := c.BindJSON(&userModel); err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"code":    400,
	//		"message": "Error registration users",
	//		"status":  err.Error(),
	//	})
	//	return
	//}

	if err := models.Validate.Struct(userModel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Error registration users",
			"status":  err.Error(),
		})
		return
	}

	if !utils.ValidatePassword(userModel.Password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"message": "Password must be at least 8 " +
				"characters long and contain " +
				"at least one uppercase letter, " +
				"one lowercase letter, and one number",
			"status": "Error",
		})
		return
	}

	hashPassword, _ := utils.HashPassword(userModel.Password)

	var users models.Users
	users.ID = utils.Uuid()
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
