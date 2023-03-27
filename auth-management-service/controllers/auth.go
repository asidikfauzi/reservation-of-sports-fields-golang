package controllers

import (
	"asidikfauzi/reservation-of-sport-fields-golang/config"
	"asidikfauzi/reservation-of-sport-fields-golang/lib/database"
	"asidikfauzi/reservation-of-sport-fields-golang/lib/utils"
	"asidikfauzi/reservation-of-sport-fields-golang/models"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"regexp"
	"time"
)

var user []models.Users

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

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

	if !emailRegex.MatchString(userModel.Email) {
		utils.BadRequest(c, "Invalid email format")
		return
	}

	if !utils.ValidatePassword(userModel.Password) {
		utils.BadRequest(c, "Password must be at least 8 "+
			"characters long and contain "+
			"at least one uppercase letter, "+
			"one lowercase letter, and one number")
	}

	checkUsername := config.DB.Where("username = ?", userModel.Username).First(&user)

	if checkUsername.RowsAffected > 0 {
		utils.BadRequest(c, "Username already exists")
		return
	}

	checkEmail := config.DB.Where("email = ?", userModel.Email).First(&user)

	if checkEmail.RowsAffected > 0 {
		utils.BadRequest(c, "Email already exists")
		return
	}

	if err := models.Validate.Struct(userModel); err != nil {
		utils.BadRequest(c, err.Error())
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
		utils.InternalServerError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Success register users",
		"status":  "success",
	})

}
