package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type ErrorMessageEmpty struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func GetErrorMessageEmpty(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "lte":
		return "Should be less than " + fe.Param()
	case "gte":
		return "Should be greater than " + fe.Param()
	case "oneof":
		return "these fields should be user and admin"
	case "email":
		return "format email not valid"
	}
	return "Unknown error"
}

func BadRequest(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"code":    400,
		"message": message,
		"status":  "Bad Request",
	})
}

func InternalServerError(c *gin.Context, message string) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
		"code":    500,
		"message": message,
		"status":  "Internal Server Error",
	})
}
